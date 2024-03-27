package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestProxyConnectionMembershipDeleteE2E(t *testing.T) {
	proxyId := testutils.CreateProxy(t)
	connectionId := testutils.CreateConnector(t)
	testutils.CreateProxyConnection(t, proxyId, connectionId)

	deleted, err := testutils.Client.NewProxyConnectionMembershipDelete().
		ProxyId(proxyId).
		ConnectionId(connectionId).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}

	testutils.AssertEqual(t, deleted.Code, "Success")
	testutils.AssertNotEmpty(t, deleted.Message)

	t.Cleanup(func() {
		testutils.DeleteConnector(t, connectionId)
		testutils.DeleteProxy(t, proxyId)
	})
}
