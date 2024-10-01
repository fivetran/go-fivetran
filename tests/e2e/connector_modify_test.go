package fivetran_test

import (
	"context"
	"testing"

	"github.com/fivetran/go-fivetran"
	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewConnectorModifyE2E(t *testing.T) {
	t.Skip("The test often fails due to timeouts. It is necessary to check its work only when this resource changes")
	connectorId := testutils.CreateTempConnector(t)
	syncFrequency := 1440
	updated, err := testutils.Client.NewConnectorModify().ConnectorID(connectorId).
		Paused(true).
		PauseAfterTrial(true).
		//IsHistoricalSync(false).
		SyncFrequency(&syncFrequency).
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

	testutils.AssertEqual(t, updated.Code, "Success")
	testutils.AssertEqual(t, updated.Data.ID, connectorId)
	testutils.AssertEqual(t, updated.Data.GroupID, testutils.PredefinedGroupId)
	testutils.AssertEqual(t, updated.Data.Service, "itunes_connect")
	testutils.AssertEqual(t, *updated.Data.ServiceVersion, 1)
	testutils.AssertEqual(t, updated.Data.Schema, "itunes_e2e_connect")
	testutils.AssertEqual(t, updated.Data.ConnectedBy, testutils.PredefinedUserId)
	testutils.AssertEqual(t, updated.Data.CreatedAt.IsZero(), false)
	testutils.AssertEqual(t, updated.Data.SucceededAt.IsZero(), true)
	testutils.AssertEqual(t, updated.Data.FailedAt.IsZero(), true)
	testutils.AssertEqual(t, *updated.Data.Paused, true)
	testutils.AssertEqual(t, *updated.Data.PauseAfterTrial, true)
	testutils.AssertEqual(t, *updated.Data.SyncFrequency, 1440)
	testutils.AssertEqual(t, updated.Data.ScheduleType, "auto")
	testutils.AssertEmpty(t, updated.Data.PrivateLinkId)
	testutils.AssertEmpty(t, updated.Data.HybridDeploymentAgentId)
	testutils.AssertEmpty(t, updated.Data.ProxyAgentId)
	testutils.AssertEqual(t, updated.Data.NetworkingMethod, "Directly")

	testutils.AssertNotEmpty(t, updated.Data.Status.SetupState)
	testutils.AssertEqual(t, updated.Data.Status.SyncState, "paused")
	testutils.AssertEqual(t, updated.Data.Status.UpdateState, "on_schedule")
	//todo: check after fix
	//testutils.AssertEqual(t, *updated.Data.Status.IsHistoricalSync, false)
	testutils.AssertHasLength(t, updated.Data.Status.Tasks, 0)
	testutils.AssertHasLength(t, updated.Data.Status.Warnings, 0)

	testutils.AssertEqual(t, updated.Data.Config.Password, "******")
	testutils.AssertEqual(t, updated.Data.Config.TimeframeMonths, "SIX")
	testutils.AssertEqual(t, updated.Data.Config.AppSyncMode, "AllApps")
	testutils.AssertEqual(t, updated.Data.Config.SalesAccountSyncMode, "AllSalesAccounts")
	testutils.AssertEqual(t, updated.Data.Config.FinanceAccountSyncMode, "AllFinanceAccounts")
	testutils.AssertEqual(t, updated.Data.Config.Username, "fivetran_updated")
}
