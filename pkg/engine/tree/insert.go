package tree

type InsertStatement struct {
	CTEs            []*CTE
	InsertOr        InsertOr
	Table           string
	Alias           *Alias
	Columns         []string
	Expressions     [][]*Expression
	Upsert          *Upsert
	ReturningClause *ReturningClause
}

type InsertOr uint8

const (
	InsertOrNone InsertOr = iota
	InsertOrReplace
)