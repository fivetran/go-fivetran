package fivetran_test

import (
	"context"
	"testing"
)

func TestNewDestinationDetailsIntegration(t *testing.T) {
	for version, c := range Clients {
		//todo: remove it after v2 fixes
		if version == "v2" {
			continue
		}

		t.Run(version, func(t *testing.T) {
			destinationId := CreateTempDestination(t)

			details, err := c.NewDestinationDetails().DestinationID(destinationId).Do(context.Background())

			if err != nil {
				t.Logf("%+v\n", details)
				t.Error(err)
			}

			AssertEqual(t, details.Code, "Success")
			AssertEqual(t, details.Data.ID, "climbed_consulted")
			AssertEqual(t, details.Data.GroupID, "climbed_consulted")
			AssertEqual(t, details.Data.Service, "snowflake")
			AssertEqual(t, details.Data.Region, "US")
			AssertEqual(t, details.Data.TimeZoneOffset, "+10")
			AssertEqual(t, details.Data.SetupStatus, "incomplete")
			AssertEqual(t, details.Data.Config.Database, "fivetran")
			AssertEqual(t, details.Data.Config.Password, "******")
			AssertEqual(t, details.Data.Config.Port, "443")
			AssertEqual(t, details.Data.Config.Host, "your-account.snowflakecomputing.com")
			AssertEqual(t, details.Data.Config.User, "fivetran_user")
		})
	}
}
