package hybriddeploymentagent_test

import (
    "context"
    "fmt"
    "net/http"
    "testing"

	"github.com/fivetran/go-fivetran/common"
    
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestHybridDeploymentAgentDeleteServiceDo(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodDelete, "/v1/hybrid-deployment-agents/agent_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, `{"code": "Success"}`)
			return response, nil
		},
	)

	service := ftClient.NewHybridDeploymentAgentDelete().AgentId("agent_id")

	// act
	response, err := service.Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	assertHybridDeploymentAgentDeleteResponse(t, response, "Success")

	// Check that the expected interactions with the mock client occurred
	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
}

func TestHybridDeploymentAgentsDeleteServiceDoMissingAgentID(t *testing.T) {
	// Create a test client
	ftClient, _ := testutils.CreateTestClient()

	// Create the ExternalLoggingDeleteService without setting the Log ID
	service := ftClient.NewHybridDeploymentAgentDelete()

	// Call the Do method to execute the request
	_, err := service.Do(context.Background())

	// Check for expected error
	expectedError := fmt.Errorf("missing required agentId")
	testutils.AssertEqual(t, err, expectedError)
}

func assertHybridDeploymentAgentDeleteResponse(t *testing.T, response common.CommonResponse, code string) {
	testutils.AssertEqual(t, response.Code, code)
}
