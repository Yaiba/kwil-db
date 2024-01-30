package demo

import (
	"context"
	"github.com/kwilteam/kwil-db/internal/engine/types"
	"github.com/kwilteam/kwil-db/parse/sql/tree"
)

// RelationTransformer transforms a Relation to logical plan.
// A Relation represents data, it can be tables or values or subqueries, or
// data composed by joins.
// It'll call
//// It implements the tree.Visitor (AstWalker) interface.
//type RelationTransformer struct {
//}
//
//func NewRelationTransformer() *RelationTransformer {
//	return &RelationTransformer{}
//}
//
//func (t *RelationTransformer) Transform(node *tree.Select) (*LogicalPlan, error) {
//	pb := &Builder{}
//
//	return pb.build(node), nil
//}

type CTEContext struct{}

type BuilderContext struct {
	// can get all schemas & related tables
	info SchemaGetter

	currentSchema *types.Schema
}

func NewBuilderContext(info SchemaGetter, curSchema *types.Schema) *BuilderContext {
	return &BuilderContext{
		info:          info,
		currentSchema: curSchema,
	}
}

// Builder implement tree.AstVisitor, it builds logical plan from a statement
// when visiting the AST.
type Builder struct {
	*tree.BaseAstVisitor

	scope *scope

	ctx *BuilderContext

	// all tables from a schema
	tables map[string]*types.Table

	// all used tables
	usedTables map[string]*types.Table

	metaCollected bool
}

func NewBuilder(ctx *BuilderContext) *Builder {
	return &Builder{
		ctx:           ctx,
		tables:        make(map[string]*types.Table),
		usedTables:    make(map[string]*types.Table),
		metaCollected: false,
	}
}

// collectTables collects all schemas/tables referred from a statement.
func (b *Builder) collectTables(node tree.AstNode) {
	collector := newTableCollector() // need to collect dbName first, then table names
	tbls := collector.collect(node)
	for _, tbl := range tbls {
		b.usedTables[tbl.Name] = b.tables[tbl.Name]
	}
}

// build builds an OperationBuilder from a statement.
func (b *Builder) build(node tree.AstNode) LogicalPlan {
	if b.scope == nil {
		b.scope = newScope()
	}

	if !b.metaCollected {
		b.collectTables(node)
		b.metaCollected = true
	}

	switch t := node.(type) {
	case *tree.Select:
		r := b.VisitSelect(t)
		return r.(LogicalPlan)

		//return b.VisitSelect(t).(LogicalPlan)
	//case *tree.Insert:
	//	if t.CTE != nil {
	//		for _, cte := range t.CTE {
	//			b.buildCTE(cte)
	//		}
	//	}
	//	return b.buildInsert(t)
	//case *tree.Delete:
	//	if t.CTE != nil {
	//		for _, cte := range t.CTE {
	//			b.buildCTE(cte)
	//		}
	//	}
	//	return b.buildDelete(t)
	//case *tree.Update:
	//	if t.CTE != nil {
	//		for _, cte := range t.CTE {
	//			b.buildCTE(cte)
	//		}
	//	}
	//	return b.buildUpdate(t)
	default:
		panic("unknown statement type")
	}
}

// buildDataSource builds a data source from a statement.
func (b *Builder) buildDataSource(node tree.AstNode) LogicalPlan {
	b.descendScope()
	b.scope.ast = node

	switch t := node.(type) {
	case *tree.TableOrSubqueryTable: // simple table
		return b.visitTableOrSubQueryTable(t, nil).(LogicalPlan)
	case *tree.TableOrSubquerySelect: // subquery in Join
		return b.visitTableOrSubQuerySelect(t).(LogicalPlan)
	case *tree.SelectStmt: // select in CTE
		b.VisitSelectStmt(t)
		//case // values
		//case // join
		return b.VisitSelectStmt(t).(LogicalPlan)
	default:
		panic("unknown data source type")
	}

	return nil
}

func (b *Builder) Visit(node tree.AstNode) any {
	return node.Accept(b)
}

// VisitSelect return a *OperationBuilder
func (b *Builder) VisitSelect(node *tree.Select) any {
	if node.CTE != nil {
		//b.descendScope()
		for _, cte := range node.CTE {
			b.VisitCTE(cte)
		}
	}
	return b.VisitSelectStmt(node.SelectStmt)
}

func (b *Builder) descendScope() {
	b.scope = b.scope.descend()
}

func (b *Builder) ascendScope() {
	b.scope = b.scope.ascend()
}

func (b *Builder) VisitCTE(node *tree.CTE) any {
	//s := b.scope.descend()
	//name := node.Table
	//
	//b.buildDataSource(node.)
	return nil
}

// VisitSelectStmt return a *LogicalPlan
// The hierarchy of the logical operators is:
// limit
// - sort
//   - aggregate/distinct
//   - aggregate/having
//   - aggregate/group
//   - filter/where
//   - scan
func (b *Builder) VisitSelectStmt(node *tree.SelectStmt) any {
	//var builder *OperationBuilder
	var plan LogicalPlan
	if len(node.SelectCores) > 1 {
		// set operation (it's tree.CompoundOperator)
		left := b.visitSelectCore(node.SelectCores[0], node.OrderBy)
		for _, core := range node.SelectCores[1:] {
			right := b.visitSelectCore(core, node.OrderBy)
			plan = NewLogicalSet(left, right, core.Compound.Operator)
			left = plan
		}
	} else {
		plan = b.visitSelectCore(node.SelectCores[0], node.OrderBy)
	}

	plan = b.withOrderByLimit(plan, node)
	return plan
}

// visitSelectCore return a *OperationBuilder
// TODO: the order by columns are needed for later sort operator
func (b *Builder) visitSelectCore(node *tree.SelectCore, orderBy *tree.OrderBy) LogicalPlan {
	//plan := b.build(node.From).root
	//plan := b.VisitFromClause(node.From).(*LogicalPlan).root
	plan := b.buildFrom(node)
	plan = b.buildFilter(plan, node.Where) // where

	// expand * in select list

	if node.GroupBy != nil {
		plan = b.buildAggregate(plan, node.GroupBy, node.Columns) // group by
		plan = b.buildFilter(plan, node.GroupBy.Having)           // having
	}

	// if orderBy , project for order

	plan = b.buildDistinct(plan, node.SelectType, node.Columns) // distinct

	plan = b.buildProject(plan, orderBy, node.Columns) // project

	// done in VisitSelectStmt and VisitTableOrSubQuerySelect
	//plan = b.buildSort()  // order by
	//plan = b.buildLimit() // limit

	return plan
	//
	//return &LogicalPlan{
	//	root: plan,
	//}
}

//// VisitSelectCore return a *OperationBuilder
//func (b *Builder) VisitSelectCore(node *tree.SelectCore) any {
//	//builder := b.build(node.From).root
//	//builder := b.VisitFromClause(node.From).(*LogicalPlan).root
//	builder := b.buildFrom(node)
//	builder = b.buildFilter(builder, node.Where) // where
//
//	// expand * in select list
//
//	if node.GroupBy != nil {
//		builder = b.buildAggregate(builder, node.GroupBy, node.Columns) // group by
//		builder = b.buildFilter(builder, node.GroupBy.Having)           // having
//	}
//
//	// if orderBy , project for order
//
//	builder = b.buildDistinct(builder, node.SelectType, node.Columns) // distinct
//
//	builder = b.buildProject(builder, node.node.Columns) // project
//
//	// done in VisitSelectStmt and VisitTableOrSubQuerySelect
//	//builder = b.buildSort()  // order by
//	//builder = b.buildLimit() // limit
//
//	return builder
//	//
//	//return &LogicalPlan{
//	//	root: builder,
//	//}
//}

func (b *Builder) buildFrom(node *tree.SelectCore) LogicalPlan {
	joinClause := node.From.JoinClause
	// simple relation
	// TODO: change SQL parse rule to make it simpler
	// return b.Visit(node.TableOrSubquery)
	left := b.visitTableOrSubquery(joinClause.TableOrSubquery, node.Columns).(LogicalPlan)

	if len(joinClause.Joins) > 0 {
		var tmpPlan LogicalPlan
		//var tmpBuilder *OperationBuilder
		// join relation, from left to right
		for _, join := range joinClause.Joins {
			if tmpPlan != nil {
				left = tmpPlan
			}
			//if tmpBuilder != nil {
			//	left = tmpBuilder
			//}

			right := b.visitTableOrSubquery(join.Table, node.Columns).(LogicalPlan)
			//right := b.visitTableOrSubquery(join.Table, node.Columns).(*OperationBuilder)

			//leftBuilder := left
			//rightBuilder := right
			//tmpPlan = &LogicalPlan{
			//	root: NewOperationBuilder(nil,
			//		operator.NewLogicalJoinOperator(join.JoinOperator,
			//			join.Constraint.ToSQL()),
			//		leftBuilder,
			//		rightBuilder),
			//}
			//tmpBuilder = NewOperationBuilder(nil,
			//	operator.NewLogicalJoinOperator(join.JoinOperator,
			//		join.Constraint.ToSQL()),
			//	leftBuilder,
			//	rightBuilder)
			tmpPlan = NewLogicalJoin(left, right, join.JoinOperator.JoinType, join.Constraint)
		}

		return tmpPlan
	} else {
		return left
	}
}

func (b *Builder) buildFilter(plan LogicalPlan, node tree.Expression) LogicalPlan {
	if node == nil {
		return plan
	}

	return NewLogicalFilter(plan, node)
}

// buildDistinct add distinct aggregate operator to the OperationBuilder.
// NOTE: distinct only operate on the columns in select list
func (b *Builder) buildDistinct(plan LogicalPlan,
	selectType tree.SelectType, cols []tree.ResultColumn) LogicalPlan {
	if selectType == tree.SelectTypeDistinct {
		//distinctOp := operator.NewLogicalDistinctOperator(cols)
		//return builder.WithNewRoot(distinctOp)
		return NewLogicalDistinct(plan, cols)
	}

	return plan
}

// buildProject add project operator to the OperationBuilder.
// NOTE: project only operate on the columns in select list
func (b *Builder) buildProject(plan LogicalPlan, orderby *tree.OrderBy,
	cols []tree.ResultColumn) LogicalPlan {
	//var newCols []tree.ResultColumn
	//
	//for _, o := range orderby.OrderingTerms {
	//	switch t := o.Expression.(type) {
	//	case *tree.ExpressionColumn:
	//		newCols = append(newCols,
	//			&tree.ResultColumnExpression{Expression: t})
	//	}
	//}
	//
	//for _, col := range cols {
	//	switch t := col.(type) {
	//	case *tree.ResultColumnExpression:
	//		switch et := t.Expression.(type) {
	//		case *tree.ExpressionColumn:
	//			if et.Table == "" {
	//			}
	//		}
	//	}
	//}

	proj := b.unfoldStar(plan, cols)

	plan = NewLogicalProjection(plan, proj)

	return plan
}

func (b *Builder) unfoldStar(plan LogicalPlan, cols []tree.ResultColumn) (
	resultCols []tree.ResultColumnExpression) {
	for _, col := range cols {
		switch t := col.(type) {
		case *tree.ResultColumnExpression:
			resultCols = append(resultCols, *t)
		case *tree.ResultColumnStar, *tree.ResultColumnTable:
			for _, field := range plan.Schema().fields {
				resultCols = append(resultCols,
					tree.ResultColumnExpression{
						Expression: &tree.ExpressionColumn{
							Table:  field.TblName,
							Column: field.ColName,
						},
					},
				)
			}
		}
	}

	return resultCols
}

// getColumnAggregateMap return a map from column to its aggregate functions
// take the following SQL as example: `select a, count(b) from t group by a`
// the map will be {b: count}
func getColumnAggregateMap(cols []tree.ResultColumn) map[*tree.AggregateFunc]*tree.ExpressionColumn {
	// NOTE: this should be done in analyzer?
	aggrColMap := make(map[*tree.AggregateFunc]*tree.ExpressionColumn)

	for _, col := range cols {
		switch ct := col.(type) {
		case *tree.ResultColumnExpression:
			switch et := ct.Expression.(type) {
			case *tree.ExpressionFunction:
				if f, ok := et.Function.(*tree.AggregateFunc); ok {
					for _, input := range et.Inputs {
						if colExpr, ok := input.(*tree.ExpressionColumn); ok {
							aggrColMap[f] = colExpr
						}
					}
				}

				//default:
				// TODO: more edges case to consider:
				// - select a, sum(b) + 2 from t group by a, b
				// recursive call to get the aggregate function
			}
		}
	}

	return aggrColMap
}

// buildAggregate add groupby aggregate operator to the OperationBuilder.
// NOTE: groupby group columns and aggregate functions on the columns
func (b *Builder) buildAggregate(plan LogicalPlan,
	groupBy *tree.GroupBy, cols []tree.ResultColumn) LogicalPlan {
	if groupBy == nil {
		return plan
	}

	//colAggrMap := getColumnAggregateMap(cols)
	colAggrs := make([]*tree.AggregateFunc, 0, len(cols))

	for _, col := range cols {
		switch ct := col.(type) {
		case *tree.ResultColumnExpression:
			switch et := ct.Expression.(type) {
			case *tree.ExpressionFunction:
				if f, ok := et.Function.(*tree.AggregateFunc); ok {
					colAggrs = append(colAggrs, f)
				}
			}
		}
	}

	//op := operator.NewLogicalAggregateOperator(groupBy.Expressions, colAggrMap)
	//return NewOperationBuilder(nil, op, builder)
	return NewLogicalAggregate(plan, groupBy.Expressions, colAggrs)
}

// withOrderByLimit add sort and limit operator to the OperationBuilder.
func (b *Builder) withOrderByLimit(plan LogicalPlan, node *tree.SelectStmt) LogicalPlan {
	if node.OrderBy != nil {
		// add sort
		//var limit, offset tree.Expression
		//if node.Limit != nil {
		//	limit = node.Limit.Expression
		//	offset = node.Limit.Offset
		//}
		//takeNOp := operator.NewLogicalTakeNOperator(node.OrderBy, limit, offset)
		//builder = builder.WithNewRoot(takeNOp)

		plan = NewLogicalTakeN(plan, node.OrderBy.OrderingTerms)
	}

	if node.Limit != nil {
		// NOTE: the tree.Limit will be changed only have limit & offset
		// TODO: make the changes in tree.Limit
		limit := node.Limit.Expression
		offset := node.Limit.Offset
		if node.Limit.SecondExpression != nil {
			offset = node.Limit.Expression
			limit = node.Limit.SecondExpression
		}

		//limitOp := operator.NewLogicalLimitOperator(
		//	limit,
		//	offset,
		//)
		//
		//builder = builder.WithNewRoot(limitOp)
		plan = NewLogicalLimit(plan, limit, offset)
	}

	return plan
}

// handled by buildFrom
//// VisitFromClause return a *LogicalPlan
//func (b *Builder) VisitFromClause(node *tree.FromClause) any {
//	return b.Visit(node.JoinClause)
//}
//
//// VisitJoinClause return a *LogicalPlan
//func (b *Builder) VisitJoinClause(node *tree.JoinClause) any {
//	// simple relation
//	// TODO: change SQL parse rule to make it simpler
//	// return b.Visit(node.TableOrSubquery)
//	left := b.visitTableOrSubquery(node.TableOrSubquery).(*LogicalPlan)
//
//	if len(node.Joins) > 0 {
//		var tmpPlan *LogicalPlan
//		// join relation, from left to right
//		for _, join := range node.Joins {
//			if tmpPlan != nil {
//				left = tmpPlan
//			}
//
//			right := b.visitTableOrSubquery(join.Table).(*LogicalPlan)
//			leftBuilder := left.root
//			rightBuilder := right.root
//			tmpPlan = &LogicalPlan{
//				root: NewOperationBuilder(nil,
//					operator.NewLogicalJoinOperator(join.JoinOperator,
//						join.Constraint.ToSQL()),
//					leftBuilder,
//					rightBuilder),
//			}
//		}
//
//		return tmpPlan
//	} else {
//		return left
//	}
//}

func (b *Builder) visitTableOrSubquery(node tree.TableOrSubquery, cols []tree.ResultColumn) any {
	switch t := node.(type) {
	case *tree.TableOrSubqueryTable:
		return b.visitTableOrSubQueryTable(t, cols)
	case *tree.TableOrSubquerySelect:
		return b.visitTableOrSubQuerySelect(t)
	default:
		panic("unknown table or subquery type")
	}
}

// visitTableOrSubQueryTable return an *OperationBuilder
func (b *Builder) visitTableOrSubQueryTable(node *tree.TableOrSubqueryTable, cols []tree.ResultColumn) any {

	// 1. build output columns from original table columns
	// 2. build temporary schema from output columns as well,
	// 3. store the output temporary
	// 4. unfold star(extend output columns), if is star, also resolve(replace) the column name to full qualified(
	//fq) name

	// get table
	//b.ctx.info.GetSchema(node.Name)

	//tbl := b.tables[node.Name]
	//
	//s := newSchema(tbl.Columns...)
	//
	//outputColumns := make([]*OutputColumn, 0, len(cols))
	//
	//var outCols []tree.ResultColumn
	//for _, col := range tbl.Columns {
	//	outputColumns = append(outputColumns, &OutputColumn{
	//		OriginalTblName: tbl.Name,
	//		OriginalColName: col.Name,
	//		TblName:         node.Name,
	//		ColName:         col.ToSQL(),
	//		DB:              "",
	//	})
	//}

	//for _, col := range cols {
	//	switch t := col.(type) {
	//	case *tree.ResultColumnExpression:
	//		outCols = append(outCols, t)
	//	default:
	//		// unfold star
	//
	//	}
	//}

	tb, err := b.ctx.info.TableByName(context.TODO(), b.ctx.currentSchema, node.Name)
	if err != nil {
		panic(err)
	}

	fields := make([]*field, 0, len(tb.Columns))
	for _, col := range tb.Columns {
		fields = append(fields, &field{
			OriginalTblName: tb.Name,
			OriginalColName: col.Name,
			TblName:         tb.Name,
			ColName:         col.Name,
			DB:              "",
		})
	}

	s := newSchema(fields...)
	s.tblName = tb.Name
	s.tblAlias = node.Alias

	imDs := &memDataSource{
		schema: s,
	}

	return NewLogicalScan(imDs)

	//return &OperationBuilder{
	//	op: operator.NewLogicalScanOperator(node.Name, node.Alias, cols),
	//}
}

// visitTableOrSubQuerySelect return a LogicalPlan
func (b *Builder) visitTableOrSubQuerySelect(node *tree.TableOrSubquerySelect) any {
	//root := b.build(node.Select)
	//builder := NewOperationBuilder(
	//	root.ctx,
	//	root.op,
	//	root.inputs...)
	////builder = b.withOrderByLimit(builder, node.Select)
	//return builder

	p := b.VisitSelectStmt(node.Select).(LogicalPlan)
	return p
}
