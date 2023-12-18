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

func TestTeamUserListServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/teams/team_id/users").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareTeamUserListResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTeamUserMembershipsList().
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
	assertTeamUserListResponse(t, response)
}

func prepareTeamUserListResponse() string {
	return fmt.Sprintf(`{
    		"code": "Success",
    		"data": {
      		"items": [
        		{
          			"user_id": "user_id_1",
          			"role": "Team Member"
        		},
        		{
          			"user_id": "user_id_2",
          			"role": "Team Manager"
        		}
      		],
      		"next_cursor": null
    		}
		}`)
}

func assertTeamUserListResponse(t *testing.T, response teams.TeamUserMembershipsListResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.Items[0].UserId, "user_id_1")
	testutils.AssertEqual(t, response.Data.Items[0].Role, "Team Member")

	testutils.AssertEqual(t, response.Data.Items[1].UserId, "user_id_2")
	testutils.AssertEqual(t, response.Data.Items[1].Role, "Team Manager")
}