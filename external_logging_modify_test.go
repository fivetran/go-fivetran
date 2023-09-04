package fivetran_test

import (
	"context"
	"testing"

	"github.com/fivetran/go-fivetran"
)

func TestNewExternalLoggingModifyE2E(t *testing.T) {
	externalLoggingId := CreateTempExternalLogging(t)
	details, err := Client.NewExternalLoggingModify().ExternalLoggingId(externalLoggingId).
		Enabled(true).
		Config(fivetran.NewExternalLoggingConfig().
			WorkspaceId("test").
			PrimaryKey("12345678")).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", details)
		t.Error(err)
	}

	AssertEqual(t, details.Code, "Success")
	AssertNotEmpty(t, details.Message)
	AssertEqual(t, details.Data.Id, externalLoggingId)
	AssertEqual(t, details.Data.Enabled, true)
	AssertEqual(t, details.Data.Service, "azure_monitor_log")
	AssertEqual(t, details.Data.Config.WorkspaceId, "test")
    AssertEqual(t, details.Data.Config.PrimaryKey, "******")
}
