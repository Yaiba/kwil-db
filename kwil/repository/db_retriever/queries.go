package dbretriever

import (
	"context"
	"fmt"
	anytype "kwil/x/types/data_types/any_type"
	"kwil/x/types/databases"
	"kwil/x/types/databases/convert"
)

func (q *dbRetriever) GetQueries(ctx context.Context, dbid int32) ([]*databases.SQLQuery[anytype.KwilAny], error) {
	queryList, err := q.gen.GetQueries(ctx, dbid)
	if err != nil {
		return nil, fmt.Errorf(`error getting queries for dbid %d: %w`, dbid, err)
	}

	queries := make([]*databases.SQLQuery[anytype.KwilAny], len(queryList))
	for i, query := range queryList {
		var q databases.SQLQuery[[]byte]
		err = q.DecodeGOB(query.Query)
		if err != nil {
			return nil, fmt.Errorf(`error decoding query %d: %w`, query.ID, err)
		}

		// convert bytes to anytype.KwilAny
		qry, err := convert.Bytes.SQLQueryToKwilAny(&q)
		if err != nil {
			return nil, fmt.Errorf(`error converting query %d: %w`, query.ID, err)
		}

		queries[i] = qry
	}

	return queries, nil
}
