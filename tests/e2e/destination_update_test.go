package fivetran_test

import (
	"context"
	"testing"

	"github.com/fivetran/go-fivetran"
	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewDestinationUpdateE2E(t *testing.T) {
	destinationId := testutils.CreateTempDestination(t)
	details, err := testutils.Client.NewDestinationUpdate().DestinationID(destinationId).
		Region("GCP_AUSTRALIA_SOUTHEAST1").
		TimeZoneOffset("+10").
		RunSetupTests(false).
		DaylightSavingTimeEnabled(true).
		Config(fivetran.NewDestinationConfig().
			Host("updated_host.snowflakecomputing.com").
			Port(444).
			Database("fivetran_updated").
			Auth("PASSWORD").
			User("fivetran_user_updated").
			Password("12345678")).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", details)
		t.Error(err)
	}

	testutils.AssertEqual(t, details.Code, "Success")
	testutils.AssertNotEmpty(t, details.Message)
	testutils.AssertEqual(t, details.Data.ID, destinationId)
	testutils.AssertEqual(t, details.Data.GroupID, destinationId)
	testutils.AssertEqual(t, details.Data.Service, "snowflake")
	testutils.AssertEqual(t, details.Data.Region, "GCP_AUSTRALIA_SOUTHEAST1")
	testutils.AssertEqual(t, details.Data.TimeZoneOffset, "+10")
	testutils.AssertEqual(t, details.Data.DaylightSavingTimeEnabled, true)
	testutils.AssertEmpty(t, details.Data.PrivateLinkId)
	testutils.AssertEmpty(t, details.Data.HybridDeploymentAgentId)
	testutils.AssertEqual(t, details.Data.NetworkingMethod, "Directly")
	testutils.AssertEqual(t, details.Data.Config.Host, "updated_host.snowflakecomputing.com")
	testutils.AssertEqual(t, details.Data.Config.Port, "444")
	testutils.AssertEqual(t, details.Data.Config.Database, "fivetran_updated")
	testutils.AssertEqual(t, details.Data.Config.User, "fivetran_user_updated")
	testutils.AssertEqual(t, details.Data.Config.Password, "******")
}

func TestNewDestinationUpdateCustomE2E(t *testing.T) {
	destinationId := testutils.CreateTempDestination(t)
	details, err := testutils.Client.NewDestinationUpdate().DestinationID(destinationId).
		Region("GCP_AUSTRALIA_SOUTHEAST1").
		TimeZoneOffset("+10").
		RunSetupTests(false).
		DaylightSavingTimeEnabled(false).
		ConfigCustom(&map[string]interface{}{
			"host":     "updated_host.snowflakecomputing.com",
			"port":     444,
			"database": "fivetran_updated",
			"auth":     "PASSWORD",
			"user":     "fivetran_user_updated",
			"password": "12345678",
		}).
		DoCustom(context.Background())

	if err != nil {
		t.Logf("%+v\n", details)
		t.Error(err)
	}

	testutils.AssertEqual(t, details.Code, "Success")
	testutils.AssertNotEmpty(t, details.Message)
	testutils.AssertEqual(t, details.Data.ID, destinationId)
	testutils.AssertEqual(t, details.Data.GroupID, destinationId)
	testutils.AssertEqual(t, details.Data.Service, "snowflake")
	testutils.AssertEqual(t, details.Data.Region, "GCP_AUSTRALIA_SOUTHEAST1")
	testutils.AssertEqual(t, details.Data.TimeZoneOffset, "+10")
	testutils.AssertEqual(t, details.Data.DaylightSavingTimeEnabled, false)
	testutils.AssertEmpty(t, details.Data.PrivateLinkId)
	testutils.AssertEmpty(t, details.Data.HybridDeploymentAgentId)
	testutils.AssertEqual(t, details.Data.NetworkingMethod, "Directly")
	testutils.AssertEqual(t, details.Data.Config["host"], "updated_host.snowflakecomputing.com")
	testutils.AssertEqual(t, details.Data.Config["port"], "444")
	testutils.AssertEqual(t, details.Data.Config["database"], "fivetran_updated")
	testutils.AssertEqual(t, details.Data.Config["user"], "fivetran_user_updated")
	testutils.AssertEqual(t, details.Data.Config["password"], "******")
}
