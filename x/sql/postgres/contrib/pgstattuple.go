// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"kwil/x/sql/ast"
	"kwil/x/sql/catalog"
)

func PgstattupleFuncs0() []*catalog.Function {
	return []*catalog.Function{
		{
			Name: "pg_relpages",
			Args: []*catalog.Argument{
				{
					Name: "relname",
					Type: &ast.TypeName{Name: "regclass"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pg_relpages",
			Args: []*catalog.Argument{
				{
					Name: "relname",
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "bigint"},
		},
		{
			Name: "pgstatginindex",
			Args: []*catalog.Argument{
				{
					Name: "relname",
					Type: &ast.TypeName{Name: "regclass"},
				},
			},
			ReturnType: &ast.TypeName{Name: "record"},
		},
		{
			Name: "pgstathashindex",
			Args: []*catalog.Argument{
				{
					Name: "relname",
					Type: &ast.TypeName{Name: "regclass"},
				},
			},
			ReturnType: &ast.TypeName{Name: "record"},
		},
		{
			Name: "pgstatindex",
			Args: []*catalog.Argument{
				{
					Name: "relname",
					Type: &ast.TypeName{Name: "regclass"},
				},
			},
			ReturnType: &ast.TypeName{Name: "record"},
		},
		{
			Name: "pgstatindex",
			Args: []*catalog.Argument{
				{
					Name: "relname",
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "record"},
		},
		{
			Name: "pgstattuple",
			Args: []*catalog.Argument{
				{
					Name: "reloid",
					Type: &ast.TypeName{Name: "regclass"},
				},
			},
			ReturnType: &ast.TypeName{Name: "record"},
		},
		{
			Name: "pgstattuple",
			Args: []*catalog.Argument{
				{
					Name: "relname",
					Type: &ast.TypeName{Name: "text"},
				},
			},
			ReturnType: &ast.TypeName{Name: "record"},
		},
		{
			Name: "pgstattuple_approx",
			Args: []*catalog.Argument{
				{
					Name: "reloid",
					Type: &ast.TypeName{Name: "regclass"},
				},
			},
			ReturnType: &ast.TypeName{Name: "record"},
		},
	}
}

func PgstattupleFuncs() []*catalog.Function {
	funcs := []*catalog.Function{}
	funcs = append(funcs, PgstattupleFuncs0()...)
	return funcs
}

func Pgstattuple() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = PgstattupleFuncs()
	return s
}