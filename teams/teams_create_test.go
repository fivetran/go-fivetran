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
	TEAM_NAME        = "Finance Team"
	TEAM_DESCRIPTION = "Finance Team description"
	TEAM_ROLE        = "Account Analyst"
)

func TestNewTeamCreate(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/teams").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertTeamCreateRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareTeamCreateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTeamsCreate().
		Name(TEAM_NAME).
		Description(TEAM_DESCRIPTION).
		Role(TEAM_ROLE).
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

	assertTeamCreateResponse(t, response)
}

func prepareTeamCreateResponse() string {
	return fmt.Sprintf(
		`{
            "code": "Success",
            "message": "Team has been created",
            "data": {
                "id": "clarification_expand",
                "name": "%v",
                "description": "%v",
                "role": "%v"
            }
        }`,
		TEAM_NAME,
		TEAM_DESCRIPTION,
		TEAM_ROLE,
	)
}

func assertTeamCreateRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "name", request, TEAM_NAME)
	testutils.AssertKey(t, "description", request, TEAM_DESCRIPTION)
	testutils.AssertKey(t, "role", request, TEAM_ROLE)
}

func assertTeamCreateResponse(t *testing.T, response teams.TeamsCreateResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)

	testutils.AssertNotEmpty(t, response.Data.Id)
	testutils.AssertEqual(t, response.Data.Name, TEAM_NAME)
	testutils.AssertEqual(t, response.Data.Description, TEAM_DESCRIPTION)
	testutils.AssertEqual(t, response.Data.Role, TEAM_ROLE)
}
