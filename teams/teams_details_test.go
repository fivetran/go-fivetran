package teams_test

import (
    "context"
    "fmt"
    "net/http"
    "testing"

	"github.com/fivetran/go-fivetran/teams"
    
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestTeamDetailsServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := testutils.CreateTestClient()
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
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
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

func assertTeamDetailsResponse(t *testing.T, response teams.TeamsDetailsResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.Id, "team_id")
	testutils.AssertEqual(t, response.Data.Name, TEAM_NAME)
	testutils.AssertEqual(t, response.Data.Description, TEAM_DESCRIPTION)
	testutils.AssertEqual(t, response.Data.Role, TEAM_ROLE)
}