package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewMetadataTableE2E(t *testing.T) {
	t.Skip("Skipping test until we will get more information about the status of this API")

	connectorId := testutils.CreateConnector(t)
	details, err := testutils.Client.NewMetadataTablesList().ConnectorId(connectorId).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", details)
		t.Error(err)
	}

	testutils.AssertEqual(t, details.Code, "Success")
	testutils.AssertNotEmpty(t, details.Data.Items[0].Id)
	testutils.AssertNotEmpty(t, details.Data.Items[0].ParentId)
	testutils.AssertNotEmpty(t, details.Data.Items[0].NameInSource)
	testutils.AssertNotEmpty(t, details.Data.Items[0].NameInDestination)

	t.Cleanup(func() { testutils.DeleteConnector(t, connectorId) })
}
