package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/common"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestDestinationDeleteServiceDo(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodDelete, "/v1/destinations/destination_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, `{"code": "Success", "message": "Destination deleted"}`)
			return response, nil
		},
	)

	service := ftClient.NewDestinationDelete().DestinationID("destination_id")

	// act
	response, err := service.Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	assertDestinationDeleteResponse(t, response, "Success", "Destination deleted")

	// Check that the expected interactions with the mock client occurred
	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)
}

func TestDestinationDeleteServiceDoMissingDestinationID(t *testing.T) {
	// Create a test client
	ftClient, _ := CreateTestClient()

	// Create the DestinationDeleteService without setting the destination ID
	service := ftClient.NewDestinationDelete()

	// Call the Do method to execute the request
	_, err := service.Do(context.Background())

	// Check for expected error
	expectedError := fmt.Errorf("missing required DestinationID")
	assertEqual(t, err, expectedError)
}

func assertDestinationDeleteResponse(t *testing.T, response common.CommonResponse, code string, massage string) {
	assertEqual(t, response.Code, code)
	assertEqual(t, response.Message, massage)
}
