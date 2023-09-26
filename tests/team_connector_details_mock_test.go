package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestTeamConnectorDetailsServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/teams/team_id/connectors/connector_id").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareTeamConnectorDetailsResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTeamConnectorMembershipDetails().
		TeamId("team_id").
		ConnectorId("connector_id").
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)
	assertTeamConnectorDetailsResponse(t, response)
}

func prepareTeamConnectorDetailsResponse() string {
	return fmt.Sprintf(`{
    		"code": "Success",
    		"data": {
          		"id": "connector_id",
          		"role": "Connector Administrator",
          		"created_at": "2020-05-25T15:26:47.306509Z"
    		}
		}`)
}

func assertTeamConnectorDetailsResponse(t *testing.T, response fivetran.TeamConnectorMembershipDetailsResponse) {
	assertEqual(t, response.Code, "Success")
	assertEqual(t, response.Data.ConnectorId, "connector_id")
	assertEqual(t, response.Data.Role, "Connector Administrator")
	assertEqual(t, response.Data.CreatedAt, "2020-05-25T15:26:47.306509Z")
}