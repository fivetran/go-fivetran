package connections_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/connections"

	"github.com/fivetran/go-fivetran/tests/mock"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestConnectionDatabaseSchemaConfigUpdateServiceMock(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()

	handler := mockClient.When(http.MethodPatch, fmt.Sprintf("/v1/connections/%v/schemas/schema_1", testutils.TEST_CONNECTION_ID)).ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertSchemaUpdateRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, prepareSchemaUpdateResponse())
			return response, nil
		})

	tableName_1 := "table_1"
	tableName_2 := "table_2"

	columnName_2 := "column_2"
	columnName_3 := "column_3"

	column_2 := fivetran.NewConnectionSchemaConfigColumn().
		Enabled(true).
		Hashed(false)

	column_3 := fivetran.NewConnectionSchemaConfigColumn().
		Hashed(true)

	table_1 := fivetran.NewConnectionSchemaConfigTable().
		Enabled(true).
		SyncMode("HISTORY")

	table_2 := fivetran.NewConnectionSchemaConfigTable().
		Enabled(false).
		Column(columnName_2, column_2).
		Column(columnName_3, column_3)

	svc := ftClient.
		NewConnectionDatabaseSchemaConfigUpdateService().
		Enabled(true).
		ConnectionId(testutils.TEST_CONNECTION_ID).
		Schema("schema_1").
		Tables(tableName_1, table_1).
		Tables(tableName_2, table_2)

	//act
	response, err := svc.
		Do(context.Background())

	// assert
	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertSchemaUpdateResponse(t, response)
}

func assertSchemaUpdateRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "enabled", request, true)

	table1 := request["tables"].(map[string]interface{})["table_1"].(map[string]interface{})

	testutils.AssertKey(t, "enabled", table1, true)
	testutils.AssertKey(t, "sync_mode", table1, "HISTORY")

	table2 := request["tables"].(map[string]interface{})["table_2"].(map[string]interface{})

	testutils.AssertKey(t, "enabled", table2, false)

	column2 := table2["columns"].(map[string]interface{})["column_2"].(map[string]interface{})

	testutils.AssertKey(t, "enabled", column2, true)
	testutils.AssertKey(t, "hashed", column2, false)

	column3 := table2["columns"].(map[string]interface{})["column_3"].(map[string]interface{})

	testutils.AssertKey(t, "hashed", column3, true)
}

func assertSchemaUpdateResponse(t *testing.T, response connections.ConnectionSchemaDetailsResponse) {
	testutils.AssertEqual(t, response.Code, "Success")

	testutils.AssertEqual(t, len(response.Data.Schemas), 2)
	testutils.AssertEqual(t, *response.Data.Schemas["schema_1"].Enabled, true)
	testutils.AssertEqual(t, *response.Data.Schemas["schema_1"].NameInDestination, "schema_1")

	testutils.AssertEqual(t, len(response.Data.Schemas["schema_1"].Tables), 2)

	table1_schema1 := response.Data.Schemas["schema_1"].Tables["table_1"]
	testutils.AssertEqual(t, *table1_schema1.Enabled, true)
	testutils.AssertEqual(t, *table1_schema1.NameInDestination, "table_1")
	testutils.AssertEqual(t, *table1_schema1.EnabledPatchSettings.Allowed, true)
	testutils.AssertEqual(t, *table1_schema1.SyncMode, "HISTORY")

	testutils.AssertEqual(t, len(table1_schema1.Columns), 0)

	table2_schema1 := response.Data.Schemas["schema_1"].Tables["table_2"]
	testutils.AssertEqual(t, *table2_schema1.Enabled, false)
	testutils.AssertEqual(t, *table2_schema1.NameInDestination, "table_2")
	testutils.AssertEqual(t, *table2_schema1.EnabledPatchSettings.Allowed, true)
	testutils.AssertEqual(t, *table2_schema1.SyncMode, "SOFT_DELETE")

	testutils.AssertEqual(t, len(table2_schema1.Columns), 3)

	column1_table2_schema1 := table2_schema1.Columns["column_1"]
	testutils.AssertEqual(t, *column1_table2_schema1.NameInDestination, "column_1")
	testutils.AssertEqual(t, *column1_table2_schema1.IsPrimaryKey, true)
	testutils.AssertEqual(t, *column1_table2_schema1.Enabled, true)
	testutils.AssertEqual(t, *column1_table2_schema1.Hashed, false)
	testutils.AssertEqual(t, *column1_table2_schema1.EnabledPatchSettings.Allowed, false)
	testutils.AssertEqual(t, *column1_table2_schema1.EnabledPatchSettings.ReasonCode, "SYSTEM_COLUMN")
	testutils.AssertEqual(t, *column1_table2_schema1.EnabledPatchSettings.Reason, "The column does not support exclusion as it is a Primary Key")

	column2_table2_schema1 := table2_schema1.Columns["column_2"]
	testutils.AssertEqual(t, *column2_table2_schema1.NameInDestination, "column_2")
	testutils.AssertEqual(t, *column2_table2_schema1.IsPrimaryKey, false)
	testutils.AssertEqual(t, *column2_table2_schema1.Enabled, true)
	testutils.AssertEqual(t, *column2_table2_schema1.Hashed, false)
	testutils.AssertEqual(t, *column2_table2_schema1.EnabledPatchSettings.Allowed, true)

	column3_table2_schema1 := table2_schema1.Columns["column_3"]
	testutils.AssertEqual(t, *column3_table2_schema1.NameInDestination, "column_3")
	testutils.AssertEqual(t, *column3_table2_schema1.IsPrimaryKey, false)
	testutils.AssertEqual(t, *column3_table2_schema1.Enabled, true)
	testutils.AssertEqual(t, *column3_table2_schema1.Hashed, true)
	testutils.AssertEqual(t, *column3_table2_schema1.EnabledPatchSettings.Allowed, true)

	testutils.AssertEqual(t, *response.Data.Schemas["schema_2"].Enabled, false)
	testutils.AssertEqual(t, *response.Data.Schemas["schema_2"].NameInDestination, "schema_2")

	table1_schema2 := response.Data.Schemas["schema_2"].Tables["table_1"]
	testutils.AssertEqual(t, *table1_schema2.Enabled, true)
	testutils.AssertEqual(t, *table1_schema2.NameInDestination, "table_1")
	testutils.AssertEqual(t, *table1_schema2.EnabledPatchSettings.Allowed, true)
	testutils.AssertEqual(t, *table1_schema2.SyncMode, "SOFT_DELETE")
	testutils.AssertEqual(t, len(table1_schema2.Columns), 0)

	table2_schema2 := response.Data.Schemas["schema_2"].Tables["table_2"]
	testutils.AssertEqual(t, *table2_schema2.Enabled, false)
	testutils.AssertEqual(t, *table2_schema2.NameInDestination, "table_2")
	testutils.AssertEqual(t, *table2_schema2.EnabledPatchSettings.Allowed, false)
	testutils.AssertEqual(t, *table2_schema2.EnabledPatchSettings.ReasonCode, "SYSTEM_TABLE")
	testutils.AssertEqual(t, *table2_schema2.SyncMode, "SOFT_DELETE")
	testutils.AssertEqual(t, len(table2_schema2.Columns), 0)
}

func prepareSchemaUpdateResponse() string {
	return `{
    "code": "Success",
    "data": {
        "enable_new_by_default": true,
        "schema_change_handling": "ALLOW_ALL",
        "schemas": {
            "schema_1": {
                "name_in_destination": "schema_1",
                "enabled": true,
                "tables": {
                    "table_1": {
                        "name_in_destination": "table_1",
                        "enabled": true,
                        "sync_mode": "HISTORY",
                        "enabled_patch_settings": {
                            "allowed": true
                        }
                    },
                    "table_2": {
                        "name_in_destination": "table_2",
                        "enabled": false,
                        "sync_mode": "SOFT_DELETE",
                        "enabled_patch_settings": {
                            "allowed": true
                        },
                        "columns": {
                            "column_1": {
                                "name_in_destination": "column_1",
                                "enabled": true,
                                "hashed": false,
                                "enabled_patch_settings": {
                                    "allowed": false,
                                    "reason_code": "SYSTEM_COLUMN",
                                    "reason": "The column does not support exclusion as it is a Primary Key"
                                },
								"is_primary_key" : true
                            },
                            "column_2": {
                                "name_in_destination": "column_2",
                                "enabled": true,
                                "hashed": false,
                                "enabled_patch_settings": {
                                    "allowed": true
                                },
								"is_primary_key" : false
                            },
                            "column_3": {
                                "name_in_destination": "column_3",
                                "enabled": true,
                                "hashed": true,
                                "enabled_patch_settings": {
                                    "allowed": true
                                },
								"is_primary_key" : false
                            }
                        }
                    }
                }
            },
            "schema_2": {
                "name_in_destination": "schema_2",
                "enabled": false,
                "tables": {
                    "table_1": {
                        "name_in_destination": "table_1",
                        "enabled": true,
                        "sync_mode": "SOFT_DELETE",
                        "enabled_patch_settings": {
                            "allowed": true
                        }
                    },
                    "table_2": {
                        "name_in_destination": "table_2",
                        "enabled": false,
                        "sync_mode": "SOFT_DELETE",
                        "enabled_patch_settings": {
                            "allowed": false,
                            "reason_code": "SYSTEM_TABLE"
                        }
                    }
                }
            }
        }
    }
}
`
}
