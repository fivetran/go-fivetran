package fivetran_test

import (
	"context"
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
	AssertEqual(t, sync.Message, "Sync has been successfully triggered for connector with id '"+connectorId+"'")
}
