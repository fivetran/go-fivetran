package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestTeamDeleteRoleInAccountServiceDo(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodDelete, "/v1/teams/team_id/role").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, `{"code": "Success", "message": "Team role in account has been removed"}`)
			return response, nil
		},
	)

	service := ftClient.NewTeamsDeleteRoleInAccount().TeamId("team_id")

	// act
	response, err := service.Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	assertTeamDeleteRoleInAccountResponse(t, response, "Success", "Team role in account has been removed")

	// Check that the expected interactions with the mock client occurred
	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)
}

func TestTeamDeleteRoleInAccountServiceDoMissingId(t *testing.T) {
	// Create a test client
	ftClient, _ := CreateTestClient()

	// Create the TeamDeleteService without setting the Team ID
	service := ftClient.NewTeamsDeleteRoleInAccount()

	// Call the Do method to execute the request
	_, err := service.Do(context.Background())

	// Check for expected error
	expectedError := fmt.Errorf("missing required teamId")
	assertEqual(t, err, expectedError)
}

func assertTeamDeleteRoleInAccountResponse(t *testing.T, response fivetran.TeamsDeleteRoleInAccountResponse, code string, massage string) {
	assertEqual(t, response.Code, code)
	assertEqual(t, response.Message, massage)
}
