package validate

import (
	"kwil/x/sql/sqlerr"
	"kwil/x/sql/sqlparse/ast"
	"kwil/x/sql/sqlparse/astutils"
	"kwil/x/sql/sqlparse/named"
)

// A query can use one (and only one) of the following formats:
// - positional parameters           $1
// - named parameter operator        @param
// - named parameter function calls  sqlc.arg(param)
func ParamStyle(n ast.Node) error {
	namedFunc := astutils.Search(n, named.IsParamFunc)
	for _, f := range namedFunc.Items {
		if fc, ok := f.(*ast.FuncCall); ok {
			args := fc.Args.Items

			if len(args) == 0 {
				continue
			}

			switch val := args[0].(type) {
			case *ast.FuncCall:
				return &sqlerr.Error{
					Code:     "", // TODO: Pick a new error code
					Message:  "Invalid argument to sqlc.arg()",
					Location: val.Location,
				}
			case *ast.ParamRef:
				return &sqlerr.Error{
					Code:     "", // TODO: Pick a new error code
					Message:  "Invalid argument to sqlc.arg()",
					Location: val.Location,
				}
			case *ast.A_Const, *ast.ColumnRef:
			default:
				return &sqlerr.Error{
					Code:    "", // TODO: Pick a new error code
					Message: "Invalid argument to sqlc.arg()",
				}

			}
		}
	}
	return nil
}