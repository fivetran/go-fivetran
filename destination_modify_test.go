package fivetran_test

import (
	"context"
	"testing"

	"github.com/fivetran/go-fivetran"
)

func TestNewDestinationModifyE2E(t *testing.T) {
	destinationId := CreateTempDestination(t)
	details, err := Client.NewDestinationModify().DestinationID(destinationId).
		Region("APAC").
		TimeZoneOffset("+10").
		RunSetupTests(false).
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

	AssertEqual(t, details.Code, "Success")
	AssertEqual(t, details.Message, "Destination has been updated")
	AssertEqual(t, details.Data.ID, destinationId)
	AssertEqual(t, details.Data.GroupID, destinationId)
	AssertEqual(t, details.Data.Service, "snowflake")
	AssertEqual(t, details.Data.Region, "APAC")
	AssertEqual(t, details.Data.TimeZoneOffset, "+10")
	AssertEqual(t, details.Data.Config.Host, "updated_host.snowflakecomputing.com")
	AssertEqual(t, details.Data.Config.Port, "444")
	AssertEqual(t, details.Data.Config.Database, "fivetran_updated")
	AssertEqual(t, details.Data.Config.User, "fivetran_user_updated")
	AssertEqual(t, details.Data.Config.Password, "******")
}
