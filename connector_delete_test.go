package fivetran_test

import (
	"context"
	"testing"
)

func TestNewConnectorDeleteIntegration(t *testing.T) {
	for version, c := range Clients {
		t.Run(version, func(t *testing.T) {
			connectorId := CreateConnector(t)
			deleted, err := c.NewConnectorDelete().ConnectorID(connectorId).Do(context.Background())

			if err != nil {
				t.Logf("%+v\n", deleted)
				t.Error(err)
			}

			AssertEqual(t, deleted.Code, "Success")
			AssertEqual(t, deleted.Message, "Connector with id '"+connectorId+"' has been deleted")
		})
	}
}
