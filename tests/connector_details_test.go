package tests

import (
	"context"
	"testing"

	"github.com/corbym/gocrest/has"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
)

func TestConnectorDetails(t *testing.T) {
	for _, c := range GetClients() {
		details, err := c.NewConnectorDetails().ConnectorID("abject_normative").Do(context.Background())

		if err != nil {
			t.Logf("%+v\n", details)
			t.Error(err)
		}

		then.AssertThat(t, details.Code, is.EqualTo("Success"))
		then.AssertThat(t, details.Data.ID, is.EqualTo("abject_normative"))
		then.AssertThat(t, details.Data.GroupID, is.EqualTo("_moonbeam"))
		then.AssertThat(t, details.Data.Service, is.EqualTo("pardot"))
		then.AssertThat(t, *details.Data.ServiceVersion, is.EqualTo(2))
		then.AssertThat(t, details.Data.Schema, is.EqualTo("pardot"))
		then.AssertThat(t, details.Data.ConnectedBy, is.EqualTo("_airworthy"))
		then.AssertThat(t, details.Data.CreatedAt.IsZero(), is.False())
		then.AssertThat(t, details.Data.SucceededAt.IsZero(), is.False())
		then.AssertThat(t, details.Data.FailedAt.IsZero(), is.False())
		then.AssertThat(t, *details.Data.Paused, is.False())
		then.AssertThat(t, *details.Data.PauseAfterTrial, is.False())
		then.AssertThat(t, *details.Data.SyncFrequency, is.EqualTo(360))
		then.AssertThat(t, details.Data.ScheduleType, is.EqualTo("auto"))

		then.AssertThat(t, details.Data.Status.SetupState, is.EqualTo("connected"))
		then.AssertThat(t, details.Data.Status.SyncState, is.EqualTo("scheduled"))
		then.AssertThat(t, details.Data.Status.UpdateState, is.EqualTo("delayed"))
		then.AssertThat(t, *details.Data.Status.IsHistoricalSync, is.False())
		then.AssertThat(t, details.Data.Status.Tasks, has.Length(0))
		then.AssertThat(t, details.Data.Status.Warnings, has.Length(0))

	
		then.AssertThat(t, details.Data.Config.Password, is.EqualTo("******"))
		then.AssertThat(t, details.Data.Config.UserKey, is.EqualTo("******"))
		then.AssertThat(t, details.Data.Config.Email, is.EqualTo("test_email@testdomain.com"))
	}
}