package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewLocalProcessingAgentDeleteE2E(t *testing.T) {
	agentId := testutils.CreateLocalProcessingAgent(t)

	deleted, err := testutils.Client.NewLocalProcessingAgentDelete().AgentId(agentId).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}

	testutils.AssertEqual(t, deleted.Code, "Success")
}
