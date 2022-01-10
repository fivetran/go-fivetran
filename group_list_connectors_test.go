package fivetran_test

import (
	"context"
	"testing"
)

func TestNewGroupListConnectorsE2E(t *testing.T) {
	connectorId := CreateTempConnector(t)
	connectors, err := Client.NewGroupListConnectors().GroupID(PredefinedGroupId).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", connectors)
		t.Error(err)
	}

	AssertEqual(t, connectors.Code, "Success")
	AssertHasLength(t, connectors.Data.Items, 1)

	AssertEqual(t, connectors.Data.Items[0].ID, connectorId)
	AssertEqual(t, connectors.Data.Items[0].GroupID, PredefinedGroupId)
	AssertEqual(t, connectors.Data.Items[0].Service, "itunes_connect")
	AssertEqual(t, *connectors.Data.Items[0].ServiceVersion, 1)
	AssertEqual(t, connectors.Data.Items[0].Schema, "itunes_e2e_connect")
	AssertEqual(t, connectors.Data.Items[0].ConnectedBy, PredefinedUserId)
	AssertEqual(t, connectors.Data.Items[0].CreatedAt.IsZero(), false)
	AssertEqual(t, connectors.Data.Items[0].SucceededAt.IsZero(), true)
	AssertEqual(t, connectors.Data.Items[0].FailedAt.IsZero(), true)

	//todo: map Paused and PauseAfterTrial
	//AssertEqual(t, *connectors.Data.Items[0].Paused, false)
	//AssertEqual(t, *connectors.Data.Items[0].PauseAfterTrial, false)

	AssertEqual(t, *connectors.Data.Items[0].SyncFrequency, 360)
	AssertEqual(t, connectors.Data.Items[0].ScheduleType, "auto")

	AssertEqual(t, connectors.Data.Items[0].Status.SetupState, "incomplete")
	AssertEqual(t, connectors.Data.Items[0].Status.SyncState, "scheduled")
	AssertEqual(t, connectors.Data.Items[0].Status.UpdateState, "on_schedule")
	AssertEqual(t, *connectors.Data.Items[0].Status.IsHistoricalSync, true)
	AssertHasLength(t, connectors.Data.Items[0].Status.Tasks, 0)
	AssertHasLength(t, connectors.Data.Items[0].Status.Warnings, 0)
}
