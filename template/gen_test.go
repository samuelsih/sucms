package template

import (
	"reflect"
	"testing"

	"github.com/samuelsih/sucms/config"
)

func TestGenerate(t *testing.T) {
	extracted := config.Extracted{
		ExpectedMigrationFilename: []string{"create_users_table.php"},
		MigrationFiles:            []string{"2023_03_23_212051_create_users_table.php"},
		SchemasDefined: config.SchemasDefined{
			"users|create_users_table.php": config.Schema{
				"id":         "bigIncrements",
				"name":       "string",
				"email":      "string",
				"password":   "string",
				"some_uuid":  "uuid",
				"created_at": "date",
				"updated_at": "date",
			},
		},
	}

	expectedTmplData := []TmplData{
		{
			Name: "users|create_users_table.php",
			Table: []string{
				"bigIncrements('id')",
				"string('name')",
				"string('email')",
				"string('password')",
				"uuid('some_uuid')",
				"date('created_at')",
				"date('updated_at')",
			},
		},
	}

	result := Generate(extracted)

	if !reflect.DeepEqual(result, expectedTmplData) {
		t.Errorf("Unexpected result from Generate. Expected: %v, got: %v", expectedTmplData, result)
	}
}
