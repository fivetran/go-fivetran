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

func TestTeamUserDetailsServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/teams/team_id/users/user_id").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareTeamUserDetailsResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTeamUserMembershipDetails().
		TeamId("team_id").
		UserId("user_id").
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertTeamUserDetailsResponse(t, response)
}

func prepareTeamUserDetailsResponse() string {
	return fmt.Sprintf(`{
    		"code": "Success",
    		"data": {
      			"user_id": "user_id",
      			"role": "Team Member"
    		}
		}`)
}

func assertTeamUserDetailsResponse(t *testing.T, response teams.TeamUserMembershipDetailsResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.UserId, "user_id")
	testutils.AssertEqual(t, response.Data.Role, "Team Member")
}