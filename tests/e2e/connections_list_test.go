package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewConnectionsListE2E(t *testing.T) {
	ConnectionId := testutils.CreateTempConnection(t)
	result, err := testutils.Client.NewConnectionsList().Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertHasLength(t, result.Data.Items, 1)
	testutils.AssertEmpty(t, result.Message)
	testutils.AssertEqual(t, result.Data.Items[0].ID, ConnectionId)
	testutils.AssertEqual(t, result.Data.Items[0].GroupID, testutils.PredefinedGroupId)
	testutils.AssertEqual(t, result.Data.Items[0].Service, "itunes_connect")
	testutils.AssertEqual(t, result.Data.Items[0].Schema, "itunes_e2e_connect")

	testutils.AssertEmpty(t, result.Data.NextCursor)
}

func TestNewConnectionsListFilterByGroupAndSchemaE2E(t *testing.T) {
	ConnectionId := testutils.CreateTempConnection(t)
	result, err := testutils.Client.NewConnectionsList().
		GroupID(testutils.PredefinedGroupId).
		Schema("itunes_e2e_connect").
		Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertHasLength(t, result.Data.Items, 1)
	testutils.AssertEmpty(t, result.Message)
	testutils.AssertEqual(t, result.Data.Items[0].ID, ConnectionId)
	testutils.AssertEqual(t, result.Data.Items[0].GroupID, testutils.PredefinedGroupId)
	testutils.AssertEqual(t, result.Data.Items[0].Service, "itunes_connect")
	testutils.AssertEqual(t, result.Data.Items[0].Schema, "itunes_e2e_connect")

	testutils.AssertEmpty(t, result.Data.NextCursor)
}
