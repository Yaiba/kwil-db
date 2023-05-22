package sql_parser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kwilteam/kwil-db/pkg/engine/tree"
	"github.com/kwilteam/kwil-db/pkg/sql_parser/sqlite"
	"reflect"
	"strings"
)

// KFSqliteVisitor is visitor that visit Antlr parsed tree and returns the AST.
type KFSqliteVisitor struct {
	sqlite.BaseSQLiteParserVisitor
	ErrorHandler

	actionCtx ActionContext
	dbCtx     DatabaseContext

	trace bool
}

var _ sqlite.SQLiteParserVisitor = &KFSqliteVisitor{}

func NewKFSqliteVisitor(eh *ErrorHandler, actionName string, ctx DatabaseContext) *KFSqliteVisitor {
	k := &KFSqliteVisitor{
		ErrorHandler: *eh,
		actionCtx:    ctx.Actions[actionName],
		dbCtx:        ctx,
		trace:        false,
	}
	return k
}

type cteTableName struct {
	table   string
	columns []string
}

type withClause struct {
	tableName cteTableName
}

func (v *KFSqliteVisitor) visitCteTableName(ctx sqlite.ICte_table_nameContext) (tableName cteTableName) {
	tableName.table = ctx.Table_name().GetText()
	colNameCtxs := ctx.AllColumn_name()
	for _, colName := range colNameCtxs {
		tableName.columns = append(tableName.columns, colName.GetText())
	}

	return tableName
}

func (v *KFSqliteVisitor) visitCommonTableExpression(ctx sqlite.ICommon_table_expressionContext) *tree.CTE {
	cte := tree.CTE{}
	cteTableCtx := ctx.Cte_table_name()
	tableName := v.visitCteTableName(cteTableCtx)
	cte.Table = tableName.table
	if tableName.columns != nil {
		cte.Columns = tableName.columns
	}

	selectStmtCoreCtx := ctx.Select_stmt_core()
	selectStmt := v.Visit(selectStmtCoreCtx).(tree.SelectStmt)
	cte.Select = &selectStmt
	return &cte
}

// VisitCommon_table_stmt is called when visiting a common_table_stmt, return []*tree.CTE.
func (v *KFSqliteVisitor) VisitCommon_table_stmt(ctx *sqlite.Common_table_stmtContext) interface{} {
	if ctx == nil {
		return nil
	}

	cteCount := len(ctx.AllCommon_table_expression())
	ctes := make([]*tree.CTE, cteCount)
	for i := 0; i < cteCount; i++ {
		cteCtx := ctx.Common_table_expression(i)
		cte := v.visitCommonTableExpression(cteCtx)
		ctes[i] = cte
	}
	return ctes
}

func getInsertType(ctx *sqlite.Insert_stmtContext) tree.InsertType {
	var insertType tree.InsertType
	if ctx.OR_() != nil {
		switch {
		case ctx.REPLACE_() != nil:
			insertType = tree.InsertTypeInsertOrReplace
		}
	} else {
		if ctx.REPLACE_() != nil {
			insertType = tree.InsertTypeReplace
		} else {
			insertType = tree.InsertTypeInsert
		}
	}

	return insertType
}

func (v *KFSqliteVisitor) visitExprList(exprList []sqlite.IExprContext) *tree.ExpressionList {
	exprCount := len(exprList)
	exprs := make([]tree.Expression, exprCount)
	for i, exprCtx := range exprList {
		exprs[i] = v.visitExpr(exprCtx)
	}
	fmt.Printf("expr list: %+v\n", exprs)
	return &tree.ExpressionList{Expressions: exprs}
}

// VisitExpr is called when visiting an expression, return tree.Expression.
func (v *KFSqliteVisitor) VisitExpr(ctx *sqlite.ExprContext) interface{} {
	return v.visitExpr(ctx)
}

func (v *KFSqliteVisitor) visitExpr(ctx sqlite.IExprContext) tree.Expression {
	if ctx == nil {
		return nil
	}

	fmt.Println("visitExpr..............: ", ctx.GetText())

	// handle expr list
	startNode, ok := ctx.GetChild(0).(antlr.TerminalNode)
	if ok && startNode.GetSymbol().GetTokenType() == sqlite.SQLiteParserOPEN_PAR {
		return v.visitExprList(ctx.AllExpr())
	}

	switch {
	case ctx.Literal_value() != nil:
		return &tree.ExpressionLiteral{Value: string(ctx.Literal_value().GetText())}
	case ctx.BIND_PARAMETER() != nil:
		return &tree.ExpressionBindParameter{Parameter: ctx.BIND_PARAMETER().GetText()}
	case ctx.Table_name() != nil || ctx.Column_name() != nil:
		expr := &tree.ExpressionColumn{}
		if ctx.Table_name() != nil {
			expr.Table = ctx.Table_name().GetText()
		}
		if ctx.Column_name() != nil {
			expr.Column = ctx.Column_name().GetText()
		}
		return expr
	// binary opertors
	//case ctx.PIPE2() != nil: // TODO: ??
	//	return &tree.ExpressionBinaryComparison{
	//		Left:     v.visitExpr(ctx.Expr(0)),
	//		Right:    v.visitExpr(ctx.Expr(1)),
	//		Operator: tree.BitwiseOperatorBitwiseOr,
	//	}
	case ctx.STAR() != nil:
		return &tree.ExpressionBinaryComparison{
			Left:     v.visitExpr(ctx.Expr(0)),
			Right:    v.visitExpr(ctx.Expr(1)),
			Operator: tree.ArithmeticOperatorMultiply,
		}
	case ctx.DIV() != nil:
		return &tree.ExpressionBinaryComparison{
			Left:     v.visitExpr(ctx.Expr(0)),
			Right:    v.visitExpr(ctx.Expr(1)),
			Operator: tree.ArithmeticOperatorDivide,
		}
	case ctx.MOD() != nil:
		return &tree.ExpressionBinaryComparison{
			Left:     v.visitExpr(ctx.Expr(0)),
			Right:    v.visitExpr(ctx.Expr(1)),
			Operator: tree.ArithmeticOperatorModulus,
		}
	case ctx.PLUS() != nil:
		return &tree.ExpressionBinaryComparison{
			Left:     v.visitExpr(ctx.Expr(0)),
			Right:    v.visitExpr(ctx.Expr(1)),
			Operator: tree.ArithmeticOperatorAdd,
		}
	case ctx.MINUS() != nil:
		return &tree.ExpressionBinaryComparison{
			Left:     v.visitExpr(ctx.Expr(0)),
			Right:    v.visitExpr(ctx.Expr(1)),
			Operator: tree.ArithmeticOperatorSubtract,
		}
	case ctx.LT2() != nil:
		return &tree.ExpressionBinaryComparison{
			Left:     v.visitExpr(ctx.Expr(0)),
			Right:    v.visitExpr(ctx.Expr(1)),
			Operator: tree.BitwiseOperatorLeftShift,
		}
	case ctx.GT2() != nil:
		return &tree.ExpressionBinaryComparison{
			Left:     v.visitExpr(ctx.Expr(0)),
			Right:    v.visitExpr(ctx.Expr(1)),
			Operator: tree.BitwiseOperatorRightShift,
		}
	case ctx.AMP() != nil:
		return &tree.ExpressionBinaryComparison{
			Left:     v.visitExpr(ctx.Expr(0)),
			Right:    v.visitExpr(ctx.Expr(1)),
			Operator: tree.BitwiseOperatorAnd,
		}
	case ctx.PIPE() != nil:
		return &tree.ExpressionBinaryComparison{
			Left:     v.visitExpr(ctx.Expr(0)),
			Right:    v.visitExpr(ctx.Expr(1)),
			Operator: tree.BitwiseOperatorOr,
		}
	case ctx.LT() != nil:
		return &tree.ExpressionBinaryComparison{
			Left:     v.visitExpr(ctx.Expr(0)),
			Right:    v.visitExpr(ctx.Expr(1)),
			Operator: tree.ComparisonOperatorLessThan,
		}
	case ctx.LT_EQ() != nil:
		return &tree.ExpressionBinaryComparison{
			Left:     v.visitExpr(ctx.Expr(0)),
			Right:    v.visitExpr(ctx.Expr(1)),
			Operator: tree.ComparisonOperatorLessThanOrEqual,
		}
	case ctx.GT() != nil:
		return &tree.ExpressionBinaryComparison{
			Left:     v.visitExpr(ctx.Expr(0)),
			Right:    v.visitExpr(ctx.Expr(1)),
			Operator: tree.ComparisonOperatorGreaterThan,
		}
	case ctx.GT_EQ() != nil:
		return &tree.ExpressionBinaryComparison{
			Left:     v.visitExpr(ctx.Expr(0)),
			Right:    v.visitExpr(ctx.Expr(1)),
			Operator: tree.ComparisonOperatorGreaterThanOrEqual,
		}
	//case ctx.ASSIGN() != nil:
	//	return &tree.ExpressionBinaryComparison{
	//		Left:     v.visitExpr(ctx.Expr(0)),
	//		Right:    v.visitExpr(ctx.Expr(1)),
	//		Operator: tree.ComparisonOperatorAssign, // TODO: assign
	//	}
	case ctx.EQ() != nil:
		return &tree.ExpressionBinaryComparison{
			Left:     v.visitExpr(ctx.Expr(0)),
			Right:    v.visitExpr(ctx.Expr(1)), // TODO: equal
			Operator: tree.ComparisonOperatorEqual,
		}
	case ctx.NOT_EQ1() != nil:
		return &tree.ExpressionBinaryComparison{
			Left:     v.visitExpr(ctx.Expr(0)),
			Right:    v.visitExpr(ctx.Expr(1)),
			Operator: tree.ComparisonOperatorNotEqual,
		}
	//case ctx.NOT_EQ2() != nil:
	//	return &tree.ExpressionBinaryComparison{
	//		Left:     v.visitExpr(ctx.Expr(0)),
	//		Right:    v.visitExpr(ctx.Expr(1)),
	//		Operator: tree.ComparisonOperatorNotEqual2,
	//	}
	case ctx.IS_() != nil:
		e := &tree.ExpressionBinaryComparison{
			Left:     v.visitExpr(ctx.Expr(0)),
			Right:    v.visitExpr(ctx.Expr(1)),
			Operator: tree.ComparisonOperatorIs,
		}
		if ctx.NOT_() != nil {
			e.Operator = tree.ComparisonOperatorIsNot
		}
		return e
	case ctx.IN_() != nil && ctx.OPEN_PAR() == nil:
		return &tree.ExpressionBinaryComparison{
			Left:     v.visitExpr(ctx.Expr(0)),
			Right:    v.visitExpr(ctx.Expr(1)),
			Operator: tree.ComparisonOperatorIn,
		}
	//case ctx.LIKE_() != nil:
	//	e := &tree.ExpressionBinaryComparison{
	//		Left:     v.visitExpr(ctx.Expr(0)),
	//		Right:    v.visitExpr(ctx.Expr(1)),
	//		Operator: tree.ComparisonOperatorLike,
	//	}
	//case ctx.MATCH_()	!= nil:
	//	e := &tree.ExpressionBinaryComparison{
	//		Left:     v.visitExpr(ctx.Expr(0)),
	//		Right:    v.visitExpr(ctx.Expr(1)),
	//		Operator: tree.ComparisonOperatorMatch,
	//	}
	//case ctx.REGEXP_() != nil:
	//	e := &tree.ExpressionBinaryComparison{
	//		Left:     v.visitExpr(ctx.Expr(0)),
	//		Right:    v.visitExpr(ctx.Expr(1)),
	//		Operator: tree.ComparisonOperatorRegexp,
	//	}
	//case ctx.Function_name() != nil:
	case ctx.COLLATE_() != nil:
		// collation_name is any_name
		collationName := ctx.Collation_name().GetText()
		var collationType tree.CollationType
		switch strings.ToLower(collationName) {
		case "binary":
			collationType = tree.CollationTypeBinary
		case "nocase":
			collationType = tree.CollationTypeNoCase
		case "rtrim":
			collationType = tree.CollationTypeRTrim
		default:
			panic(fmt.Sprintf("unknown collation type %s", collationName))
		}
		return &tree.ExpressionCollate{
			Expression: v.visitExpr(ctx.Expr(0)),
			Collation:  collationType,
		}
	case ctx.LIKE_() != nil:
		expr := &tree.ExpressionStringCompare{
			Left:     v.visitExpr(ctx.Expr(0)),
			Operator: tree.StringOperatorLike,
			Right:    v.visitExpr(ctx.Expr(1)),
			Escape:   nil,
		}
		if ctx.NOT_() != nil {
			expr.Operator = tree.StringOperatorNotLike
		}
		return expr
	case ctx.REGEXP_() != nil:
		expr := &tree.ExpressionStringCompare{
			Left:     v.visitExpr(ctx.Expr(0)),
			Operator: tree.StringOperatorRegexp,
			Right:    v.visitExpr(ctx.Expr(1)),
			Escape:   nil,
		}
		if ctx.NOT_() != nil {
			expr.Operator = tree.StringOperatorNotRegexp
		}

		return expr
	case ctx.MATCH_() != nil:
		expr := &tree.ExpressionStringCompare{
			Left:     v.visitExpr(ctx.Expr(0)),
			Operator: tree.StringOperatorMatch,
			Right:    v.visitExpr(ctx.Expr(1)),
			Escape:   nil,
		}
		if ctx.NOT_() != nil {
			expr.Operator = tree.StringOperatorNotMatch
		}
		return expr
	//case ctx.GLOB_() != nil:
	//	expr := &tree.ExpressionStringCompare{
	//		Left:     v.visitExpr(ctx.Expr(0)),
	//		Operator: tree.StringOperatorGlob,
	//		Right:    v.visitExpr(ctx.Expr(1)),
	//		Escape:   nil,
	//	}
	//	if ctx.NOT_() != nil {
	//		expr.Operator = tree.StringOperatorNotGlob
	//	}
	//	return expr
	case ctx.ISNULL_() != nil:
		return &tree.ExpressionIsNull{
			Expression: v.visitExpr(ctx.Expr(0)),
			IsNull:     true,
		}
	//case ctx.NOTNULL_() != nil:
	//case ctx.NULL_() != nil:
	//	if ctx.NOT_() != nil {}
	//case ctx.IS_() != nil and ctx.DISTINCT_()
	case ctx.BETWEEN_() != nil:
		expr := &tree.ExpressionBetween{
			Expression: v.visitExpr(ctx.Expr(0)),
			NotBetween: false,
			Left:       v.visitExpr(ctx.Expr(1)),
			Right:      v.visitExpr(ctx.Expr(2)),
		}
		if ctx.NOT_() != nil {
			expr.NotBetween = true
		}
		return expr
	//case ctx.IN_() != nil && ctx.OPEN_PAR() != nil:
	//	expr := &tree.ExpressionIn{
	//		Expression: v.visitExpr(ctx.Expr(0)),
	//		NotIn:      false,
	//	}
	//	if ctx.NOT_() != nil {
	//		expr.NotIn = true
	//	}
	//	if ctx.Select_stmt() != nil {
	//		expr.InExpressions = v.visitSelectStmt(ctx.Select_stmt())
	//	} else {
	//		expr.InExpressions = v.visitExprList(ctx.AllExpr()[1:])
	//	}
	case ctx.Select_stmt_core() != nil:
		stmt := v.Visit(ctx.Select_stmt_core()).(*tree.SelectStmt)
		expr := &tree.ExpressionSelect{
			IsNot:    false,
			IsExists: false,
			Select:   stmt,
		}
		if ctx.NOT_() != nil {
			expr.IsNot = true
		}
		if ctx.EXISTS_() != nil {
			expr.IsExists = true
		}
		return expr
		//case ctx.CASE_() != nil:
		//	exprCount := len(ctx.AllExpr())
		//	if exprCount == 4 {
		//		return &tree.ExpressionCase{
		//			CaseExpression: v.visitExpr(ctx.Expr(0)),
		//			//WhenThenPairs: v.visitExprList(ctx.AllExpr()[1:3]),
		//			ElseExpression: v.visitExpr(ctx.Expr(3)),
		//		}
		//	}

		// expr list

	}

	return nil
}

func (v *KFSqliteVisitor) visitValuesClause(ctx sqlite.IValues_clauseContext) [][]tree.Expression {
	if ctx == nil {
		return nil
	}

	allValueRowCtx := ctx.AllValue_row()
	rows := make([][]tree.Expression, len(allValueRowCtx))
	for i, valueRowCtx := range allValueRowCtx {
		allExprCtx := valueRowCtx.AllExpr()
		exprs := make([]tree.Expression, len(allExprCtx))
		for j, exprCtx := range allExprCtx {
			exprs[j] = v.Visit(exprCtx).(tree.Expression)
		}
		rows[i] = exprs
	}
	return rows
}

func (v *KFSqliteVisitor) VisitValues_clause(ctx *sqlite.Values_clauseContext) interface{} {
	if ctx == nil {
		return nil
	}

	return v.visitValuesClause(ctx)
}

// VisitUpsert_clause is called when visiting a upsert_clause, return *tree.Upsert
func (v *KFSqliteVisitor) VisitUpsert_clause(ctx *sqlite.Upsert_clauseContext) interface{} {
	clause := tree.Upsert{
		Type: tree.UpsertTypeDoNothing,
	}

	// conflict target
	conflictTarget := tree.ConflictTarget{}
	allIndexedColumnCtx := ctx.AllIndexed_column()
	indexedColumns := make([]string, len(allIndexedColumnCtx))
	for i, indexedColumnCtx := range allIndexedColumnCtx {
		indexedColumns[i] = indexedColumnCtx.Column_name().GetText()
	}
	conflictTarget.IndexedColumns = indexedColumns

	if ctx.GetTarget_expr() != nil {
		conflictTarget.Where = v.visitExpr(ctx.GetTarget_expr())
	}

	if len(allIndexedColumnCtx) != 0 {
		clause.ConflictTarget = &conflictTarget
	}

	if ctx.NOTHING_() != nil {
		return &clause
	}

	// conflict update
	clause.Type = tree.UpsertTypeDoUpdate
	updateCount := len(ctx.AllUpsert_update())
	updates := make([]*tree.UpdateSetClause, updateCount)
	for i, updateCtx := range ctx.AllUpsert_update() {
		updates[i] = v.Visit(updateCtx).(*tree.UpdateSetClause)
	}

	clause.Updates = updates

	if ctx.GetUpdate_expr() != nil {
		clause.Where = v.visitExpr(ctx.GetUpdate_expr())
	}
	return &clause
}

// VisitUpsert_update is called when visiting a upsert_update, return *tree.UpdateSetClause
func (v *KFSqliteVisitor) VisitUpsert_update(ctx *sqlite.Upsert_updateContext) interface{} {
	clause := tree.UpdateSetClause{}
	if ctx.Column_name_list() != nil {
		clause.Columns = v.Visit(ctx.Column_name_list()).([]string)
	} else {
		clause.Columns = []string{ctx.Column_name().GetText()}
	}

	clause.Expression = v.Visit(ctx.Expr()).(tree.Expression)
	return &clause
}

func (v *KFSqliteVisitor) VisitColumn_name_list(ctx *sqlite.Column_name_listContext) interface{} {
	names := make([]string, len(ctx.AllColumn_name()))
	for i, nameCtx := range ctx.AllColumn_name() {
		names[i] = nameCtx.GetText()
	}
	return names
}

// VisitColumn_name is called when visiting a column_name, return string
func (v *KFSqliteVisitor) VisitColumn_name(ctx *sqlite.Column_nameContext) interface{} {
	return ctx.GetText()
}

// VisitReturning_clause is called when visiting a returning_clause, return *tree.ReturningClause
func (v *KFSqliteVisitor) VisitReturning_clause(ctx *sqlite.Returning_clauseContext) interface{} {
	clause := tree.ReturningClause{}
	return &clause
}

func (v *KFSqliteVisitor) VisitInsert_stmt(ctx *sqlite.Insert_stmtContext) interface{} {
	t := tree.Insert{}
	var insertStmt tree.InsertStmt

	if ctx.Common_table_stmt() != nil {
		t.CTE = v.Visit(ctx.Common_table_stmt()).([]*tree.CTE)
	}

	insertStmt.InsertType = getInsertType(ctx)
	insertStmt.Table = ctx.Table_name().GetText()
	if ctx.Table_alias() != nil {
		insertStmt.TableAlias = ctx.Table_alias().GetText()
	}

	allColumnNameCtx := ctx.AllColumn_name()
	if len(allColumnNameCtx) > 0 {
		insertStmt.Columns = make([]string, len(allColumnNameCtx))
		for i, nc := range allColumnNameCtx {
			insertStmt.Columns[i] = nc.GetText()
		}
	}

	insertStmt.Values = v.Visit(ctx.Values_clause()).([][]tree.Expression)
	if ctx.Upsert_clause() != nil {
		insertStmt.Upsert = v.Visit(ctx.Upsert_clause()).(*tree.Upsert)
	}
	if ctx.Returning_clause() != nil {
		insertStmt.ReturningClause = v.Visit(ctx.Returning_clause()).(*tree.ReturningClause)
	}

	t.InsertStmt = &insertStmt
	return &t
}

//// VisitCompound_operator is called when visiting a compound_operator, return *tree.CompoundOperator
//func (v *KFSqliteVisitor) VisitCompound_operator(ctx *sqlite.Compound_operatorContext) interface{} {
//
//}
//
//// VisitOrder_by_stmt is called when visiting a order_by_stmt, return *tree.OrderBy
//func (v *KFSqliteVisitor) VisitOrder_by_stmt(ctx *sqlite.Order_by_stmtContext) interface{} {
//
//}
//
//// VisitLimit_stmt is called when visiting a limit_stmt, return *tree.Limit
//func (v *KFSqliteVisitor) VisitLimit_stmt(ctx *sqlite.Limit_stmtContext) interface{} {
//
//}

// VisitTable_or_subquery is called when visiting a table_or_subquery, return tree.TableOrSubquery
func (v *KFSqliteVisitor) VisitTable_or_subquery(ctx *sqlite.Table_or_subqueryContext) interface{} {
	switch {
	case ctx.Table_name() != nil:
		t := tree.TableOrSubqueryTable{
			Name: ctx.Table_name().GetText(),
		}
		if ctx.Table_alias() != nil {
			t.Alias = ctx.Table_alias().GetText()
		}
		return &t
	case len(ctx.AllTable_or_subquery()) != 0:
		ts := make([]tree.TableOrSubquery, len(ctx.AllTable_or_subquery()))
		for i, subCtx := range ctx.AllTable_or_subquery() {
			ts[i] = v.Visit(subCtx).(tree.TableOrSubquery)
		}
		return &tree.TableOrSubqueryList{
			TableOrSubqueries: ts,
		}
	case ctx.Select_stmt_core() != nil:
		t := tree.TableOrSubquerySelect{
			Select: v.Visit(ctx.Select_stmt_core()).(*tree.SelectStmt),
		}
		if ctx.Table_alias() != nil {
			t.Alias = ctx.Table_alias().GetText()
		}
		return &t
	}
	return nil
}

// VisitJoin_operator is called when visiting a join_operator, return *tree.JoinOperator
func (v *KFSqliteVisitor) VisitJoin_operator(ctx *sqlite.Join_operatorContext) interface{} {
	jp := tree.JoinOperator{}

	if ctx.NATURAL_() != nil {
		jp.Natural = true
	}
	if ctx.OUTER_() != nil {
		jp.Outer = true
	}

	if ctx.INNER_() != nil {
		jp.JoinType = tree.JoinTypeInner
		return &jp
	}

	switch {
	case ctx.LEFT_() != nil:
		jp.JoinType = tree.JoinTypeLeft
	case ctx.RIGHT_() != nil:
		jp.JoinType = tree.JoinTypeRight
	case ctx.FULL_() != nil:
		jp.JoinType = tree.JoinTypeFull
	}

	return &jp
}

// VisitJoin_clause is called when visiting a join_clause, return *tree.JoinClause
func (v *KFSqliteVisitor) VisitJoin_clause(ctx *sqlite.Join_clauseContext) interface{} {
	clause := tree.JoinClause{}

	// just table_or_subquery
	if len(ctx.AllTable_or_subquery()) == 1 {
		clause.TableOrSubquery = v.Visit(ctx.Table_or_subquery(0)).(tree.TableOrSubquery)
		return &clause
	}

	// with joins
	joins := make([]*tree.JoinPredicate, len(ctx.AllJoin_operator()))
	for i, subCtx := range ctx.AllJoin_operator() {
		jp := tree.JoinPredicate{}
		jp.JoinOperator = v.Visit(subCtx).(*tree.JoinOperator)
		jp.Table = v.Visit(ctx.Table_or_subquery(i + 1)).(tree.TableOrSubquery)
		jp.Constraint = v.Visit(ctx.Join_constraint(i).Expr()).(tree.Expression)
		joins[i] = &jp
	}
	clause.Joins = joins

	return &clause
}

// VisitSelect_core is called when visiting a select_core, return *tree.SelectCore
func (v *KFSqliteVisitor) VisitSelect_core(ctx *sqlite.Select_coreContext) interface{} {
	t := tree.SelectCore{
		SelectType: tree.SelectTypeAll,
	}

	if ctx.DISTINCT_() != nil {
		t.SelectType = tree.SelectTypeDistinct
	}

	//NOTE: Columns will be changed in SelectCore
	//assume all columns are * or table.* or table.column
	t.Columns = make([]string, len(ctx.AllResult_column()))
	for i, columnCtx := range ctx.AllResult_column() {
		t.Columns[i] = columnCtx.GetText()
	}

	if ctx.FROM_() != nil {
		fromClause := tree.FromClause{
			JoinClause: &tree.JoinClause{},
		}

		if ctx.Join_clause() != nil {
			fromClause.JoinClause = v.Visit(ctx.Join_clause()).(*tree.JoinClause)
		} else {
			// table_or_subquery list
			tos := make([]tree.TableOrSubquery, len(ctx.AllTable_or_subquery()))
			for i, tableOrSubqueryCtx := range ctx.AllTable_or_subquery() {
				tos[i] = v.Visit(tableOrSubqueryCtx).(tree.TableOrSubquery)
			}

			fromClause.JoinClause.TableOrSubquery = &tree.TableOrSubqueryList{
				TableOrSubqueries: tos,
			}
		}

		t.From = &fromClause
	}

	if ctx.GetWhereExpr() != nil {
		t.Where = v.Visit(ctx.GetWhereExpr()).(tree.Expression)
	}

	if ctx.GROUP_() != nil {
		exprs := make([]tree.Expression, len(ctx.GetGroupByExpr()))
		for i, exprCtx := range ctx.GetGroupByExpr() {
			exprs[i] = v.Visit(exprCtx).(tree.Expression)
		}

		groupBy := &tree.GroupBy{
			Expressions: exprs,
		}

		if ctx.HAVING_() != nil {
			groupBy.Having = v.Visit(ctx.GetHavingExpr()).(tree.Expression)
		}

		t.GroupBy = groupBy
	}

	return &t
}

// VisitSelect_stmt_core is called when visiting a select_stmt_core, return *tree.SelectStmt
func (v *KFSqliteVisitor) VisitSelect_stmt_core(ctx *sqlite.Select_stmt_coreContext) interface{} {
	t := tree.SelectStmt{}
	selectCores := make([]*tree.SelectCore, len(ctx.AllSelect_core()))

	// first select_core
	selectCores[0] = v.Visit(ctx.Select_core(0)).(*tree.SelectCore)

	// rest select_core
	for i, selectCoreCtx := range ctx.AllSelect_core()[1:] {
		compoundOp := v.Visit(ctx.Compound_operator(i)).(tree.CompoundOperatorType)
		core := v.Visit(selectCoreCtx).(*tree.SelectCore)
		core.Compound = &tree.CompoundOperator{Operator: compoundOp}
		selectCores[i+1] = core
	}

	t.SelectCores = selectCores

	if ctx.Order_by_stmt() != nil {
		t.OrderBy = v.Visit(ctx.Order_by_stmt()).(*tree.OrderBy)
	}

	if ctx.Limit_stmt() != nil {
		t.Limit = v.Visit(ctx.Limit_stmt()).(*tree.Limit)
	}

	return &t
}

// VisitSelect_stmt is called when visiting a select_stmt, return *tree.Select
func (v *KFSqliteVisitor) VisitSelect_stmt(ctx *sqlite.Select_stmtContext) interface{} {
	t := tree.Select{}

	if ctx.Common_table_stmt() != nil {
		t.CTE = v.Visit(ctx.Common_table_stmt()).([]*tree.CTE)
	}

	t.SelectStmt = v.Visit(ctx.Select_stmt_core()).(*tree.SelectStmt)
	return &t
}

type KVisitorReturn struct {
	Value interface{}
	Error error
}

func (v *KFSqliteVisitor) VisitSql_stmt_list(ctx *sqlite.Sql_stmt_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *KFSqliteVisitor) VisitSql_stmt(ctx *sqlite.Sql_stmtContext) interface{} {
	// Sql_stmtContext will only have one stmt
	return v.VisitChildren(ctx).(asts)[0]
}

// VisitParse is called first by Visitor.Visit
func (v *KFSqliteVisitor) VisitParse(ctx *sqlite.ParseContext) interface{} {
	// ParseContext will only have one Sql_stmt_listContext
	sqlStmtListContext := ctx.Sql_stmt_list(0)
	// return asts type
	return v.VisitChildren(sqlStmtListContext).(asts)
}

// Visit dispatch to the visit method of the ctx
// e.g. if the tree is a ParseContext, then dispatch call VisitParse.
// Overwrite is needed,
// refer to https://github.com/antlr/antlr4/pull/1841#issuecomment-576791512
func (v *KFSqliteVisitor) Visit(tree antlr.ParseTree) interface{} {
	//if tree == nil {
	//	return nil
	//}
	fmt.Printf("visit tree: %v, %s\n", reflect.TypeOf(tree), tree.GetText())
	return tree.Accept(v)
}

// VisitChildren visits the children of the specified node.
// Overwrite is needed,
// refer to https://github.com/antlr/antlr4/pull/1841#issuecomment-576791512
// calling function need to convert the result to asts
func (v *KFSqliteVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	var result asts
	n := node.GetChildCount()
	for i := 0; i < n; i++ {
		if !v.shouldVisitNextChild(node, result) {
			fmt.Println("should not VisitNextChild")
			break
		}
		child := node.GetChild(i)
		c := child.(antlr.ParseTree)
		childResult := v.Visit(c)
		result = append(result, childResult)
	}
	return result
}

func (v *KFSqliteVisitor) shouldVisitNextChild(node antlr.RuleNode, currentResult interface{}) bool {
	// apply default logic
	return true
}

// asts is a slice of ast
type asts []interface{}

func (a asts) ToSQL() (string, error) {
	panic("not implemented")
}