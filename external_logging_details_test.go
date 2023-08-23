package fivetran_test

import (
    "context"
    "testing"
)

func TestNewExternalLoggingDetailsE2E(t *testing.T) {
    externalLoggingId := CreateTempExternalLogging(t)

    details, err := Client.NewExternalLoggingDetails().ExternalLoggingId(externalLoggingId).Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", details)
        t.Error(err)
    }

    AssertEqual(t, details.Code, "Success")
    AssertEqual(t, details.Data.Id, PredefinedGroupId)
    AssertEqual(t, details.Data.Service, "azure_monitor_log")
    AssertEqual(t, details.Data.Enabled, true)
}
