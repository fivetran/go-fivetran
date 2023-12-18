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
	TEAM_GROUP_ROLE = "Destination Administrator"
)

func TestNewTeamGroupCreate(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/teams/team_id/groups").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertTeamGroupCreateRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareTeamGroupCreateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTeamGroupMembershipCreate().
		TeamId("team_id").
		GroupId("group_id").
		Role(TEAM_GROUP_ROLE).
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

	assertTeamGroupCreateResponse(t, response)
}

func prepareTeamGroupCreateResponse() string {
	return fmt.Sprintf(
		`{
            "code": "Created",
            "message": "Group membership has been created",
            "data": {
                "id": "group_id",
                "role": "%v",
                "created_at": "2021-09-29T10:50:51.397153Z"
            }
        }`,
		TEAM_GROUP_ROLE,
	)
}

func assertTeamGroupCreateRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "id", request, "group_id")
	testutils.AssertKey(t, "role", request, TEAM_GROUP_ROLE)
}

func assertTeamGroupCreateResponse(t *testing.T, response teams.TeamGroupMembershipCreateResponse) {
	testutils.AssertEqual(t, response.Code, "Created")
	testutils.AssertNotEmpty(t, response.Message)
	testutils.AssertEqual(t, response.Data.GroupId, "group_id")
	testutils.AssertEqual(t, response.Data.Role, TEAM_GROUP_ROLE)
	testutils.AssertNotEmpty(t, response.Data.CreatedAt)
}
