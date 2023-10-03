package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestNewTeamModify(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/teams/clarification_expand").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := RequestBodyToJson(t, req)
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
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)

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
	assertKey(t, "description", request, "Finance Team Updated")
}

func assertTeamModifyResponse(t *testing.T, response fivetran.TeamsModifyResponse) {
	assertEqual(t, response.Code, "Success")
	assertNotEmpty(t, response.Message)

	assertNotEmpty(t, response.Data.Id)
	assertEqual(t, response.Data.Name, TEAM_NAME)
	assertEqual(t, response.Data.Description, "Finance Team Updated")
	assertEqual(t, response.Data.Role, TEAM_ROLE)
}
