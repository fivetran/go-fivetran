package fivetran_test

import (
	"context"
	"testing"

	"github.com/fivetran/go-fivetran"
)

func TestNewDestinationCreateIntegration(t *testing.T) {
	for version, c := range Clients {
		//todo: remove it after v2 fixes
		if version == "v2" {
			continue
		}
		t.Run(version, func(t *testing.T) {
			created, err := c.NewDestinationCreate().
				GroupID("climbed_consulted").
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
			AssertEqual(t, created.Data.ID, "climbed_consulted")
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

			t.Cleanup(func() { DeleteDestination(t, "climbed_consulted") })
		})
	}
}
