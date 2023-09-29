package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/common"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestTeamUserDeleteServiceDo(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodDelete, "/v1/teams/team_id/users/user_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, `{"code": "Success", "message": "User membership has been deleted"}`)
			return response, nil
		},
	)

	service := ftClient.NewTeamUserMembershipDelete().TeamId("team_id").UserId("user_id")

	// act
	response, err := service.Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	assertTeamUserDeleteResponse(t, response, "Success", "User membership has been deleted")

	// Check that the expected interactions with the mock client occurred
	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)
}

func TestTeamUserDeleteServiceDoMissingId(t *testing.T) {
	// Create a test client
	ftClient, _ := CreateTestClient()

	// Create the TeamDeleteService without setting the Team ID
	service := ftClient.NewTeamUserMembershipDelete()

	// Call the Do method to execute the request
	_, err := service.Do(context.Background())

	// Check for expected error
	expectedError := fmt.Errorf("missing required teamId")
	assertEqual(t, err, expectedError)
}

func TestTeamUserDeleteServiceDoMissingConnectorId(t *testing.T) {
	// Create a test client
	ftClient, _ := CreateTestClient()

	service := ftClient.NewTeamUserMembershipDelete().TeamId("team_id")

	// Call the Do method to execute the request
	_, err := service.Do(context.Background())

	// Check for expected error
	expectedError := fmt.Errorf("missing required userId")
	assertEqual(t, err, expectedError)
}

func assertTeamUserDeleteResponse(t *testing.T, response common.CommonResponse, code string, massage string) {
	assertEqual(t, response.Code, code)
	assertEqual(t, response.Message, massage)
}
