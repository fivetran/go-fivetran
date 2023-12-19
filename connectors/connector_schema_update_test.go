package connectors_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/connectors"
	
	"github.com/fivetran/go-fivetran/tests/mock"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestConnectorSchemaUpdateFullMappingMock(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPatch, fmt.Sprintf("/v1/connectors/%v/schemas", testutils.TEST_CONNECTOR_ID)).ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertSchemaRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, prepareSchemaResponse())
			return response, nil
		})

	tableName := "table_1"
	columnName := "column_2"
	schemaName := "schema_1"

	column := fivetran.NewConnectorSchemaConfigColumn().
		Enabled(true).
		Hashed(true)

	table := fivetran.NewConnectorSchemaConfigTable().
		Enabled(true).
		SyncMode("HISTORY").
		Column(columnName, column)

	schema := fivetran.NewConnectorSchemaConfigSchema().
		Enabled(true).
		Table(tableName, table)

	svc := ftClient.
		NewConnectorSchemaUpdateService().
		ConnectorID(testutils.TEST_CONNECTOR_ID).
		SchemaChangeHandling("BLOCK_ALL").
		Schema(schemaName, schema)

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
	assertSchemaResponse(t, response)
}

func assertSchemaRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "schema_change_handling", request, "BLOCK_ALL")

	schema := request["schemas"].(map[string]interface{})["schema_1"].(map[string]interface{})

	testutils.AssertKey(t, "enabled", schema, true)

	table := schema["tables"].(map[string]interface{})["table_1"].(map[string]interface{})

	testutils.AssertKey(t, "enabled", table, true)
	testutils.AssertKey(t, "sync_mode", table, "HISTORY")

	column := table["columns"].(map[string]interface{})["column_2"].(map[string]interface{})

	testutils.AssertKey(t, "enabled", column, true)
	testutils.AssertKey(t, "hashed", column, true)
}

func assertSchemaResponse(t *testing.T, response connectors.ConnectorSchemaDetailsResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.SchemaChangeHandling, "BLOCK_ALL")
	testutils.AssertEqual(t, len(response.Data.Schemas), 1)
	testutils.AssertEqual(t, *response.Data.Schemas["schema_1"].Enabled, true)
	testutils.AssertEqual(t, *response.Data.Schemas["schema_1"].NameInDestination, "schema_1")
	testutils.AssertEqual(t, len(response.Data.Schemas["schema_1"].Tables), 2)

	table1 := response.Data.Schemas["schema_1"].Tables["table_1"]
	testutils.AssertEqual(t, *table1.Enabled, true)
	testutils.AssertEqual(t, *table1.NameInDestination, "table_1")
	testutils.AssertEqual(t, *table1.EnabledPatchSettings.Allowed, true)
	testutils.AssertEqual(t, *table1.SyncMode, "HISTORY")
	testutils.AssertEqual(t, len(table1.Columns), 2)

	table2 := response.Data.Schemas["schema_1"].Tables["table_2"]
	testutils.AssertEqual(t, *table2.Enabled, true)
	testutils.AssertEqual(t, *table2.NameInDestination, "table_2")
	testutils.AssertEqual(t, *table2.EnabledPatchSettings.Allowed, false)
	testutils.AssertEqual(t, *table2.EnabledPatchSettings.ReasonCode, "SYSTEM_TABLE")
	testutils.AssertEqual(t, *table2.EnabledPatchSettings.Reason, "The table does not support exclusion")
	testutils.AssertEqual(t, *table2.SyncMode, "SOFT_DELETE")
	testutils.AssertEqual(t, len(table2.Columns), 0)

	column1 := table1.Columns["column_1"]
	testutils.AssertEqual(t, *column1.Enabled, true)
	testutils.AssertEqual(t, *column1.EnabledPatchSettings.Allowed, false)
	testutils.AssertEqual(t, *column1.EnabledPatchSettings.ReasonCode, "SYSTEM_COLUMN")
	testutils.AssertEqual(t, *column1.EnabledPatchSettings.Reason, "The column does not support exclusion as it is a Primary Key")
	testutils.AssertEqual(t, *column1.Hashed, false)
	testutils.AssertEqual(t, *column1.NameInDestination, "column_1")

	column2 := table1.Columns["column_2"]
	testutils.AssertEqual(t, *column2.Enabled, true)
	testutils.AssertEqual(t, *column2.Hashed, true)
	testutils.AssertEqual(t, *column2.EnabledPatchSettings.Allowed, true)
}

func prepareSchemaResponse() string {
	return `{
				"code":"Success",
				"data":{
					"schema_change_handling": "BLOCK_ALL",
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
											"hashed": true,
											"enabled_patch_settings": {
												"allowed": true
											}
										}
									}
								},
								"table_2": {
									"name_in_destination": "table_2",
									"enabled": true,
									"sync_mode": "SOFT_DELETE",
									"enabled_patch_settings": {
										"allowed": false, 
										"reason_code": "SYSTEM_TABLE",
										"reason": "The table does not support exclusion"
									}
								}
							}
						}
					}
				}
			}`
}
