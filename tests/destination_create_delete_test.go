package tests

import (
	"context"
	"testing"
	"github.com/fivetran/go-fivetran"
	"github.com/corbym/gocrest/then"
	"github.com/corbym/gocrest/is"
)

func TestDestinationCreateDelete(t *testing.T) {
	for _, c := range GetClients() {
		//create
		created, err := c.NewDestinationCreate().
		GroupID("_moonbeam_bright").
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

		then.AssertThat(t, created.Code, is.EqualTo("Success"))
		then.AssertThat(t, created.Message, is.EqualTo("Destination has been created"))
		then.AssertThat(t, created.Data.ID, is.EqualTo("_moonbeam_bright"))
		then.AssertThat(t, created.Data.Service, is.EqualTo("snowflake"))
		then.AssertThat(t, created.Data.Region, is.EqualTo("US"))
		then.AssertThat(t, created.Data.TimeZoneOffset, is.EqualTo("+10"))
		then.AssertThat(t, created.Data.SetupStatus, is.EqualTo("incomplete"))
		then.AssertThat(t, created.Data.Config.Host, is.EqualTo("your-account.snowflakecomputing.com"))
		then.AssertThat(t, created.Data.Config.Port, is.EqualTo("443"))
		then.AssertThat(t, created.Data.Config.Database, is.EqualTo("fivetran"))
		then.AssertThat(t, created.Data.Config.Auth, is.EqualTo("PASSWORD"))
		then.AssertThat(t, created.Data.Config.User, is.EqualTo("fivetran_user"))
		then.AssertThat(t, created.Data.Config.Password, is.EqualTo("******"))

		//delete
		deleted, err := c.NewDestinationDelete().DestinationID("_moonbeam_bright").Do(context.Background())

		if err != nil {
			t.Logf("%+v\n", deleted)
			t.Error(err)
		}

		then.AssertThat(t, deleted.Code, is.EqualTo("Success"))
		then.AssertThat(t, deleted.Message, is.EqualTo("Destination with id '_moonbeam_bright' has been deleted"))
	}
}