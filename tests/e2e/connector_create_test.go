package fivetran_test

import (
	"context"
	"testing"

	"github.com/fivetran/go-fivetran"
	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewConnectorCreateE2E(t *testing.T) {
	created, err := testutils.Client.NewConnectorCreate().
		GroupID(testutils.PredefinedGroupId).
		Service("itunes_connect").
		RunSetupTests(false).
		NetworkingMethod("Directly").
		Config(fivetran.NewConnectorConfig().
			Schema("itunes_e2e_connect").
			Username("fivetran").
			Password("fivetran-api-e2e")).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertNotEmpty(t, created.Message)
	testutils.AssertNotEmpty(t, created.Data.ID)
	testutils.AssertEqual(t, created.Data.GroupID, testutils.PredefinedGroupId)
	testutils.AssertEqual(t, created.Data.Service, "itunes_connect")
	testutils.AssertEqual(t, *created.Data.ServiceVersion, 1)
	testutils.AssertEqual(t, created.Data.Schema, "itunes_e2e_connect")
	testutils.AssertEqual(t, created.Data.ConnectedBy, testutils.PredefinedUserId)
	testutils.AssertEqual(t, created.Data.CreatedAt.IsZero(), false)
	testutils.AssertEqual(t, created.Data.SucceededAt.IsZero(), true)
	testutils.AssertEqual(t, created.Data.FailedAt.IsZero(), true)
	testutils.AssertEqual(t, *created.Data.Paused, false)
	testutils.AssertEqual(t, *created.Data.PauseAfterTrial, false)
	testutils.AssertEqual(t, *created.Data.SyncFrequency, "360")
	testutils.AssertEqual(t, created.Data.ScheduleType, "auto")
	testutils.AssertEmpty(t, created.Data.PrivateLinkId)
	testutils.AssertEmpty(t, created.Data.LocalProcessingAgentId)
	testutils.AssertEmpty(t, created.Data.ProxyAgentId)
	testutils.AssertEqual(t, created.Data.NetworkingMethod, "Directly")

	testutils.AssertEqual(t, created.Data.Status.SetupState, "incomplete")
	testutils.AssertEqual(t, created.Data.Status.SyncState, "scheduled")
	testutils.AssertEqual(t, created.Data.Status.UpdateState, "on_schedule")
	testutils.AssertEqual(t, *created.Data.Status.IsHistoricalSync, true)
	testutils.AssertHasLength(t, created.Data.Status.Tasks, 0)
	testutils.AssertHasLength(t, created.Data.Status.Warnings, 0)

	testutils.AssertEqual(t, created.Data.Config.Password, "******")
	testutils.AssertEqual(t, created.Data.Config.TimeframeMonths, "TWELVE")
	testutils.AssertEqual(t, created.Data.Config.AppSyncMode, "AllApps")
	testutils.AssertEqual(t, created.Data.Config.SalesAccountSyncMode, "AllSalesAccounts")
	testutils.AssertEqual(t, created.Data.Config.FinanceAccountSyncMode, "AllFinanceAccounts")
	testutils.AssertEqual(t, created.Data.Config.Username, "fivetran")
	t.Cleanup(func() { testutils.DeleteConnector(t, created.Data.ID) })
}
