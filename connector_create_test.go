package fivetran_test

import (
	"context"
	"testing"

	"github.com/fivetran/go-fivetran"
)

func TestNewConnectorCreateIntegration(t *testing.T) {
	for version, c := range Clients {
		t.Run(version, func(t *testing.T) {
			created, err := c.NewConnectorCreate().
				GroupID("climbed_consulted").
				Service("itunes_connect").
				RunSetupTests(false).
				Config(fivetran.NewConnectorConfig().
					Schema("itunes_e2e_connect").
					Username("fivetran").
					Password("fivetran-api-e2e")).
				Do(context.Background())

			if err != nil {
				t.Logf("%+v\n", created)
				t.Error(err)
			}

			AssertEqual(t, created.Code, "Success")
			AssertEqual(t, created.Message, "Connector has been created")
			AssertNotEmpty(t, created.Data.ID)
			AssertEqual(t, created.Data.GroupID, "climbed_consulted")
			AssertEqual(t, created.Data.Service, "itunes_connect")
			AssertEqual(t, *created.Data.ServiceVersion, 1)
			AssertEqual(t, created.Data.Schema, "itunes_e2e_connect")
			AssertEqual(t, created.Data.ConnectedBy, "cherry_spoilt")
			AssertEqual(t, created.Data.CreatedAt.IsZero(), false)
			AssertEqual(t, created.Data.SucceededAt.IsZero(), true)
			AssertEqual(t, created.Data.FailedAt.IsZero(), true)
			AssertEqual(t, *created.Data.Paused, false)
			AssertEqual(t, *created.Data.PauseAfterTrial, false)
			AssertEqual(t, *created.Data.SyncFrequency, 360)
			AssertEqual(t, created.Data.ScheduleType, "auto")

			AssertEqual(t, created.Data.Status.SetupState, "incomplete")
			AssertEqual(t, created.Data.Status.SyncState, "scheduled")
			AssertEqual(t, created.Data.Status.UpdateState, "on_schedule")
			AssertEqual(t, *created.Data.Status.IsHistoricalSync, true)
			AssertHasLength(t, created.Data.Status.Tasks, 0)
			AssertHasLength(t, created.Data.Status.Warnings, 0)

			AssertEqual(t, created.Data.Config.Password, "******")
			AssertEqual(t, created.Data.Config.TimeframeMonths, "TWELVE")
			AssertEqual(t, created.Data.Config.AppSyncMode, "AllApps")
			AssertEqual(t, created.Data.Config.SalesAccountSyncMode, "AllSalesAccounts")
			AssertEqual(t, created.Data.Config.FinanceAccountSyncMode, "AllFinanceAccounts")
			AssertEqual(t, created.Data.Config.Username, "fivetran")
			t.Cleanup(func() { DeleteConnector(t, created.Data.ID) })
		})
	}
}
