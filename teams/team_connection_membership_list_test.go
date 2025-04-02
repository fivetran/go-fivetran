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

func TestTeamConnectionListServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/teams/team_id/connections").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareTeamConnectionListResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTeamConnectionMembershipsList().
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
	assertTeamConnectionListResponse(t, response)
}

func prepareTeamConnectionListResponse() string {
	return fmt.Sprintf(`{
    		"code": "Success",
    		"data": {
      		"items": [
        		{
          			"id": "connection_id_1",
          			"role": "Connector Administrator",
          			"created_at": "2020-05-25T15:26:47.306509Z"
        		},
        		{
          			"id": "connection_id_2",
          			"role": "Connection Reviewer",
          			"created_at": "2020-05-25T15:26:47.306509Z"
        		}
      		],
      		"next_cursor": null
    		}
		}`)
}

func assertTeamConnectionListResponse(t *testing.T, response teams.TeamConnectionMembershipsListResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.Items[0].ConnectionId, "connection_id_1")
	testutils.AssertEqual(t, response.Data.Items[0].Role, "Connector Administrator")
	testutils.AssertEqual(t, response.Data.Items[0].CreatedAt, "2020-05-25T15:26:47.306509Z")

	testutils.AssertEqual(t, response.Data.Items[1].ConnectionId, "connection_id_2")
	testutils.AssertEqual(t, response.Data.Items[1].Role, "Connection Reviewer")
	testutils.AssertEqual(t, response.Data.Items[1].CreatedAt, "2020-05-25T15:26:47.306509Z")
}