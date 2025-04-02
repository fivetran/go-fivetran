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

const (
	TEAM_CONNECTION_ROLE = "Connector Collaborator"
)

func TestNewTeamConnectionCreate(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/teams/team_id/connections").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertTeamConnectionCreateRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareTeamConnectionCreateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTeamConnectionMembershipCreate().
		TeamId("team_id").
		ConnectionId("connection_id").
		Role(TEAM_CONNECTION_ROLE).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	// assert
	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)

	assertTeamConnectionCreateResponse(t, response)
}

func prepareTeamConnectionCreateResponse() string {
	return fmt.Sprintf(
		`{
            "code": "Created",
            "message": "Connection membership has been created",
            "data": {
                "id": "connection_id",
                "role": "%v",
                "created_at": "2021-09-29T10:50:51.397153Z"
            }
        }`,
		TEAM_CONNECTION_ROLE,
	)
}

func assertTeamConnectionCreateRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "id", request, "connection_id")
	testutils.AssertKey(t, "role", request, TEAM_CONNECTION_ROLE)
}

func assertTeamConnectionCreateResponse(t *testing.T, response teams.TeamConnectionMembershipCreateResponse) {
	testutils.AssertEqual(t, response.Code, "Created")
	testutils.AssertNotEmpty(t, response.Message)
	testutils.AssertEqual(t, response.Data.ConnectionId, "connection_id")
	testutils.AssertEqual(t, response.Data.Role, TEAM_CONNECTION_ROLE)
	testutils.AssertNotEmpty(t, response.Data.CreatedAt)
}
