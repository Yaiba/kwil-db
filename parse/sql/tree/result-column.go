package tree

import (
	sqlwriter "github.com/kwilteam/kwil-db/parse/sql/tree/sql-writer"
)

type ResultColumn interface {
	resultColumn()
	ToSQL() string
	Walk(w AstWalker) error
}

type ResultColumnStar struct {
	*BaseAstNode
}

func (r *ResultColumnStar) Accept(v AstVisitor) any {
	return v.VisitResultColumnStar(r)
}

func (r *ResultColumnStar) resultColumn() {}
func (r *ResultColumnStar) ToSQL() string {
	return "*"
}
func (r *ResultColumnStar) Walk(w AstWalker) error {
	return run(
		w.EnterResultColumnStar(r),
		w.ExitResultColumnStar(r),
	)
}

type ResultColumnExpression struct {
	*BaseAstNode

	Expression Expression
	Alias      string
}

func (r *ResultColumnExpression) Accept(v AstVisitor) any {
	return v.VisitResultColumnExpression(r)
}

func (r *ResultColumnExpression) resultColumn() {}
func (r *ResultColumnExpression) ToSQL() string {
	stmt := sqlwriter.NewWriter()
	stmt.WriteString(r.Expression.ToSQL())
	if r.Alias != "" {
		stmt.Token.As()
		stmt.WriteIdent(r.Alias)
	}
	return stmt.String()
}
func (r *ResultColumnExpression) Walk(w AstWalker) error {
	return run(
		w.EnterResultColumnExpression(r),
		accept(w, r.Expression),
		w.ExitResultColumnExpression(r),
	)
}

type ResultColumnTable struct {
	*BaseAstNode

	TableName string
}

func (r *ResultColumnTable) Accept(v AstVisitor) any {
	return v.VisitResultColumnTable(r)
}

func (r *ResultColumnTable) resultColumn() {}
func (r *ResultColumnTable) ToSQL() string {
	stmt := sqlwriter.NewWriter()
	stmt.WriteIdent(r.TableName)
	stmt.Token.Period()
	stmt.Token.Asterisk()
	return stmt.String()
}
func (r *ResultColumnTable) Walk(w AstWalker) error {
	return run(
		w.EnterResultColumnTable(r),
		w.ExitResultColumnTable(r),
	)
}
