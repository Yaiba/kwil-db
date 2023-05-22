package sqlddlgenerator_test

import (
	"fmt"
	"testing"

	"github.com/kwilteam/kwil-db/pkg/engine/dto"
	sqlitegenerator "github.com/kwilteam/kwil-db/pkg/engine/sqldb/sql-ddl-generator"
)

func Test_Generate(t *testing.T) {
	ddl, err := sqlitegenerator.GenerateDDL(&dto.Table{
		Name: "test",
		Columns: []*dto.Column{
			{
				Name: "id",
				Type: dto.INT,
				Attributes: []*dto.Attribute{
					{
						Type: "primary_key", // testing string case insensitivity
					},
					{
						Type: dto.NOT_NULL,
					},
				},
			},
			{
				Name: "name",
				Type: dto.TEXT,
				Attributes: []*dto.Attribute{
					{
						Type: dto.NOT_NULL,
					},
					{
						Type:  dto.DEFAULT,
						Value: "foo",
					},
				},
			},
		},
		Indexes: []*dto.Index{
			{
				Name:    "test_index",
				Type:    dto.UNIQUE_BTREE,
				Columns: []string{"id", "name"},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(ddl)
}