package fivetran_test

import (
	"context"
	"testing"

	"github.com/fivetran/go-fivetran"
)

func TestNewExternalLoggingModifyE2E(t *testing.T) {
	externalLoggingId := CreateTempExternalLogging(t)
	details, err := Client.NewExternalLoggingModify().ExternalLoggingID(externalLoggingId).
		Region("GCP_AUSTRALIA_SOUTHEAST1").
		TimeZoneOffset("+10").
		RunSetupTests(false).
		Config(fivetran.NewExternalLoggingConfig().
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

	AssertEqual(t, details.Code, "Success")
	AssertNotEmpty(t, details.Message)
	AssertEqual(t, details.Data.ID, externalLoggingId)
	AssertEqual(t, details.Data.GroupID, externalLoggingId)
	AssertEqual(t, details.Data.Service, "snowflake")
	AssertEqual(t, details.Data.Region, "GCP_AUSTRALIA_SOUTHEAST1")
	AssertEqual(t, details.Data.TimeZoneOffset, "+10")
	AssertEqual(t, details.Data.Config.Host, "updated_host.snowflakecomputing.com")
	AssertEqual(t, details.Data.Config.Port, "444")
	AssertEqual(t, details.Data.Config.Database, "fivetran_updated")
	AssertEqual(t, details.Data.Config.User, "fivetran_user_updated")
	AssertEqual(t, details.Data.Config.Password, "******")
}
