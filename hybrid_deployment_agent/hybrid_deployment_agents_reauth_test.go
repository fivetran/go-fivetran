package hybriddeploymentagent_test

import (
    "context"
    "net/http"
    "testing"
	"github.com/fivetran/go-fivetran/hybrid_deployment_agents"
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewHybridDeploymentAgentReAuthMappingMock(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/hybrid-deployment-agents/agent_id/re-auth").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareHybridDeploymentAgentResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewHybridDeploymentAgentReAuth().
		AgentId("agent_id").
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

	assertHybridDeploymentAgentResponse(t, response)
}

func prepareHybridDeploymentAgentResponse() string {
	return `{
		    "code": "Success",
		    "message": "Success",
		    "data": {
       			"id": "id",
       			"display_name": "display_name",
       			"group_id": "group_id",
       			"registered_at": "1970-01-01T00:00:00.000000Z",
       			"files": {
          			"config_json": "config_json",
          			"auth_json": "auth_json",
          			"docker_compose_yaml": "docker_compose_yaml"
       			}
    		}
		}`
}

func assertHybridDeploymentAgentResponse(t *testing.T, response hybriddeploymentagent.HybridDeploymentAgentCreateResponse) {
	testutils.AssertEqual(t, response.Code, "Success")

	testutils.AssertEqual(t, response.Data.Id, "id")
	testutils.AssertEqual(t, response.Data.DisplayName, "display_name")
	testutils.AssertEqual(t, response.Data.GroupId, "group_id")
	testutils.AssertEqual(t, response.Data.RegisteredAt, "1970-01-01T00:00:00.000000Z")
	testutils.AssertEqual(t, response.Data.Files.ConfigJson, "config_json")
	testutils.AssertEqual(t, response.Data.Files.AuthJson, "auth_json")
	testutils.AssertEqual(t, response.Data.Files.DockerComposeYaml, "docker_compose_yaml")
}
