package fivetran_test

import (
	"context"
	"strings"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewConnectorSyncE2E(t *testing.T) {
	connectorId := testutils.CreateTempConnector(t)
	sync, err := testutils.Client.NewConnectorSync().
		ConnectorID(connectorId).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", sync)
		t.Error(err)
	}

	testutils.AssertEqual(t, sync.Code, "Success")
	testutils.AssertNotEmpty(t, sync.Message)
	testutils.AssertEqual(t, strings.Contains(sync.Message, connectorId), true)
}
