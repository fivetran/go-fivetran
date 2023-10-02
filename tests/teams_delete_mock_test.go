package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/common"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestTeamDeleteServiceDo(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodDelete, "/v1/teams/team_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, `{"code": "Success", "message": "Team has been deleted"}`)
			return response, nil
		},
	)

	service := ftClient.NewTeamsDelete().TeamId("team_id")

	// act
	response, err := service.Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	assertTeamDeleteResponse(t, response, "Success", "Team has been deleted")

	// Check that the expected interactions with the mock client occurred
	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)
}

func TestTeamDeleteServiceDoMissingId(t *testing.T) {
	// Create a test client
	ftClient, _ := CreateTestClient()

	// Create the TeamDeleteService without setting the Team ID
	service := ftClient.NewTeamsDelete()

	// Call the Do method to execute the request
	_, err := service.Do(context.Background())

	// Check for expected error
	expectedError := fmt.Errorf("missing required teamId")
	assertEqual(t, err, expectedError)
}

func assertTeamDeleteResponse(t *testing.T, response common.CommonResponse, code string, massage string) {
	assertEqual(t, response.Code, code)
	assertEqual(t, response.Message, massage)
}
