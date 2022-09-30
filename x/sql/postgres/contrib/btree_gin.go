// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"kwil/x/sql/ast"
	"kwil/x/sql/catalog"
)

func BtreeGinFuncs0() []*catalog.Function {
	return []*catalog.Function{
		{
			Name: "gin_enum_cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
				{
					Type: &ast.TypeName{Name: "anyenum"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
		{
			Name: "gin_numeric_cmp",
			Args: []*catalog.Argument{
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
				{
					Type: &ast.TypeName{Name: "numeric"},
				},
			},
			ReturnType: &ast.TypeName{Name: "integer"},
		},
	}
}

func BtreeGinFuncs() []*catalog.Function {
	funcs := []*catalog.Function{}
	funcs = append(funcs, BtreeGinFuncs0()...)
	return funcs
}

func BtreeGin() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = BtreeGinFuncs()
	return s
}