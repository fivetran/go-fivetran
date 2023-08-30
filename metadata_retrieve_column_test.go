package fivetran_test

import (
    "context"
    "testing"
)

func TestNewMetadataColumnE2E(t *testing.T) {
    t.Skip("Skipping test until we will get more information about the status of this API")

    connectorId := CreateConnector(t)
    details, err := Client.NewMetadataColumnList().ConnectorId(connectorId).Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", details)
        t.Error(err)
    }

    AssertEqual(t, details.Code, "Success")
    AssertNotEmpty(t, details.Data.Items[0].Id)
    AssertNotEmpty(t, details.Data.Items[0].ParentId)
    AssertNotEmpty(t, details.Data.Items[0].NameInSource)
    AssertNotEmpty(t, details.Data.Items[0].NameInDestination)
    AssertNotEmpty(t, details.Data.Items[0].TypeInSource)
    AssertNotEmpty(t, details.Data.Items[0].TypeInDestination)
    AssertNotEmpty(t, details.Data.Items[0].IsPrimaryKey)
    AssertNotEmpty(t, details.Data.Items[0].IsForeignKey)

    t.Cleanup(func() { DeleteConnector(t, connectorId) })    
}
