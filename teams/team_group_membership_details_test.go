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

func TestTeamGroupDetailsServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/teams/team_id/groups/group_id").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareTeamGroupDetailsResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTeamGroupMembershipDetails().
		TeamId("team_id").
		GroupId("group_id").
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertTeamGroupDetailsResponse(t, response)
}

func prepareTeamGroupDetailsResponse() string {
	return fmt.Sprintf(`{
    		"code": "Success",
    		"data": {
          		"id": "group_id",
          		"role": "Destination Administrator",
          		"created_at": "2020-05-25T15:26:47.306509Z"
    		}
		}`)
}

func assertTeamGroupDetailsResponse(t *testing.T, response teams.TeamGroupMembershipDetailsResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.GroupId, "group_id")
	testutils.AssertEqual(t, response.Data.Role, "Destination Administrator")
	testutils.AssertEqual(t, response.Data.CreatedAt, "2020-05-25T15:26:47.306509Z")
}