package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewConnectionDetailsE2E(t *testing.T) {

	ConnectionId := testutils.CreateTempConnection(t)
	details, err := testutils.Client.NewConnectionDetails().ConnectionID(ConnectionId).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", details)
		t.Error(err)
	}

	testutils.AssertEqual(t, details.Code, "Success")
	testutils.AssertEqual(t, details.Data.ID, ConnectionId)
	testutils.AssertEqual(t, details.Data.GroupID, testutils.PredefinedGroupId)
	testutils.AssertEqual(t, details.Data.Service, "itunes_connect")
	testutils.AssertEqual(t, *details.Data.ServiceVersion, 1)
	testutils.AssertEqual(t, details.Data.Schema, "itunes_e2e_connect")
	testutils.AssertEqual(t, details.Data.ConnectedBy, testutils.PredefinedUserId)
	testutils.AssertEqual(t, details.Data.CreatedAt.IsZero(), false)
	testutils.AssertEqual(t, details.Data.SucceededAt.IsZero(), true)
	testutils.AssertEqual(t, details.Data.FailedAt.IsZero(), true)
	testutils.AssertEqual(t, *details.Data.Paused, false)
	testutils.AssertEqual(t, *details.Data.PauseAfterTrial, false)
	testutils.AssertEqual(t, *details.Data.SyncFrequency, 360)
	testutils.AssertEqual(t, details.Data.ScheduleType, "auto")

	testutils.AssertEqual(t, details.Data.Status.SetupState, "incomplete")
	testutils.AssertEqual(t, details.Data.Status.SyncState, "scheduled")
	testutils.AssertEqual(t, details.Data.Status.UpdateState, "on_schedule")
	testutils.AssertEqual(t, *details.Data.Status.IsHistoricalSync, true)
	testutils.AssertHasLength(t, details.Data.Status.Tasks, 0)
	testutils.AssertHasLength(t, details.Data.Status.Warnings, 0)

	testutils.AssertEqual(t, details.Data.Config.Password, "******")
	testutils.AssertEqual(t, details.Data.Config.TimeframeMonths, "TWELVE")
	testutils.AssertEqual(t, details.Data.Config.AppSyncMode, "AllApps")
	testutils.AssertEqual(t, details.Data.Config.SalesAccountSyncMode, "AllSalesAccounts")
	testutils.AssertEqual(t, details.Data.Config.FinanceAccountSyncMode, "AllFinanceAccounts")
	testutils.AssertEqual(t, details.Data.Config.Username, "fivetran")
}
