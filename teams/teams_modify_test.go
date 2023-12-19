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

func TestNewTeamModify(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/teams/clarification_expand").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertTeamModifyRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, prepareTeamModifyResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTeamsModify().
		TeamId("clarification_expand").
		Description("Finance Team Updated").
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

	assertTeamModifyResponse(t, response)
}

func prepareTeamModifyResponse() string {
	return fmt.Sprintf(
		`{
            "code": "Success",
            "message": "Team has been updated",
            "data": {
                "id": "clarification_expand",
                "name": "%v",
                "description": "%v",
                "role": "%v"
            }
        }`,
		TEAM_NAME,
		"Finance Team Updated",
		TEAM_ROLE,
	)
}

func assertTeamModifyRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "description", request, "Finance Team Updated")
}

func assertTeamModifyResponse(t *testing.T, response teams.TeamsModifyResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)

	testutils.AssertNotEmpty(t, response.Data.Id)
	testutils.AssertEqual(t, response.Data.Name, TEAM_NAME)
	testutils.AssertEqual(t, response.Data.Description, "Finance Team Updated")
	testutils.AssertEqual(t, response.Data.Role, TEAM_ROLE)
}
