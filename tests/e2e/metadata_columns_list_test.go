package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewMetadataColumnsE2E(t *testing.T) {
	t.Skip("Skipping test until we will get more information about the status of this API")

	connectorId := testutils.CreateConnector(t)
	details, err := testutils.Client.NewMetadataColumnsList().ConnectorId(connectorId).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", details)
		t.Error(err)
	}

	testutils.AssertEqual(t, details.Code, "Success")
	testutils.AssertNotEmpty(t, details.Data.Items[0].Id)
	testutils.AssertNotEmpty(t, details.Data.Items[0].ParentId)
	testutils.AssertNotEmpty(t, details.Data.Items[0].NameInSource)
	testutils.AssertNotEmpty(t, details.Data.Items[0].NameInDestination)
	testutils.AssertNotEmpty(t, details.Data.Items[0].TypeInSource)
	testutils.AssertNotEmpty(t, details.Data.Items[0].TypeInDestination)
	testutils.AssertNotEmpty(t, details.Data.Items[0].IsPrimaryKey)
	testutils.AssertNotEmpty(t, details.Data.Items[0].IsForeignKey)

	t.Cleanup(func() { testutils.DeleteConnector(t, connectorId) })
}
