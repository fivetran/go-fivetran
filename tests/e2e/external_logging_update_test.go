package fivetran_test

import (
	"context"
	"testing"

	"github.com/fivetran/go-fivetran"
	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewExternalLoggingUpdateE2E(t *testing.T) {
	externalLoggingId := testutils.CreateTempExternalLogging(t)
	details, err := testutils.Client.NewExternalLoggingUpdate().ExternalLoggingId(externalLoggingId).
		Enabled(true).
		Config(fivetran.NewExternalLoggingConfig().
			WorkspaceId("test").
			PrimaryKey("12345678")).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", details)
		t.Error(err)
	}

	testutils.AssertEqual(t, details.Code, "Success")
	testutils.AssertNotEmpty(t, details.Message)
	testutils.AssertEqual(t, details.Data.Id, externalLoggingId)
	testutils.AssertEqual(t, details.Data.Enabled, true)
	testutils.AssertEqual(t, details.Data.Service, "azure_monitor_log")
	testutils.AssertEqual(t, details.Data.Config.WorkspaceId, "test")
	testutils.AssertEqual(t, details.Data.Config.PrimaryKey, "******")
}
