// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"kwil/x/sql/catalog"
)

func DblinkFuncs0() []*catalog.Function {
	return []*catalog.Function{
		{
			Name: "dblink",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
			},
			ReturnType: &catalog.QualName{Name: "record"},
		},
		{
			Name: "dblink",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "boolean"},
				},
			},
			ReturnType: &catalog.QualName{Name: "record"},
		},
		{
			Name: "dblink",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "text"},
				},
			},
			ReturnType: &catalog.QualName{Name: "record"},
		},
		{
			Name: "dblink",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "boolean"},
				},
			},
			ReturnType: &catalog.QualName{Name: "record"},
		},
		{
			Name: "dblink_build_sql_delete",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "int2vector"},
				},
				{
					Type: &catalog.QualName{Name: "integer"},
				},
				{
					Type: &catalog.QualName{Name: "text[]"},
				},
			},
			ReturnType: &catalog.QualName{Name: "text"},
		},
		{
			Name: "dblink_build_sql_insert",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "int2vector"},
				},
				{
					Type: &catalog.QualName{Name: "integer"},
				},
				{
					Type: &catalog.QualName{Name: "text[]"},
				},
				{
					Type: &catalog.QualName{Name: "text[]"},
				},
			},
			ReturnType: &catalog.QualName{Name: "text"},
		},
		{
			Name: "dblink_build_sql_update",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "int2vector"},
				},
				{
					Type: &catalog.QualName{Name: "integer"},
				},
				{
					Type: &catalog.QualName{Name: "text[]"},
				},
				{
					Type: &catalog.QualName{Name: "text[]"},
				},
			},
			ReturnType: &catalog.QualName{Name: "text"},
		},
		{
			Name: "dblink_cancel_query",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
			},
			ReturnType: &catalog.QualName{Name: "text"},
		},
		{
			Name: "dblink_close",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
			},
			ReturnType: &catalog.QualName{Name: "text"},
		},
		{
			Name: "dblink_close",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "boolean"},
				},
			},
			ReturnType: &catalog.QualName{Name: "text"},
		},
		{
			Name: "dblink_close",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "text"},
				},
			},
			ReturnType: &catalog.QualName{Name: "text"},
		},
		{
			Name: "dblink_close",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "boolean"},
				},
			},
			ReturnType: &catalog.QualName{Name: "text"},
		},
		{
			Name: "dblink_connect",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
			},
			ReturnType: &catalog.QualName{Name: "text"},
		},
		{
			Name: "dblink_connect",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "text"},
				},
			},
			ReturnType: &catalog.QualName{Name: "text"},
		},
		{
			Name: "dblink_connect_u",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
			},
			ReturnType: &catalog.QualName{Name: "text"},
		},
		{
			Name: "dblink_connect_u",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "text"},
				},
			},
			ReturnType: &catalog.QualName{Name: "text"},
		},
		{
			Name:       "dblink_current_query",
			Args:       []*catalog.Argument{},
			ReturnType: &catalog.QualName{Name: "text"},
		},
		{
			Name:       "dblink_disconnect",
			Args:       []*catalog.Argument{},
			ReturnType: &catalog.QualName{Name: "text"},
		},
		{
			Name: "dblink_disconnect",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
			},
			ReturnType: &catalog.QualName{Name: "text"},
		},
		{
			Name: "dblink_error_message",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
			},
			ReturnType: &catalog.QualName{Name: "text"},
		},
		{
			Name: "dblink_exec",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
			},
			ReturnType: &catalog.QualName{Name: "text"},
		},
		{
			Name: "dblink_exec",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "boolean"},
				},
			},
			ReturnType: &catalog.QualName{Name: "text"},
		},
		{
			Name: "dblink_exec",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "text"},
				},
			},
			ReturnType: &catalog.QualName{Name: "text"},
		},
		{
			Name: "dblink_exec",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "boolean"},
				},
			},
			ReturnType: &catalog.QualName{Name: "text"},
		},
		{
			Name: "dblink_fdw_validator",
			Args: []*catalog.Argument{
				{
					Name: "options",
					Type: &catalog.QualName{Name: "text[]"},
				},
				{
					Name: "catalog",
					Type: &catalog.QualName{Name: "oid"},
				},
			},
			ReturnType: &catalog.QualName{Name: "void"},
		},
		{
			Name: "dblink_fetch",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "integer"},
				},
			},
			ReturnType: &catalog.QualName{Name: "record"},
		},
		{
			Name: "dblink_fetch",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "integer"},
				},
				{
					Type: &catalog.QualName{Name: "boolean"},
				},
			},
			ReturnType: &catalog.QualName{Name: "record"},
		},
		{
			Name: "dblink_fetch",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "integer"},
				},
			},
			ReturnType: &catalog.QualName{Name: "record"},
		},
		{
			Name: "dblink_fetch",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "integer"},
				},
				{
					Type: &catalog.QualName{Name: "boolean"},
				},
			},
			ReturnType: &catalog.QualName{Name: "record"},
		},
		{
			Name:       "dblink_get_connections",
			Args:       []*catalog.Argument{},
			ReturnType: &catalog.QualName{Name: "text[]"},
		},
		{
			Name:       "dblink_get_notify",
			Args:       []*catalog.Argument{},
			ReturnType: &catalog.QualName{Name: "record"},
		},
		{
			Name: "dblink_get_notify",
			Args: []*catalog.Argument{
				{
					Name: "conname",
					Type: &catalog.QualName{Name: "text"},
				},
			},
			ReturnType: &catalog.QualName{Name: "record"},
		},
		{
			Name: "dblink_get_pkey",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
			},
			ReturnType: &catalog.QualName{Name: "dblink_pkey_results"},
		},
		{
			Name: "dblink_get_result",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
			},
			ReturnType: &catalog.QualName{Name: "record"},
		},
		{
			Name: "dblink_get_result",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "boolean"},
				},
			},
			ReturnType: &catalog.QualName{Name: "record"},
		},
		{
			Name: "dblink_is_busy",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
			},
			ReturnType: &catalog.QualName{Name: "integer"},
		},
		{
			Name: "dblink_open",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "text"},
				},
			},
			ReturnType: &catalog.QualName{Name: "text"},
		},
		{
			Name: "dblink_open",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "boolean"},
				},
			},
			ReturnType: &catalog.QualName{Name: "text"},
		},
		{
			Name: "dblink_open",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "text"},
				},
			},
			ReturnType: &catalog.QualName{Name: "text"},
		},
		{
			Name: "dblink_open",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "boolean"},
				},
			},
			ReturnType: &catalog.QualName{Name: "text"},
		},
		{
			Name: "dblink_send_query",
			Args: []*catalog.Argument{
				{
					Type: &catalog.QualName{Name: "text"},
				},
				{
					Type: &catalog.QualName{Name: "text"},
				},
			},
			ReturnType: &catalog.QualName{Name: "integer"},
		},
	}
}

func DblinkFuncs() []*catalog.Function {
	funcs := []*catalog.Function{}
	funcs = append(funcs, DblinkFuncs0()...)
	return funcs
}

func Dblink() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = DblinkFuncs()
	return s
}