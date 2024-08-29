package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewGroupListConnectorsE2E(t *testing.T) {
	connectorId := testutils.CreateTempConnector(t)
	connectors, err := testutils.Client.NewGroupListConnectors().GroupID(testutils.PredefinedGroupId).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", connectors)
		t.Error(err)
	}

	testutils.AssertEqual(t, connectors.Code, "Success")
	testutils.AssertHasLength(t, connectors.Data.Items, 1)

	testutils.AssertEqual(t, connectors.Data.Items[0].ID, connectorId)
	testutils.AssertEqual(t, connectors.Data.Items[0].GroupID, testutils.PredefinedGroupId)
	testutils.AssertEqual(t, connectors.Data.Items[0].Service, "itunes_connect")
	testutils.AssertEqual(t, *connectors.Data.Items[0].ServiceVersion, 1)
	testutils.AssertEqual(t, connectors.Data.Items[0].Schema, "itunes_e2e_connect")
	testutils.AssertEqual(t, connectors.Data.Items[0].ConnectedBy, testutils.PredefinedUserId)
	testutils.AssertEqual(t, connectors.Data.Items[0].CreatedAt.IsZero(), false)
	testutils.AssertEqual(t, connectors.Data.Items[0].SucceededAt.IsZero(), true)
	testutils.AssertEqual(t, connectors.Data.Items[0].FailedAt.IsZero(), true)

	//todo: map Paused and PauseAfterTrial
	//testutils.AssertEqual(t, *connectors.Data.Items[0].Paused, false)
	//testutils.AssertEqual(t, *connectors.Data.Items[0].PauseAfterTrial, false)

	testutils.AssertEqual(t, *connectors.Data.Items[0].SyncFrequency, "360")
	testutils.AssertEqual(t, connectors.Data.Items[0].ScheduleType, "auto")

	testutils.AssertEqual(t, connectors.Data.Items[0].Status.SetupState, "incomplete")
	testutils.AssertEqual(t, connectors.Data.Items[0].Status.SyncState, "scheduled")
	testutils.AssertEqual(t, connectors.Data.Items[0].Status.UpdateState, "on_schedule")
	testutils.AssertEqual(t, *connectors.Data.Items[0].Status.IsHistoricalSync, true)
	testutils.AssertHasLength(t, connectors.Data.Items[0].Status.Tasks, 0)
	testutils.AssertHasLength(t, connectors.Data.Items[0].Status.Warnings, 0)
}
