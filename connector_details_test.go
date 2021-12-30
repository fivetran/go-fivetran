package fivetran_test

import (
	"context"
	"testing"
)

func TestNewConnectorDetailsIntegration(t *testing.T) {
	for version, c := range Clients {
		t.Run(version, func(t *testing.T) {
			connectorId := CreateTempConnector(t)
			details, err := c.NewConnectorDetails().ConnectorID(connectorId).Do(context.Background())

			if err != nil {
				t.Logf("%+v\n", details)
				t.Error(err)
			}

			AssertEqual(t, details.Code, "Success")
			AssertEqual(t, details.Data.ID, connectorId)
			AssertEqual(t, details.Data.GroupID, PredefinedGroupId)
			AssertEqual(t, details.Data.Service, "itunes_connect")
			AssertEqual(t, *details.Data.ServiceVersion, 1)
			AssertEqual(t, details.Data.Schema, "itunes_e2e_connect")
			AssertEqual(t, details.Data.ConnectedBy, PredefinedUserId)
			AssertEqual(t, details.Data.CreatedAt.IsZero(), false)
			AssertEqual(t, details.Data.SucceededAt.IsZero(), true)
			AssertEqual(t, details.Data.FailedAt.IsZero(), true)
			AssertEqual(t, *details.Data.Paused, false)
			AssertEqual(t, *details.Data.PauseAfterTrial, false)
			AssertEqual(t, *details.Data.SyncFrequency, 360)
			AssertEqual(t, details.Data.ScheduleType, "auto")

			AssertEqual(t, details.Data.Status.SetupState, "incomplete")
			AssertEqual(t, details.Data.Status.SyncState, "scheduled")
			AssertEqual(t, details.Data.Status.UpdateState, "on_schedule")
			AssertEqual(t, *details.Data.Status.IsHistoricalSync, true)
			AssertHasLength(t, details.Data.Status.Tasks, 0)
			AssertHasLength(t, details.Data.Status.Warnings, 0)

			AssertEqual(t, details.Data.Config.Password, "******")
			AssertEqual(t, details.Data.Config.TimeframeMonths, "TWELVE")
			AssertEqual(t, details.Data.Config.AppSyncMode, "AllApps")
			AssertEqual(t, details.Data.Config.SalesAccountSyncMode, "AllSalesAccounts")
			AssertEqual(t, details.Data.Config.FinanceAccountSyncMode, "AllFinanceAccounts")
			AssertEqual(t, details.Data.Config.Username, "fivetran")
		})
	}
}
