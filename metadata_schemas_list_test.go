package fivetran_test

import (
    "context"
    "testing"
)

func TestNewMetadataSchemasE2E(t *testing.T) {
    t.Skip("Skipping test until we will get more information about the status of this API")

    connectorId := CreateConnector(t)
    details, err := Client.NewMetadataSchemasList().ConnectorId(connectorId).Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", details)
        t.Error(err)
    }

    AssertEqual(t, details.Code, "Success")
    AssertNotEmpty(t, details.Data.Items[0].Id)
    AssertNotEmpty(t, details.Data.Items[0].NameInSource)
    AssertNotEmpty(t, details.Data.Items[0].NameInDestination)

    t.Cleanup(func() { DeleteConnector(t, connectorId) })
}
