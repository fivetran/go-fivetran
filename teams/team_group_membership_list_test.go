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

func TestTeamGroupListServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/teams/team_id/groups").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareTeamGroupListResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTeamGroupMembershipsList().
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
	assertTeamGroupListResponse(t, response)
}

func prepareTeamGroupListResponse() string {
	return fmt.Sprintf(`{
    		"code": "Success",
    		"data": {
      		"items": [
        		{
          			"id": "group_id_1",
          			"role": "Destination Administrator",
          			"created_at": "2020-05-25T15:26:47.306509Z"
        		},
        		{
          			"id": "group_id_2",
          			"role": "Destination Reviewer",
          			"created_at": "2020-05-25T15:26:47.306509Z"
        		}
      		],
      		"next_cursor": null
    		}
		}`)
}

func assertTeamGroupListResponse(t *testing.T, response teams.TeamGroupMembershipsListResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.Items[0].GroupId, "group_id_1")
	testutils.AssertEqual(t, response.Data.Items[0].Role, "Destination Administrator")
	testutils.AssertEqual(t, response.Data.Items[0].CreatedAt, "2020-05-25T15:26:47.306509Z")

	testutils.AssertEqual(t, response.Data.Items[1].GroupId, "group_id_2")
	testutils.AssertEqual(t, response.Data.Items[1].Role, "Destination Reviewer")
	testutils.AssertEqual(t, response.Data.Items[1].CreatedAt, "2020-05-25T15:26:47.306509Z")
}