package hybriddeploymentagent_test

import (
    "context"
    "net/http"
    "testing"
		"github.com/fivetran/go-fivetran/hybrid_deployment_agent"
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestHybridDeploymentAgentDetailsServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/hybrid-deployment-agents/agent_id").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareHybridDeploymentAgentDetailsResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewHybridDeploymentAgentDetails().
		AgentId("agent_id").
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertHybridDeploymentAgentDetailsResponse(t, response)
}

func prepareHybridDeploymentAgentDetailsResponse() string {
	return `{
  "code": "Success",
  "data": {
    "id": "id",
    "display_name": "display_name",
    "group_id": "group_id",
    "registered_at": "1970-01-01T00:00:00.000000Z",
    "usage": [
      {
        "connection_id": "connection_id",
        "schema": "schema",
        "service": "service"
      }
    ]
  }
}`
}

func assertHybridDeploymentAgentDetailsResponse(t *testing.T, response hybriddeploymentagent.HybridDeploymentAgentDetailsResponse) {
	testutils.AssertEqual(t, response.Data.Id, "id")
	testutils.AssertEqual(t, response.Data.DisplayName, "display_name")
	testutils.AssertEqual(t, response.Data.GroupId, "group_id")
	testutils.AssertEqual(t, response.Data.RegisteredAt, "1970-01-01T00:00:00.000000Z")
	testutils.AssertEqual(t, response.Data.Usage[0].ConnectionId, "connection_id")
	testutils.AssertEqual(t, response.Data.Usage[0].Schema, "schema")
	testutils.AssertEqual(t, response.Data.Usage[0].Service, "service")
}