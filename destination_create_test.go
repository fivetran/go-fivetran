package fivetran_test

import (
	"context"
	"testing"

	"github.com/fivetran/go-fivetran"
)

func TestNewDestinationCreateIntegration(t *testing.T) {
	for version, c := range Clients {
		t.Run(version, func(t *testing.T) {
			if version == "v2" {
				t.Skip("Will be tested after port fixes in v2")
			}
			created, err := c.NewDestinationCreate().
				GroupID(PredefinedGroupId).
				Service("snowflake").
				TimeZoneOffset("+10").
				RunSetupTests(false).
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

			AssertEqual(t, created.Code, "Success")
			AssertEqual(t, created.Message, "Destination has been created")
			AssertEqual(t, created.Data.ID, PredefinedGroupId)
			AssertEqual(t, created.Data.Service, "snowflake")
			AssertEqual(t, created.Data.Region, "US")
			AssertEqual(t, created.Data.TimeZoneOffset, "+10")
			AssertEqual(t, created.Data.SetupStatus, "incomplete")
			AssertEqual(t, created.Data.Config.Host, "your-account.snowflakecomputing.com")
			AssertEqual(t, created.Data.Config.Port, "443")
			AssertEqual(t, created.Data.Config.Database, "fivetran")
			AssertEqual(t, created.Data.Config.Auth, "PASSWORD")
			AssertEqual(t, created.Data.Config.User, "fivetran_user")
			AssertEqual(t, created.Data.Config.Password, "******")

			t.Cleanup(func() { DeleteDestination(t, PredefinedGroupId) })
		})
	}
}
