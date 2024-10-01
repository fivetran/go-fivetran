package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewHybridDeploymentAgentDeleteE2E(t *testing.T) {
	agentId := testutils.CreateHybridDeploymentAgent(t)

	deleted, err := testutils.Client.NewHybridDeploymentAgentDelete().AgentId(agentId).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}

	testutils.AssertEqual(t, deleted.Code, "Success")
}
