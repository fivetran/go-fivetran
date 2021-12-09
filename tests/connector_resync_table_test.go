package tests

import (
	"context"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
)

func TestConnectorResyncTable(t *testing.T) {
	for _, c := range GetClients() {
		resync, err := c.NewConnectorReSyncTable().ConnectorID("prodigy_vaso").Schema("schema_name1").Table("table_name1").Do(context.Background())

		if err != nil {
			t.Logf("%+v\n", resync)
			t.Error(err)
		}

		then.AssertThat(t, resync.Code, is.EqualTo("Success"))
		then.AssertThat(t, resync.Message, is.EqualTo("Re-sync has been successfully triggered for 'schema_name1.table_name1'"))
	}
}