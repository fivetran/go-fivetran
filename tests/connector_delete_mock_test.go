package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestConnectorDeleteService(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodDelete, "/v1/connectors/connector_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareConnectorDeleteResponse("Success", "Connector has been deleted"))
			return response, nil
		})

	service := ftClient.NewConnectorDelete().ConnectorID("connector_id")

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

	assertConnectorDeleteResponse(t, response, "Success", "Connector has been deleted")
}

func TestRespStatusConnectorDeleteService(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	mockClient.When(http.MethodDelete, "/v1/connectors/connector_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusNotFound, prepareConnectorDeleteResponse("NotFound_Integration", "Cannot find entity 'Integration' with id"))
			return response, nil
		})

	service := ftClient.NewConnectorDelete().ConnectorID("connector_id")

	// act
	response, err := service.Do(context.Background())

	// assert
	if err != nil {
		assertConnectorDeleteResponse(t, response, "NotFound_Integration", "Cannot find entity 'Integration' with id 'connector_id'")

	} else {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

}

func prepareConnectorDeleteResponse(code string, message string) string {
	var s = "{\"code\": \"" + code + "\" , \"message\": \"" + message + "\"}"
	return s
}

func assertConnectorDeleteResponse(t *testing.T, response fivetran.ConnectorDeleteResponse, expectCode string, expectMessage string) {
	assertEqual(t, response.Code, expectCode)
	if response.Message != expectMessage {
		t.Errorf("expected message '%s', got '%s'", expectMessage, response.Message)
	}
}
