package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestProxyConnectionMembershipsListE2E(t *testing.T) {
	proxyId := testutils.CreateProxy(t)
	connectionId := testutils.CreateConnector(t)
	testutils.CreateProxyConnection(t, proxyId, connectionId)

	result, err := testutils.Client.NewProxyConnectionMembershipsList().ProxyId(proxyId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertEqual(t, result.Data.Items[0].ConnectionId, connectionId)

	t.Cleanup(func() {
		testutils.DeleteProxyConnection(t, proxyId, connectionId)
		testutils.DeleteConnector(t, connectionId)
		testutils.DeleteProxy(t, proxyId)
	})
}
