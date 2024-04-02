package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewProxyConnectionMembershipCreateE2E(t *testing.T) {
	proxyId := testutils.CreateProxy(t)
	connectionId := testutils.CreateConnector(t)

	created, err := testutils.Client.NewProxyConnectionMembershipCreate().
		ProxyId(proxyId).
		ConnectionId(connectionId).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertNotEmpty(t, created.Message)

	t.Cleanup(func() {
		testutils.DeleteProxyConnection(t, proxyId, connectionId)
		testutils.DeleteConnector(t, connectionId)
		testutils.DeleteProxy(t, proxyId)
	})
}
