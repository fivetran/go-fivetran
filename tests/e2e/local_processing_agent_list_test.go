package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestLocalProcessingAgentListE2E(t *testing.T) {
	agentId := testutils.CreateLocalProcessingAgent(t)

	result, err := testutils.Client.NewLocalProcessingAgentList().Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertEqual(t, result.Data.Items[0].Id, agentId)
	testutils.AssertNotEmpty(t, result.Data.Items[0].DisplayName)
	testutils.AssertNotEmpty(t, result.Data.Items[0].GroupId)

	t.Cleanup(func() { testutils.DeleteLocalProcessingAgent(t, agentId) })
}
