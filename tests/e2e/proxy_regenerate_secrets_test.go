package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewProxyRegenerateE2E(t *testing.T) {
	proxyId := testutils.CreateProxy(t)

	created, err := testutils.Client.NewProxyRegenerateSecrets().ProxyId(proxyId).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertNotEmpty(t, created.Message)
	testutils.AssertNotEmpty(t, created.Data.AgentId)
	testutils.AssertNotEmpty(t, created.Data.AuthToken)
	testutils.AssertNotEmpty(t, created.Data.ClientCert)
	testutils.AssertNotEmpty(t, created.Data.ClientPrivateKey)
}