package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewHybridDeploymentAgentCreateE2E(t *testing.T) {
	groupId := testutils.CreateGroup(t)
	created, err := testutils.Client.NewHybridDeploymentAgentCreate().
		DisplayName("go_sdk_test_lpa").
		GroupId(groupId).
		EnvType("DOCKER").
		AcceptTerms(true).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertNotEmpty(t, created.Message)
	testutils.AssertNotEmpty(t, created.Data.Id)
	testutils.AssertEqual(t, created.Data.DisplayName, "go_sdk_test_lpa")
	testutils.AssertEqual(t, created.Data.GroupId, groupId)

	t.Cleanup(func() { 
		testutils.DeleteHybridDeploymentAgent(t, created.Data.Id) 
		testutils.DeleteGroup(t, groupId)
	})
}
