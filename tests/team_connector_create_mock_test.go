package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

const (
	TEAM_CONNECTOR_ROLE = "Connector Collaborator"
)

func TestNewTeamConnectorCreate(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/teams/team_id/connectors").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := RequestBodyToJson(t, req)
			assertTeamConnectorCreateRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareTeamConnectorCreateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTeamConnectorMembershipCreate().
		TeamId("team_id").
		ConnectorId("connector_id").
		Role(TEAM_CONNECTOR_ROLE).
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

	assertTeamConnectorCreateResponse(t, response)
}

func prepareTeamConnectorCreateResponse() string {
	return fmt.Sprintf(
		`{
            "code": "Created",
            "message": "Connector membership has been created",
            "data": {
                "id": "connector_id",
                "role": "%v",
                "created_at": "2021-09-29T10:50:51.397153Z"
            }
        }`,
		TEAM_CONNECTOR_ROLE,
	)
}

func assertTeamConnectorCreateRequest(t *testing.T, request map[string]interface{}) {
	assertKey(t, "id", request, "connector_id")
	assertKey(t, "role", request, TEAM_CONNECTOR_ROLE)
}

func assertTeamConnectorCreateResponse(t *testing.T, response fivetran.TeamConnectorMembershipCreateResponse) {
	assertEqual(t, response.Code, "Created")
	assertNotEmpty(t, response.Message)
	assertEqual(t, response.Data.ConnectorId, "connector_id")
	assertEqual(t, response.Data.Role, TEAM_CONNECTOR_ROLE)
	assertNotEmpty(t, response.Data.CreatedAt)
}
