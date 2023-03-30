package specifications

import (
	"context"
	"encoding/json"
	"fmt"
	"kwil/pkg/databases"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type userTable struct {
	ID       int32  `json:"id"`
	UserName string `json:"username"`
	Age      int32  `json:"age"`
	Wallet   string `json:"wallet"`
}

type hasuraTable map[string][]userTable
type hasuraResp map[string]hasuraTable

type ExecuteQueryDsl interface {
	// ExecuteQuery executes QUERY to a database
	// @yaiba TODO: owner is not needed?? because user can only execute queries using his private key
	ExecuteQuery(ctx context.Context, dbid string, queryName string, queryInputs []map[string]any) error
	QueryDatabase(ctx context.Context, query string) ([]byte, error)
}

func ExecuteDBInsertSpecification(ctx context.Context, t *testing.T, execute ExecuteQueryDsl) {
	t.Logf("Executing insert query specification")
	// Given a valid database schema
	db := SchemaLoader.Load(t)
	dbID := databases.GenerateSchemaId(db.Owner, db.Name)

	userQueryName := "create_user"
	userQ := userTable{
		ID:       1111,
		UserName: "test_user",
		Age:      22,
		Wallet:   strings.ToLower(db.Owner),
	}

	/*
			{"id": userQ.ID},
		{"username": userQ.UserName},
		{"age": userQ.Age},
		{"address": userQ.Wallet},
	*/

	userQueryInput := []map[string]any{
		map[string]any{
			"id":       userQ.ID,
			"username": userQ.UserName,
			"age":      userQ.Age,
			"address":  userQ.Wallet,
		},
	}

	// TODO test insert post table
	// When i execute query to database
	err := execute.ExecuteQuery(ctx, dbID, userQueryName, userQueryInput)
	assert.NoError(t, err)

	// Then i expect row to be inserted
	query := fmt.Sprintf(`query MyQuery { %s (where: {id: {_eq: %d}}) {id username age wallet degen}}`,
		qualifiedUserTableName, userQ.ID)
	resByte, err := execute.QueryDatabase(ctx, query)
	assert.NoError(t, err)

	var resp hasuraResp
	err = json.Unmarshal(resByte, &resp)
	assert.NoError(t, err)

	data := resp["data"]
	res := data[qualifiedUserTableName][0]
	assert.EqualValues(t, userQ, res)
}
