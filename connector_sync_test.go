package fivetran_test

import (
	"context"
	"strings"
	"testing"
)

func TestNewConnectorSyncE2E(t *testing.T) {
	connectorId := CreateTempConnector(t)
	sync, err := Client.NewConnectorSync().
		ConnectorID(connectorId).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", sync)
		t.Error(err)
	}

	AssertEqual(t, sync.Code, "Success")
	AssertNotEmpty(t, sync.Message)
	AssertEqual(t, strings.Contains(sync.Message, connectorId), true)
}
