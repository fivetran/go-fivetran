package tests

import (
	"context"
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
)

func TestConnectorSync(t *testing.T) {
	for _, c := range GetClients() {
		sync, err := c.NewConnectorSync().
		ConnectorID("goes_headlock").
		Do(context.Background())

		if err != nil {
			t.Logf("%+v\n", sync)
			t.Error(err)
		}

		then.AssertThat(t, sync.Code, is.EqualTo("Success"))
		then.AssertThat(t, sync.Message, is.EqualTo("Sync has been successfully triggered for connector with id 'goes_headlock'"))
	}
}