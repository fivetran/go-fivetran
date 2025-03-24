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

func TestTeamConnectionDetailsServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/teams/team_id/connections/connection_id").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareTeamConnectionDetailsResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTeamConnectionMembershipDetails().
		TeamId("team_id").
		ConnectionId("connection_id").
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertTeamConnectionDetailsResponse(t, response)
}

func prepareTeamConnectionDetailsResponse() string {
	return fmt.Sprintf(`{
    		"code": "Success",
    		"data": {
          		"id": "connection_id",
          		"role": "Connection Administrator",
          		"created_at": "2020-05-25T15:26:47.306509Z"
    		}
		}`)
}

func assertTeamConnectionDetailsResponse(t *testing.T, response teams.TeamConnectionMembershipDetailsResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.ConnectionId, "connection_id")
	testutils.AssertEqual(t, response.Data.Role, "Connection Administrator")
	testutils.AssertEqual(t, response.Data.CreatedAt, "2020-05-25T15:26:47.306509Z")
}