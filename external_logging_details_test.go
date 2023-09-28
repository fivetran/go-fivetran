package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewExternalLoggingDetailsE2E(t *testing.T) {
	externalLoggingId := testutils.CreateTempExternalLogging(t)

	details, err := testutils.Client.NewExternalLoggingDetails().ExternalLoggingId(externalLoggingId).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", details)
		t.Error(err)
	}

	testutils.AssertEqual(t, details.Code, "Success")
	testutils.AssertEqual(t, details.Data.Id, testutils.PredefinedGroupId)
	testutils.AssertEqual(t, details.Data.Service, "azure_monitor_log")
	testutils.AssertEqual(t, details.Data.Enabled, true)
	testutils.AssertEqual(t, details.Data.Config.WorkspaceId, "workspace_id")
	testutils.AssertEqual(t, details.Data.Config.PrimaryKey, "******")
}
