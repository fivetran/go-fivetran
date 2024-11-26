package connectors_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/connectors"

	"github.com/fivetran/go-fivetran/tests/mock"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewConnectorColumnConfigListServiceMock(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()

	handler := mockClient.When(http.MethodGet, fmt.Sprintf("/v1/connectors/%v/schemas/schema_1/tables/table_1/columns", testutils.TEST_CONNECTOR_ID)).ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareConnectorColumnConfigListResponse())
			return response, nil
		})

	svc := ftClient.
		NewConnectorColumnConfigListService().
		ConnectorId(testutils.TEST_CONNECTOR_ID).
		Schema("schema_1").
		Table("table_1")

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
	assertConnectorColumnConfigListResponse(t, response)
}

func assertConnectorColumnConfigListResponse(t *testing.T, response connectors.ConnectorColumnConfigListResponse) {
	testutils.AssertEqual(t, response.Code, "Success")

	testutils.AssertEqual(t, len(response.Data.Columns), 3)

	column_1 := response.Data.Columns["column_1"]
	testutils.AssertEqual(t, *column_1.NameInDestination, "column_1")
	testutils.AssertEqual(t, *column_1.IsPrimaryKey, true)
	testutils.AssertEqual(t, *column_1.Enabled, true)
	testutils.AssertEqual(t, *column_1.Hashed, false)
	testutils.AssertEqual(t, *column_1.EnabledPatchSettings.Allowed, false)
	testutils.AssertEqual(t, *column_1.EnabledPatchSettings.ReasonCode, "SYSTEM_COLUMN")
	testutils.AssertEqual(t, *column_1.EnabledPatchSettings.Reason, "The column does not support exclusion as it is a Primary Key")

	column_2 := response.Data.Columns["column_2"]
	testutils.AssertEqual(t, *column_2.NameInDestination, "column_2")
	testutils.AssertEqual(t, *column_2.IsPrimaryKey, false)
	testutils.AssertEqual(t, *column_2.Enabled, true)
	testutils.AssertEqual(t, *column_2.Hashed, false)
	testutils.AssertEqual(t, *column_2.EnabledPatchSettings.Allowed, true)

	column_3 := response.Data.Columns["column_3"]
	testutils.AssertEqual(t, *column_3.NameInDestination, "column_3")
	testutils.AssertEqual(t, *column_3.IsPrimaryKey, false)
	testutils.AssertEqual(t, *column_3.Enabled, false)
	testutils.AssertEqual(t, *column_3.Hashed, false)
	testutils.AssertEqual(t, *column_3.EnabledPatchSettings.Allowed, true)
}

func prepareConnectorColumnConfigListResponse() string {
	return `{
    "code": "Success",
    "data":{
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
				""is_primary_key" : true
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
                "enabled": false,
                "hashed": false,
                "enabled_patch_settings": {
                    "allowed": true
				},
                "is_primary_key" : false
        }
    }
}


`
}
