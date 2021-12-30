package fivetran_test

import (
	"context"
	"testing"

	"github.com/fivetran/go-fivetran"
)

func TestNewConnectorModifyE2E(t *testing.T) {
	for version, c := range Clients {
		t.Run(version, func(t *testing.T) {
			connectorId := CreateTempConnector(t)

			updated, err := c.NewConnectorModify().ConnectorID(connectorId).
				Paused(true).
				PauseAfterTrial(true).
				//IsHistoricalSync(false).
				SyncFrequency(1440).
				DailySyncTime("03:00").
				Config(fivetran.NewConnectorConfig().
					Username("fivetran_updated").
					Password("fivetran_password_updated").
					TimeframeMonths("SIX")).
				Do(context.Background())

			if err != nil {
				t.Logf("%+v\n", updated)
				t.Error(err)
			}

			AssertEqual(t, updated.Code, "Success")
			AssertEqual(t, updated.Data.ID, connectorId)
			AssertEqual(t, updated.Data.GroupID, PredefinedGroupId)
			AssertEqual(t, updated.Data.Service, "itunes_connect")
			AssertEqual(t, *updated.Data.ServiceVersion, 1)
			AssertEqual(t, updated.Data.Schema, "itunes_e2e_connect")
			AssertEqual(t, updated.Data.ConnectedBy, PredefinedUserId)
			AssertEqual(t, updated.Data.CreatedAt.IsZero(), false)
			AssertEqual(t, updated.Data.SucceededAt.IsZero(), true)
			AssertEqual(t, updated.Data.FailedAt.IsZero(), true)
			AssertEqual(t, *updated.Data.Paused, true)
			AssertEqual(t, *updated.Data.PauseAfterTrial, true)
			AssertEqual(t, *updated.Data.SyncFrequency, 1440)
			AssertEqual(t, updated.Data.ScheduleType, "auto")

			AssertEqual(t, updated.Data.Status.SetupState, "incomplete")
			AssertEqual(t, updated.Data.Status.SyncState, "paused")
			AssertEqual(t, updated.Data.Status.UpdateState, "on_schedule")
			//todo: check after fix
			//AssertEqual(t, *updated.Data.Status.IsHistoricalSync, false)
			AssertHasLength(t, updated.Data.Status.Tasks, 0)
			AssertHasLength(t, updated.Data.Status.Warnings, 0)

			AssertEqual(t, updated.Data.Config.Password, "******")
			AssertEqual(t, updated.Data.Config.TimeframeMonths, "SIX")
			AssertEqual(t, updated.Data.Config.AppSyncMode, "AllApps")
			AssertEqual(t, updated.Data.Config.SalesAccountSyncMode, "AllSalesAccounts")
			AssertEqual(t, updated.Data.Config.FinanceAccountSyncMode, "AllFinanceAccounts")
			AssertEqual(t, updated.Data.Config.Username, "fivetran_updated")
		})
	}
}
