package tests

import (
	"context"
	"testing"

	"github.com/corbym/gocrest/has"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
)

func TestConnectorSetupTests(t *testing.T) {
	for _, c := range GetClients() {
		details, err := c.NewConnectorSetupTests().
			ConnectorID("armadillo_vertebrate").
			TrustCertificates(true).
			TrustFingerprints(false).
			Do(context.Background())

		if err != nil {
			t.Logf("%+v\n", details)
			t.Error(err)
		}

		then.AssertThat(t, details.Code, is.EqualTo("Success"))
		then.AssertThat(t, details.Message, is.EqualTo("Setup tests were completed"))

		then.AssertThat(t, details.Data.ID, is.EqualTo("armadillo_vertebrate"))
		then.AssertThat(t, details.Data.GroupID, is.EqualTo("_moonbeam"))
		then.AssertThat(t, details.Data.Service, is.EqualTo("zendesk"))
		then.AssertThat(t, *details.Data.ServiceVersion, is.EqualTo(2))
		then.AssertThat(t, details.Data.Schema, is.EqualTo("zendesk"))
		then.AssertThat(t, details.Data.ConnectedBy, is.EqualTo("_airworthy"))
		then.AssertThat(t, details.Data.CreatedAt.IsZero(), is.False())
		then.AssertThat(t, details.Data.SucceededAt.IsZero(), is.False())
		then.AssertThat(t, details.Data.FailedAt.IsZero(), is.False())
		//todo:incomment after fix T-168261
		//then.AssertThat(t, *details.Data.Paused, is.False())
		//then.AssertThat(t, *details.Data.PauseAfterTrial, is.False())
		then.AssertThat(t, *details.Data.SyncFrequency, is.EqualTo(360))
		then.AssertThat(t, details.Data.ScheduleType, is.EqualTo("auto"))


		//then.AssertThat(t, details.Data.Status.SetupState, is.EqualTo("broken"))

		then.AssertThat(t, details.Data.Status.SyncState, is.EqualTo("scheduled"))
		then.AssertThat(t, details.Data.Status.UpdateState, is.EqualTo("delayed"))
		then.AssertThat(t, *details.Data.Status.IsHistoricalSync, is.False())
		then.AssertThat(t, details.Data.Status.Tasks, has.Length(1))
		then.AssertThat(t, details.Data.Status.Warnings, has.Length(0))
	}
}
