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

func TestUserConnectionDeleteServiceDo(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodDelete, "/v1/users/user_id/connections/connection_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, `{"code": "Success", "message": "Connection membership has been deleted"}`)
			return response, nil
		},
	)

	service := ftClient.NewUserConnectionMembershipDelete().UserId("user_id").ConnectionId("connection_id")

	// act
	response, err := service.Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	assertUserConnectionDeleteResponse(t, response, "Success", "Connection membership has been deleted")

	// Check that the expected interactions with the mock client occurred
	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
}

func TestUserConnectionDeleteServiceDoMissingId(t *testing.T) {
	// Create a test client
	ftClient, _ := testutils.CreateTestClient()

	// Create the UserDeleteService without setting the User ID
	service := ftClient.NewUserConnectionMembershipDelete()

	// Call the Do method to execute the request
	_, err := service.Do(context.Background())

	// Check for expected error
	expectedError := fmt.Errorf("missing required userId")
	testutils.AssertEqual(t, err, expectedError)
}

func TestUserConnectionDeleteServiceDoMissingConnectionId(t *testing.T) {
	// Create a test client
	ftClient, _ := testutils.CreateTestClient()

	service := ftClient.NewUserConnectionMembershipDelete().UserId("user_id")

	// Call the Do method to execute the request
	_, err := service.Do(context.Background())

	// Check for expected error
	expectedError := fmt.Errorf("missing required connectionId")
	testutils.AssertEqual(t, err, expectedError)
}

func assertUserConnectionDeleteResponse(t *testing.T, response common.CommonResponse, code string, massage string) {
	testutils.AssertEqual(t, response.Code, code)
	testutils.AssertEqual(t, response.Message, massage)
}
