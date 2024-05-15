package localprocessingagent_test

import (
    "context"
    "fmt"
    "net/http"
    "testing"

	"github.com/fivetran/go-fivetran/common"
    
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestLocalProcessingAgentDeleteServiceDo(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodDelete, "/v1/local-processing-agents/agent_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, `{"code": "Success"}`)
			return response, nil
		},
	)

	service := ftClient.NewLocalProcessingAgentDelete().AgentId("agent_id")

	// act
	response, err := service.Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	assertLocalProcessingAgentDeleteResponse(t, response, "Success")

	// Check that the expected interactions with the mock client occurred
	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
}

func TestLocalProcessingAgentsDeleteServiceDoMissingAgentID(t *testing.T) {
	// Create a test client
	ftClient, _ := testutils.CreateTestClient()

	// Create the ExternalLoggingDeleteService without setting the Log ID
	service := ftClient.NewLocalProcessingAgentDelete()

	// Call the Do method to execute the request
	_, err := service.Do(context.Background())

	// Check for expected error
	expectedError := fmt.Errorf("missing required agentId")
	testutils.AssertEqual(t, err, expectedError)
}

func assertLocalProcessingAgentDeleteResponse(t *testing.T, response common.CommonResponse, code string) {
	testutils.AssertEqual(t, response.Code, code)
}
