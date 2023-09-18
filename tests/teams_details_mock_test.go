package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestTeamDetailsServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/teams/team_id").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareTeamDetailsResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTeamsDetails().
		TeamId("team_id").
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)
	assertTeamDetailsResponse(t, response)
}

func prepareTeamDetailsResponse() string {
	return fmt.Sprintf(`{
  		"code": "Success",
  		"data": {
      		"id": "team_id",
      		"name": "%v",
      		"description": "%v",
      		"role": "%v"
    	}
	}`,
		TEAM_NAME,
		TEAM_DESCRIPTION,
		TEAM_ROLE)
}

func assertTeamDetailsResponse(t *testing.T, response fivetran.TeamsDetailsResponse) {
	assertEqual(t, response.Code, "Success")
	assertEqual(t, response.Data.Id, "team_id")
	assertEqual(t, response.Data.Name, TEAM_NAME)
	assertEqual(t, response.Data.Description, TEAM_DESCRIPTION)
	assertEqual(t, response.Data.Role, TEAM_ROLE)
}