package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/connectors"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestNewConnectorTableConfigModifyServiceMock(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()

	handler := mockClient.When(http.MethodPatch, fmt.Sprintf("/v1/connectors/%v/schemas/schema_1/tables/table_1", TEST_CONNECTOR_ID)).ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := RequestBodyToJson(t, req)
			assertConnectorTableConfigModifyRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, prepareConnectorTableConfigModifyResponse())
			return response, nil
		})

	columnName_2 := "column_2"
	columnName_3 := "column_3"

	column_2 := fivetran.NewConnectorSchemaConfigColumn().
		Enabled(true).
		Hashed(false)

	column_3 := fivetran.NewConnectorSchemaConfigColumn().
		Hashed(true)

	svc := ftClient.
		NewConnectorTableConfigModifyService().
		Enabled(true).
		SyncMode("HISTORY").
		ConnectorId(TEST_CONNECTOR_ID).
		Schema("schema_1").
		Table("table_1").
		Columns(columnName_2, column_2).
		Columns(columnName_3, column_3)

	//act
	response, err := svc.
		Do(context.Background())

	// assert
	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)
	assertConnectorTableConfigModifyResponse(t, response)
}

func assertConnectorTableConfigModifyRequest(t *testing.T, request map[string]interface{}) {
	assertKey(t, "enabled", request, true)
	assertKey(t, "sync_mode", request, "HISTORY")

	column2 := request["columns"].(map[string]interface{})["column_2"].(map[string]interface{})

	assertKey(t, "enabled", column2, true)
	assertKey(t, "hashed", column2, false)

	column3 := request["columns"].(map[string]interface{})["column_3"].(map[string]interface{})

	assertKey(t, "hashed", column3, true)
}

func assertConnectorTableConfigModifyResponse(t *testing.T, response connectors.ConnectorSchemaDetailsResponse) {
	assertEqual(t, response.Code, "Success")

	assertEqual(t, len(response.Data.Schemas), 2)
	assertEqual(t, *response.Data.Schemas["schema_1"].Enabled, true)
	assertEqual(t, *response.Data.Schemas["schema_1"].NameInDestination, "schema_1")

	assertEqual(t, len(response.Data.Schemas["schema_1"].Tables), 2)

	table1_schema1 := response.Data.Schemas["schema_1"].Tables["table_1"]
	assertEqual(t, *table1_schema1.Enabled, true)
	assertEqual(t, *table1_schema1.NameInDestination, "table_1")
	assertEqual(t, *table1_schema1.EnabledPatchSettings.Allowed, true)
	assertEqual(t, *table1_schema1.SyncMode, "HISTORY")

	assertEqual(t, len(table1_schema1.Columns), 0)

	table2_schema1 := response.Data.Schemas["schema_1"].Tables["table_2"]
	assertEqual(t, *table2_schema1.Enabled, false)
	assertEqual(t, *table2_schema1.NameInDestination, "table_2")
	assertEqual(t, *table2_schema1.EnabledPatchSettings.Allowed, true)
	assertEqual(t, *table2_schema1.SyncMode, "SOFT_DELETE")

	assertEqual(t, len(table2_schema1.Columns), 3)

	column1_table2_schema1 := table2_schema1.Columns["column_1"]
	assertEqual(t, *column1_table2_schema1.NameInDestination, "column_1")
	assertEqual(t, *column1_table2_schema1.Enabled, true)
	assertEqual(t, *column1_table2_schema1.Hashed, false)
	assertEqual(t, *column1_table2_schema1.EnabledPatchSettings.Allowed, false)
	assertEqual(t, *column1_table2_schema1.EnabledPatchSettings.ReasonCode, "SYSTEM_COLUMN")
	assertEqual(t, *column1_table2_schema1.EnabledPatchSettings.Reason, "The column does not support exclusion as it is a Primary Key")

	column2_table2_schema1 := table2_schema1.Columns["column_2"]
	assertEqual(t, *column2_table2_schema1.NameInDestination, "column_2")
	assertEqual(t, *column2_table2_schema1.Enabled, true)
	assertEqual(t, *column2_table2_schema1.Hashed, false)
	assertEqual(t, *column2_table2_schema1.EnabledPatchSettings.Allowed, true)

	column3_table2_schema1 := table2_schema1.Columns["column_3"]
	assertEqual(t, *column3_table2_schema1.NameInDestination, "column_3")
	assertEqual(t, *column3_table2_schema1.Enabled, true)
	assertEqual(t, *column3_table2_schema1.Hashed, true)
	assertEqual(t, *column3_table2_schema1.EnabledPatchSettings.Allowed, true)

	assertEqual(t, *response.Data.Schemas["schema_2"].Enabled, false)
	assertEqual(t, *response.Data.Schemas["schema_2"].NameInDestination, "schema_2")

	table1_schema2 := response.Data.Schemas["schema_2"].Tables["table_1"]
	assertEqual(t, *table1_schema2.Enabled, true)
	assertEqual(t, *table1_schema2.NameInDestination, "table_1")
	assertEqual(t, *table1_schema2.EnabledPatchSettings.Allowed, true)
	assertEqual(t, *table1_schema2.SyncMode, "SOFT_DELETE")
	assertEqual(t, len(table1_schema2.Columns), 0)

	table2_schema2 := response.Data.Schemas["schema_2"].Tables["table_2"]
	assertEqual(t, *table2_schema2.Enabled, false)
	assertEqual(t, *table2_schema2.NameInDestination, "table_2")
	assertEqual(t, *table2_schema2.EnabledPatchSettings.Allowed, false)
	assertEqual(t, *table2_schema2.EnabledPatchSettings.ReasonCode, "SYSTEM_TABLE")
	assertEqual(t, *table2_schema2.SyncMode, "SOFT_DELETE")
	assertEqual(t, len(table2_schema2.Columns), 0)
}

func prepareConnectorTableConfigModifyResponse() string {
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
                                }
                            },
                            "column_2": {
                                "name_in_destination": "column_2",
                                "enabled": true,
                                "hashed": false,
                                "enabled_patch_settings": {
                                    "allowed": true
                                }
                            },
                            "column_3": {
                                "name_in_destination": "column_3",
                                "enabled": true,
                                "hashed": true,
                                "enabled_patch_settings": {
                                    "allowed": true
                                }
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
