package fivetran_test

import (
	"context"
	"strings"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewConnectionSyncE2E(t *testing.T) {
	t.Skip("Problems with scheduler on staging")
	ConnectionId := testutils.CreateTempConnection(t)
	sync, err := testutils.Client.NewConnectionSync().
		ConnectionID(ConnectionId).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", sync)
		t.Error(err)
	}

	testutils.AssertEqual(t, sync.Code, "Success")
	testutils.AssertNotEmpty(t, sync.Message)
	testutils.AssertEqual(t, strings.Contains(sync.Message, ConnectionId), true)
}
