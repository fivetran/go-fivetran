package fivetran_test

import (
    "context"
    "testing"
)

func TestNewMetadataTableE2E(t *testing.T) {
    connectorId := CreateTempConnector(t)
    details, err := Client.NewMetadataTableList().ConnectorId(connectorId).Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", details)
        t.Error(err)
    }

    AssertEqual(t, details.Code, "Success")
    AssertNotEmpty(t, details.Data.Items[0].Id)
    AssertNotEmpty(t, details.Data.Items[0].ParentId)
    AssertNotEmpty(t, details.Data.Items[0].NameInSource)
    AssertNotEmpty(t, details.Data.Items[0].NameInDestination)
}
