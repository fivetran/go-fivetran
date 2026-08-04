package fivetran_test

import (
	"context"
	"testing"

	"github.com/fivetran/go-fivetran"
	testutils "github.com/fivetran/go-fivetran/test_utils"
)

const (
	DATA_DELAY_THRESHOLD = 1
)

func TestNewConnectionCreateE2E(t *testing.T) {
	dataDelayThreshold := 1
	created, err := testutils.Client.NewConnectionCreate().
		GroupID(testutils.PredefinedGroupId).
		Service("itunes_connect").
		RunSetupTests(false).
		NetworkingMethod("Directly").
		DataDelayThreshold(&dataDelayThreshold).
		DataDelaySensitivity("CUSTOM").
		DestinationSchemaNames("FIVETRAN_NAMING").
		Config(fivetran.NewConnectionConfig().
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
	testutils.AssertEqual(t, *created.Data.SyncFrequency, 360)
	testutils.AssertEqual(t, created.Data.ScheduleType, "auto")
	testutils.AssertEmpty(t, created.Data.PrivateLinkId)
	testutils.AssertEmpty(t, created.Data.HybridDeploymentAgentId)
	testutils.AssertEmpty(t, created.Data.ProxyAgentId)
	testutils.AssertEqual(t, created.Data.NetworkingMethod, "Directly")
	testutils.AssertEqual(t, *created.Data.DataDelayThreshold, 1)
	testutils.AssertEqual(t, created.Data.DataDelaySensitivity, "CUSTOM")

	testutils.AssertEqual(t, created.Data.Status.SetupState, "incomplete")
	testutils.AssertEqual(t, created.Data.Status.SyncState, "scheduled")
	testutils.AssertEqual(t, created.Data.Status.UpdateState, "on_schedule")
	testutils.AssertEqual(t, *created.Data.Status.IsHistoricalSync, true)
	testutils.AssertHasLength(t, created.Data.Status.Tasks, 0)
	testutils.AssertHasLength(t, created.Data.Status.Warnings, 0)

	testutils.AssertEqual(t, created.Data.Config.TimeframeMonths, "TWELVE")
	testutils.AssertEqual(t, created.Data.Config.AppSyncMode, "AllApps")
	t.Cleanup(func() { testutils.DeleteConnection(t, created.Data.ID) })
}

func TestNewConnectionCreateSourceNamingE2E(t *testing.T) {
	testutils.CreateTempDestination(t)

	dataDelayThreshold := 1
	created, err := testutils.Client.NewConnectionCreate().
		GroupID(testutils.PredefinedGroupId).
		Service("postgres_rds").
		RunSetupTests(false).
		NetworkingMethod("Directly").
		DataDelayThreshold(&dataDelayThreshold).
		DataDelaySensitivity("CUSTOM").
		DestinationSchemaNames("SOURCE_NAMING").
		Config(fivetran.NewConnectionConfig().
			SchemaPrefix("postgres_e2e_source_naming").
			Host("terraform-pgsql-connection-test.cp0rdhwjbsae.us-east-1.rds.amazonaws.com").
			Port(5432).
			Database("fivetran").
			User("postgres").
			Password("mYP4ssw0rd").
			UpdateMethod("XMIN")).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Fatal(err)
	}

	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertNotEmpty(t, created.Message)
	testutils.AssertNotEmpty(t, created.Data.ID)
	testutils.AssertEqual(t, created.Data.GroupID, testutils.PredefinedGroupId)
	testutils.AssertEqual(t, created.Data.Service, "postgres_rds")
	testutils.AssertEqual(t, created.Data.Schema, "postgres_e2e_source_naming")
	testutils.AssertEqual(t, created.Data.ConnectedBy, testutils.PredefinedUserId)
	testutils.AssertEqual(t, created.Data.CreatedAt.IsZero(), false)
	testutils.AssertEqual(t, created.Data.SucceededAt.IsZero(), true)
	testutils.AssertEqual(t, created.Data.FailedAt.IsZero(), true)
	testutils.AssertEqual(t, *created.Data.Paused, true)
	testutils.AssertEqual(t, *created.Data.PauseAfterTrial, false)
	testutils.AssertEqual(t, *created.Data.SyncFrequency, 360)
	testutils.AssertEqual(t, created.Data.ScheduleType, "auto")
	testutils.AssertEmpty(t, created.Data.PrivateLinkId)
	testutils.AssertEmpty(t, created.Data.HybridDeploymentAgentId)
	testutils.AssertEmpty(t, created.Data.ProxyAgentId)
	testutils.AssertEqual(t, created.Data.NetworkingMethod, "Directly")
	testutils.AssertEqual(t, *created.Data.DataDelayThreshold, 1)
	testutils.AssertEqual(t, created.Data.DataDelaySensitivity, "CUSTOM")
	testutils.AssertEqual(t, *created.Data.DestinationSchemaNames, "SOURCE_NAMING")

	testutils.AssertEqual(t, created.Data.Status.SetupState, "incomplete")
	testutils.AssertEqual(t, created.Data.Status.SyncState, "paused")
	testutils.AssertEqual(t, created.Data.Status.UpdateState, "on_schedule")
	testutils.AssertEqual(t, *created.Data.Status.IsHistoricalSync, true)
	testutils.AssertHasLength(t, created.Data.Status.Tasks, 0)
	testutils.AssertHasLength(t, created.Data.Status.Warnings, 0)

	t.Cleanup(func() { testutils.DeleteConnection(t, created.Data.ID) })
}
