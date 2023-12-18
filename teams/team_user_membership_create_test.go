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
	TEAM_USER_ROLE = "Destination Administrator"
)

func TestNewTeamUserCreate(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/teams/team_id/users").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertTeamUserCreateRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareTeamUserCreateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTeamUserMembershipCreate().
		TeamId("team_id").
		UserId("user_id").
		Role(TEAM_USER_ROLE).
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

	assertTeamUserCreateResponse(t, response)
}

func prepareTeamUserCreateResponse() string {
	return fmt.Sprintf(
		`{
            "code": "Created",
            "message": "User has been added to the team",
            "data": {
                "user_id": "user_id",
                "role": "%v"
            }
        }`,
		TEAM_USER_ROLE,
	)
}

func assertTeamUserCreateRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "user_id", request, "user_id")
	testutils.AssertKey(t, "role", request, TEAM_USER_ROLE)
}

func assertTeamUserCreateResponse(t *testing.T, response teams.TeamUserMembershipCreateResponse) {
	testutils.AssertEqual(t, response.Code, "Created")
	testutils.AssertNotEmpty(t, response.Message)
	testutils.AssertEqual(t, response.Data.UserId, "user_id")
	testutils.AssertEqual(t, response.Data.Role, TEAM_USER_ROLE)
}
