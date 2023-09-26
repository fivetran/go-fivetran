package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestTeamConnectorListServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/teams/team_id/connectors").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareTeamConnectorListResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewTeamConnectorMembershipsList().
		TeamId("team_id").
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)
	assertTeamConnectorListResponse(t, response)
}

func prepareTeamConnectorListResponse() string {
	return fmt.Sprintf(`{
    		"code": "Success",
    		"data": {
      		"items": [
        		{
          			"id": "connector_id_1",
          			"role": "Connector Administrator",
          			"created_at": "2020-05-25T15:26:47.306509Z"
        		},
        		{
          			"id": "connector_id_2",
          			"role": "Connector Reviewer",
          			"created_at": "2020-05-25T15:26:47.306509Z"
        		}
      		],
      		"next_cursor": null
    		}
		}`)
}

func assertTeamConnectorListResponse(t *testing.T, response fivetran.TeamConnectorMembershipsListResponse) {
	assertEqual(t, response.Code, "Success")
	assertEqual(t, response.Data.Items[0].ConnectorId, "connector_id_1")
	assertEqual(t, response.Data.Items[0].Role, "Connector Administrator")
	assertEqual(t, response.Data.Items[0].CreatedAt, "2020-05-25T15:26:47.306509Z")

	assertEqual(t, response.Data.Items[1].ConnectorId, "connector_id_2")
	assertEqual(t, response.Data.Items[1].Role, "Connector Reviewer")
	assertEqual(t, response.Data.Items[1].CreatedAt, "2020-05-25T15:26:47.306509Z")
}