package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestConnectorReSyncTableBadRequest(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/connectors/connector_id/schemas/schema/tables/table/resync").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusBadRequest, prepareConnectorReSyncTableResponse("BadRequest", "Invalid request"))
			return response, nil
		})

	service := ftClient.NewConnectorReSyncTable().
		ConnectorID("connector_id").
		Schema("schema").
		Table("table")

	// act
	response, err := service.Do(context.Background())

	if err == nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	// assert
	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)

	assertConnectorReSyncTableResponse(t, response, "BadRequest", "Invalid request")
}

func TestConnectorReSyncTable_WithNilConnectorID(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()

	service := ftClient.NewConnectorReSyncTable().Schema("schema_name").Table("table_name")

	// act
	response, err := service.Do(context.Background())

	// assert
	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 0)
	assertEqual(t, response, fivetran.ConnectorReSyncTableResponse{})
	assertIsNotNil(t, err)
	assertEqual(t, err.Error(), "missing required ConnectorID")
}

func TestConnectorReSyncTableWithNilSchema(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()

	service := ftClient.NewConnectorReSyncTable().ConnectorID("connector_id").Table("table_name")

	// act
	response, err := service.Do(context.Background())

	// assert
	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 0)
	assertEqual(t, response, fivetran.ConnectorReSyncTableResponse{})
	assertIsNotNil(t, err)
	assertEqual(t, err.Error(), "missing required Schema")
}

func TestConnectorReSyncTableWithNilTable(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()

	service := ftClient.NewConnectorReSyncTable().ConnectorID("connector_id").Schema("schema_name")

	// act
	response, err := service.Do(context.Background())

	// assert
	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 0)
	assertEqual(t, response, fivetran.ConnectorReSyncTableResponse{})
	assertIsNotNil(t, err)
	assertEqual(t, err.Error(), "missing required Table")
}

func TestConnectorReSyncTable(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/connectors/connector_id/schemas/schema_name/tables/table_name/resync").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareConnectorReSyncTableResponse("Success", "Table resync initiated"))
			return response, nil
		})

	service := ftClient.NewConnectorReSyncTable().ConnectorID("connector_id").Schema("schema_name").Table("table_name")

	// act
	response, err := service.Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	// assert
	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)

	assertConnectorReSyncTableResponse(t, response, "Success", "Table resync initiated")
}

func prepareConnectorReSyncTableResponse(code string, message string) string {
	return `{
		"code": "` + code + `",
		"message": "` + message + `"
	}`
}

func assertConnectorReSyncTableResponse(t *testing.T, response fivetran.ConnectorReSyncTableResponse, code string, message string) {
	assertEqual(t, response.Code, code)
	if response.Message != message {
		t.Errorf("expected message `%s` , got '%s'", message, response.Message)
	}
}
