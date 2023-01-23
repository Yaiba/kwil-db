package schemabuilder_test

import (
	schemabuilder "kwil/x/execution/sql-builder/schema-builder"
	"kwil/x/execution/validator"
	"kwil/x/types/databases/clean"
	"kwil/x/types/databases/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GenerateDDL(t *testing.T) {
	ddl, err := schemabuilder.GenerateDDL(&mocks.Db1)
	if err != nil {
		t.Errorf("failed to generate ddl: %v", err)
	}

	// validate
	clean.Clean(&mocks.Db1)
	vldtr := validator.Validator{}
	err = vldtr.Validate(&mocks.Db1)
	if err != nil {
		t.Errorf("failed to validate database: %v", err)
	}

	for _, stmt := range mocks.ALL_MOCK_DDL {
		if !assert.Contains(t, ddl, stmt) {
			t.Errorf("missing ddl statement: %v", stmt)
		}
	}
}
