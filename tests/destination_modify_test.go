package tests

import (
	"context"
	"testing"
	"github.com/fivetran/go-fivetran"
	"github.com/corbym/gocrest/then"
	"github.com/corbym/gocrest/is"
)

func TestDestinationModify(t *testing.T) {
	for _, c := range GetClients() {
		details, err := c.NewDestinationModify().DestinationID("_test").
		Region("APAC").
		TimeZoneOffset("+10").
		RunSetupTests(false).
		Config(fivetran.NewDestinationConfig().
			ProjectID("project").
			DataSetLocation("US")).
		//todo: add secret key and bucket
		Do(context.Background())
		
		if err != nil {
			t.Logf("%+v\n", details)
			t.Error(err)
		}

		then.AssertThat(t, details.Code, is.EqualTo("Success"))
		then.AssertThat(t, details.Message, is.EqualTo("Destination has been updated"))
		then.AssertThat(t, details.Data.ID, is.EqualTo("_test"))
		then.AssertThat(t, details.Data.GroupID, is.EqualTo("_test"))
		then.AssertThat(t, details.Data.Service, is.EqualTo("managed_big_query"))
		then.AssertThat(t, details.Data.Region, is.EqualTo("APAC"))
		then.AssertThat(t, details.Data.TimeZoneOffset, is.EqualTo("+10"))
		then.AssertThat(t, details.Data.SetupStatus, is.EqualTo("connected"))
		then.AssertThat(t, details.Data.Config.ProjectID, is.EqualTo("project"))
		//todo: add assert on secret key and bucket
	}
}