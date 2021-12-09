package tests

import (
	"context"
	"testing"

	"github.com/corbym/gocrest/then"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/has"
)
func TestGroupListConnectors(t *testing.T) {
	for _, c := range GetClients() {
		connectors, err := c.NewGroupListConnectors().GroupID("_moonbeam").Do(context.Background())
		if err != nil {
			t.Logf("%+v\n", connectors)
			t.Error(err)
		}
		then.AssertThat(t, connectors.Code, is.EqualTo("Success"))
		then.AssertThat(t, connectors.Data.Items, has.Length(9))
		then.AssertThat(t, connectors.Data.Items[0].ID, is.EqualTo("abject_normative"))
		then.AssertThat(t, connectors.Data.Items[0].GroupID, is.EqualTo("_moonbeam"))
		then.AssertThat(t, connectors.Data.Items[0].Service, is.EqualTo("pardot"))
		then.AssertThat(t, *connectors.Data.Items[0].ServiceVersion, is.EqualTo(2))
		then.AssertThat(t, connectors.Data.Items[0].Schema, is.EqualTo("pardot"))
		then.AssertThat(t, connectors.Data.Items[0].ConnectedBy, is.EqualTo("_airworthy"))
		then.AssertThat(t, connectors.Data.Items[0].CreatedAt.IsZero(), is.False())
		then.AssertThat(t, connectors.Data.Items[0].SucceededAt.IsZero(), is.False())
		then.AssertThat(t, connectors.Data.Items[0].FailedAt.IsZero(), is.False())
		then.AssertThat(t, *connectors.Data.Items[0].SyncFrequency, is.EqualTo(360))
		then.AssertThat(t, connectors.Data.Items[0].Status.SetupState, is.EqualTo("connected"))
		then.AssertThat(t, connectors.Data.Items[0].Status.SyncState, is.EqualTo("scheduled"))
		then.AssertThat(t, connectors.Data.Items[0].Status.UpdateState, is.EqualTo("delayed"))
		then.AssertThat(t, *connectors.Data.Items[0].Status.IsHistoricalSync, is.False())
		then.AssertThat(t, connectors.Data.Items[0].Status.Tasks, has.Length(0))
		then.AssertThat(t, connectors.Data.Items[0].Status.Warnings, has.Length(0))
		then.AssertThat(t, connectors.Data.NextCursor, is.Empty())
	}
}