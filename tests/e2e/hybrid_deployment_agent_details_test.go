package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestHybridDeploymentAgentDetailsE2E(t *testing.T) {
	agentId := testutils.CreateHybridDeploymentAgent(t)

	result, err := testutils.Client.NewHybridDeploymentAgentDetails().AgentId(agentId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertEqual(t, result.Data.Id, agentId)
	testutils.AssertNotEmpty(t, result.Data.DisplayName)
	testutils.AssertNotEmpty(t, result.Data.GroupId)

	t.Cleanup(func() { testutils.DeleteHybridDeploymentAgent(t, agentId) })
}
