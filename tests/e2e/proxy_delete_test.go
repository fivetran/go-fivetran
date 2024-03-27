package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewProxyDeleteE2E(t *testing.T) {
	proxyId := testutils.CreateProxy(t)

	deleted, err := testutils.Client.NewProxyDelete().ProxyId(proxyId).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}

	testutils.AssertEqual(t, deleted.Code, "Success")
	testutils.AssertNotEmpty(t, deleted.Message)
}
