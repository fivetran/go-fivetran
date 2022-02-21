package fivetran_test

import (
	"context"
	"strings"
	"testing"
)

func TestNewConnectorDeleteE2E(t *testing.T) {
	connectorId := CreateConnector(t)
	deleted, err := Client.NewConnectorDelete().ConnectorID(connectorId).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}

	AssertEqual(t, deleted.Code, "Success")
	AssertNotEmpty(t, deleted.Message)
	AssertEqual(t, strings.Contains(deleted.Message, connectorId), true)
}
