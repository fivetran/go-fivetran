package tests

import (
    "context"
    "fmt"
    "net/http"
    "testing"

    "github.com/fivetran/go-fivetran"
    "github.com/fivetran/go-fivetran/tests/mock"
)

func TestNewConnectorColumnConfigListServiceMock(t *testing.T) {
    // arrange
    ftClient, mockClient := CreateTestClient()

    handler := mockClient.When(http.MethodGet, fmt.Sprintf("/v1/connectors/%v/schemas/schema_1/tables/table_1/columns", TEST_CONNECTOR_ID)).ThenCall(
        func(req *http.Request) (*http.Response, error) {
            response := mock.NewResponse(req, http.StatusOK, prepareConnectorColumnConfigListResponse())
            return response, nil
        })

    svc := ftClient.
        NewConnectorColumnConfigListService().
        ConnectorId(TEST_CONNECTOR_ID).
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
    assertEqual(t, len(interactions), 1)
    assertEqual(t, interactions[0].Handler, handler)
    assertEqual(t, handler.Interactions, 1)
    assertConnectorColumnConfigListResponse(t, response)
}

func assertConnectorColumnConfigListResponse(t *testing.T, response fivetran.ConnectorColumnConfigListResponse) {
    assertEqual(t, response.Code, "Success")

    assertEqual(t, len(response.Columns), 3)

    column_1 := response.Columns["column_1"]
    assertEqual(t, *column_1.NameInDestination, "column_1")
    assertEqual(t, *column_1.Enabled, true)
    assertEqual(t, *column_1.Hashed, false)
    assertEqual(t, *column_1.EnabledPatchSettings.Allowed, false)
    assertEqual(t, *column_1.EnabledPatchSettings.ReasonCode, "SYSTEM_COLUMN")
    assertEqual(t, *column_1.EnabledPatchSettings.Reason, "The column does not support exclusion as it is a Primary Key")

    column_2 := response.Columns["column_2"]
    assertEqual(t, *column_2.NameInDestination, "column_2")
    assertEqual(t, *column_2.Enabled, true)
    assertEqual(t, *column_2.Hashed, false)
    assertEqual(t, *column_2.EnabledPatchSettings.Allowed, true)

    column_3 := response.Columns["column_3"]
    assertEqual(t, *column_3.NameInDestination, "column_3")
    assertEqual(t, *column_3.Enabled, false)
    assertEqual(t, *column_3.Hashed, false)
    assertEqual(t, *column_3.EnabledPatchSettings.Allowed, true)
}

func prepareConnectorColumnConfigListResponse() string {
    return `{
    "code": "Success",
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
            "enabled": false,
            "hashed": false,
            "enabled_patch_settings": {
                "allowed": true
            }
        }
    }
}


`
}
