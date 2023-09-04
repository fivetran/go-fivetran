package fivetran_test

import (
    "context"
    "testing"

    "github.com/fivetran/go-fivetran"
)

func TestNewExternalLoggingCreateE2E(t *testing.T) {
    created, err := Client.NewExternalLoggingCreate().
        GroupId(PredefinedGroupId).
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

    AssertEqual(t, created.Code, "Success")
    AssertNotEmpty(t, created.Message)
    AssertEqual(t, created.Data.Id, PredefinedGroupId)
    AssertEqual(t, created.Data.Service, "azure_monitor_log")
    AssertEqual(t, created.Data.Enabled, true)
    AssertEqual(t, created.Data.Config.WorkspaceId, "workspace_id")
    AssertEqual(t, created.Data.Config.PrimaryKey, "******")

    t.Cleanup(func() { DeleteExternalLogging(t, PredefinedGroupId) })
}
