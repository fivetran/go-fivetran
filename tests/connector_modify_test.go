package tests

import (
	"context"
	"testing"

	"github.com/corbym/gocrest/has"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/fivetran/go-fivetran"
)

func TestConnectorModify(t *testing.T) {
	for _, c := range GetClients() {
		updated, err := c.NewConnectorModify().ConnectorID("woollen_oblige").
		Paused(true).
		SyncFrequency(1440).
		DailySyncTime("03:00").
		Config(fivetran.NewConnectorConfig().
			APIQuota(5000).
			Endpoint("endpoint").
			Identity("identity")).
		Do(context.Background())

		if err != nil {
			t.Logf("%+v\n", updated)
			t.Error(err)
		}

		then.AssertThat(t, updated.Code, is.EqualTo("Success"))
		then.AssertThat(t, updated.Data.ID, is.EqualTo("woollen_oblige"))
		then.AssertThat(t, updated.Data.GroupID, is.EqualTo("_moonbeam"))
		then.AssertThat(t, updated.Data.Service, is.EqualTo("marketo"))
		then.AssertThat(t, *updated.Data.ServiceVersion, is.EqualTo(2))
		then.AssertThat(t, updated.Data.Schema, is.EqualTo("marketo"))
		then.AssertThat(t, updated.Data.ConnectedBy, is.Empty())
		then.AssertThat(t, updated.Data.CreatedAt.IsZero(), is.False())
		then.AssertThat(t, updated.Data.SucceededAt.IsZero(), is.True())
		then.AssertThat(t, updated.Data.FailedAt.IsZero(), is.True())
		then.AssertThat(t, *updated.Data.Paused, is.True())
		then.AssertThat(t, *updated.Data.PauseAfterTrial, is.False())
		then.AssertThat(t, *updated.Data.SyncFrequency, is.EqualTo(1440))
		then.AssertThat(t, updated.Data.ScheduleType, is.EqualTo("auto"))

		then.AssertThat(t, updated.Data.Status.SetupState, is.EqualTo("broken"))
		then.AssertThat(t, updated.Data.Status.SyncState, is.EqualTo("paused"))
		then.AssertThat(t, updated.Data.Status.UpdateState, is.EqualTo("on_schedule"))
		then.AssertThat(t, *updated.Data.Status.IsHistoricalSync, is.True())
		then.AssertThat(t, updated.Data.Status.Tasks, has.Length(1))
		then.AssertThat(t, updated.Data.Status.Warnings, has.Length(0))

		then.AssertThat(t, *updated.Data.Config.APIQuota, is.EqualTo(5000))
		then.AssertThat(t, updated.Data.Config.Endpoint, is.EqualTo("endpoint"))
		then.AssertThat(t, updated.Data.Config.Identity, is.EqualTo("identity"))
	}
}