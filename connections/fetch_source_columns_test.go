package connections_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/connections"

	"github.com/fivetran/go-fivetran/tests/mock"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewFetchSourceColumnsServiceMock(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()

	handler := mockClient.When(http.MethodPost, fmt.Sprintf("/v1/connections/%v/schemas/schema_1/fetch-source-columns", testutils.TEST_CONNECTION_ID)).ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareFetchSourceColumnsResponse())
			return response, nil
		})

	svc := ftClient.
		NewFetchSourceColumnsService().
		ConnectionId(testutils.TEST_CONNECTION_ID).
		Schema("schema_1").
		Tables([]string{"table_1", "table_2"})

	// act
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
	assertFetchSourceColumnsResponse(t, response)
}

func assertFetchSourceColumnsResponse(t *testing.T, response connections.MultipleTableColumnsConfigResponse) {
	testutils.AssertEqual(t, response.Code, "Success")

	testutils.AssertEqual(t, len(response.Data.Tables), 2)

	table_1 := response.Data.Tables["table_1"]
	testutils.AssertEqual(t, len(table_1), 2)

	column_1 := table_1["column_1"]
	testutils.AssertEqual(t, *column_1.NameInDestination, "column_1")
	testutils.AssertEqual(t, *column_1.IsPrimaryKey, true)
	testutils.AssertEqual(t, *column_1.Enabled, true)
	testutils.AssertEqual(t, *column_1.Hashed, false)

	column_2 := table_1["column_2"]
	testutils.AssertEqual(t, *column_2.NameInDestination, "column_2")
	testutils.AssertEqual(t, *column_2.IsPrimaryKey, false)
	testutils.AssertEqual(t, *column_2.Enabled, true)

	table_2 := response.Data.Tables["table_2"]
	testutils.AssertEqual(t, len(table_2), 1)

	table_2_column := table_2["column_1"]
	testutils.AssertEqual(t, *table_2_column.NameInDestination, "column_1")
	testutils.AssertEqual(t, *table_2_column.IsPrimaryKey, false)
	testutils.AssertEqual(t, *table_2_column.Enabled, false)
}

func prepareFetchSourceColumnsResponse() string {
	return `{
    "code": "Success",
    "data":{
        "tables": {
            "table_1": {
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
                }
            },
            "table_2": {
                "column_1": {
                    "name_in_destination": "column_1",
                    "enabled": false,
                    "hashed": false,
                    "enabled_patch_settings": {
                        "allowed": true
                    },
                    "is_primary_key" : false
                }
            }
        }
    }
}
`
}
