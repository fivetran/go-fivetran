package fivetran_test

import (
	"context"
	"testing"

	"github.com/fivetran/go-fivetran"
	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewDestinationCreateE2E(t *testing.T) {
	created, err := testutils.Client.NewDestinationCreate().
		GroupID(testutils.PredefinedGroupId).
		Service("snowflake").
		TimeZoneOffset("+10").
		RunSetupTests(false).
		DaylightSavingTimeEnabled(true).
		Config(fivetran.NewDestinationConfig().
			Host("your-account.snowflakecomputing.com").
			Port(443).
			Database("fivetran").
			Auth("PASSWORD").
			User("fivetran_user").
			Password("123456")).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertNotEmpty(t, created.Message)
	testutils.AssertEqual(t, created.Data.ID, testutils.PredefinedGroupId)
	testutils.AssertEqual(t, created.Data.Service, "snowflake")
	testutils.AssertEqual(t, created.Data.Region, "GCP_US_EAST4")
	testutils.AssertEqual(t, created.Data.TimeZoneOffset, "+10")
	testutils.AssertEqual(t, created.Data.DaylightSavingTimeEnabled, true)
	testutils.AssertEqual(t, created.Data.SetupStatus, "incomplete")
	testutils.AssertEqual(t, created.Data.Config.Host, "your-account.snowflakecomputing.com")
	testutils.AssertEqual(t, created.Data.Config.Port, "443")
	testutils.AssertEqual(t, created.Data.Config.Database, "fivetran")
	testutils.AssertEqual(t, created.Data.Config.Auth, "PASSWORD")
	testutils.AssertEqual(t, created.Data.Config.User, "fivetran_user")
	testutils.AssertEqual(t, created.Data.Config.Password, "******")

	t.Cleanup(func() { testutils.DeleteDestination(t, testutils.PredefinedGroupId) })
}

func TestNewDestinationCreateCustomE2E(t *testing.T) {
	created, err := testutils.Client.NewDestinationCreate().
		GroupID(testutils.PredefinedGroupId).
		Service("snowflake").
		TimeZoneOffset("+10").
		RunSetupTests(false).
		DaylightSavingTimeEnabled(false).
		ConfigCustom(&map[string]interface{}{
			"host":     "your-account.snowflakecomputing.com",
			"port":     443,
			"database": "fivetran",
			"auth":     "PASSWORD",
			"user":     "fivetran_user",
			"password": "123456",
		}).
		DoCustom(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertNotEmpty(t, created.Message)
	testutils.AssertEqual(t, created.Data.ID, testutils.PredefinedGroupId)
	testutils.AssertEqual(t, created.Data.Service, "snowflake")
	testutils.AssertEqual(t, created.Data.Region, "GCP_US_EAST4")
	testutils.AssertEqual(t, created.Data.TimeZoneOffset, "+10")
	testutils.AssertEqual(t, created.Data.DaylightSavingTimeEnabled, false)
	testutils.AssertEqual(t, created.Data.SetupStatus, "incomplete")
	testutils.AssertEqual(t, created.Data.Config["host"], "your-account.snowflakecomputing.com")
	testutils.AssertEqual(t, created.Data.Config["port"], "443")
	testutils.AssertEqual(t, created.Data.Config["database"], "fivetran")
	testutils.AssertEqual(t, created.Data.Config["auth"], "PASSWORD")
	testutils.AssertEqual(t, created.Data.Config["user"], "fivetran_user")
	testutils.AssertEqual(t, created.Data.Config["password"], "******")

	t.Cleanup(func() { testutils.DeleteDestination(t, testutils.PredefinedGroupId) })
}
