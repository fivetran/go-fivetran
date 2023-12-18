package users_test

import (
    "context"
    "fmt"
    "net/http"
    "testing"

	"github.com/fivetran/go-fivetran/common"
    
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestUserConnectorDeleteServiceDo(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodDelete, "/v1/users/user_id/connectors/connector_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, `{"code": "Success", "message": "Connector membership has been deleted"}`)
			return response, nil
		},
	)

	service := ftClient.NewUserConnectorMembershipDelete().UserId("user_id").ConnectorId("connector_id")

	// act
	response, err := service.Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	assertUserConnectorDeleteResponse(t, response, "Success", "Connector membership has been deleted")

	// Check that the expected interactions with the mock client occurred
	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
}

func TestUserConnectorDeleteServiceDoMissingId(t *testing.T) {
	// Create a test client
	ftClient, _ := testutils.CreateTestClient()

	// Create the UserDeleteService without setting the User ID
	service := ftClient.NewUserConnectorMembershipDelete()

	// Call the Do method to execute the request
	_, err := service.Do(context.Background())

	// Check for expected error
	expectedError := fmt.Errorf("missing required userId")
	testutils.AssertEqual(t, err, expectedError)
}

func TestUserConnectorDeleteServiceDoMissingConnectorId(t *testing.T) {
	// Create a test client
	ftClient, _ := testutils.CreateTestClient()

	service := ftClient.NewUserConnectorMembershipDelete().UserId("user_id")

	// Call the Do method to execute the request
	_, err := service.Do(context.Background())

	// Check for expected error
	expectedError := fmt.Errorf("missing required connectorId")
	testutils.AssertEqual(t, err, expectedError)
}

func assertUserConnectorDeleteResponse(t *testing.T, response common.CommonResponse, code string, massage string) {
	testutils.AssertEqual(t, response.Code, code)
	testutils.AssertEqual(t, response.Message, massage)
}
