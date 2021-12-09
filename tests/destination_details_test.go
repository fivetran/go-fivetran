package tests

import (
	"context"
	"testing"
	"github.com/corbym/gocrest/then"
	"github.com/corbym/gocrest/is"
)

func TestDestinationDetails(t *testing.T) {
	for _, c := range GetClients() {
		details, err := c.NewDestinationDetails().DestinationID("_moonbeam").Do(context.Background())
		
		if err != nil {
			t.Logf("%+v\n", details)
			t.Error(err)
		}

		then.AssertThat(t, details.Code, is.EqualTo("Success"))
		then.AssertThat(t, details.Data.ID, is.EqualTo("_moonbeam"))
		then.AssertThat(t, details.Data.GroupID, is.EqualTo("_moonbeam"))
		then.AssertThat(t, details.Data.Service, is.EqualTo("redshift"))
		then.AssertThat(t, details.Data.Region, is.EqualTo("US"))
		then.AssertThat(t, details.Data.TimeZoneOffset, is.EqualTo("-8"))
		then.AssertThat(t, details.Data.SetupStatus, is.EqualTo("connected"))
		then.AssertThat(t, details.Data.Config.Database, is.EqualTo("_moonbeam"))
		then.AssertThat(t, details.Data.Config.Password, is.EqualTo("******"))
		then.AssertThat(t, details.Data.Config.Port, is.EqualTo("5432"))
		then.AssertThat(t, details.Data.Config.Host, is.EqualTo("localhost"))
		then.AssertThat(t, details.Data.Config.User, is.EqualTo("moonbeam"))
		then.AssertThat(t, details.Data.Config.ConnectionMethod, is.EqualTo("Directly"))
	}
}