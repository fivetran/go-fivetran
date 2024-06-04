package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestLocalProcessingAgentDetailsE2E(t *testing.T) {
	agentId := testutils.CreateLocalProcessingAgent(t)

	result, err := testutils.Client.NewLocalProcessingAgentDetails().AgentId(agentId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertEqual(t, result.Data.Id, agentId)
	testutils.AssertNotEmpty(t, result.Data.DisplayName)
	testutils.AssertNotEmpty(t, result.Data.GroupId)

	t.Cleanup(func() { testutils.DeleteLocalProcessingAgent(t, agentId) })
}
