package fivetran_test

import (
	"context"
	"testing"

	"github.com/fivetran/go-fivetran"
	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewExternalLoggingCreateE2E(t *testing.T) {
	created, err := testutils.Client.NewExternalLoggingCreate().
		GroupId(testutils.PredefinedGroupId).
		Service("azure_monitor_log").
		Enabled(true).
		Config(fivetran.NewExternalLoggingConfig().
			WorkspaceId("workspace_id").
			PrimaryKey("PASSWORD")).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertNotEmpty(t, created.Message)
	testutils.AssertEqual(t, created.Data.Id, testutils.PredefinedGroupId)
	testutils.AssertEqual(t, created.Data.Service, "azure_monitor_log")
	testutils.AssertEqual(t, created.Data.Enabled, true)
	testutils.AssertEqual(t, created.Data.Config.WorkspaceId, "workspace_id")
	testutils.AssertEqual(t, created.Data.Config.PrimaryKey, "******")

	t.Cleanup(func() { testutils.DeleteExternalLogging(t, testutils.PredefinedGroupId) })
}
