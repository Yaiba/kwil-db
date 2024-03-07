package tree

type Ast interface {
	ToSQL() string
	Accepter
}
