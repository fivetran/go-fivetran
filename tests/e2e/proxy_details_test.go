package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestProxyDetailsE2E(t *testing.T) {
	proxyId := testutils.CreateProxy(t)

	result, err := testutils.Client.NewProxyDetails().ProxyId(proxyId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertEqual(t, result.Data.Id, proxyId)
	testutils.AssertNotEmpty(t, result.Data.AccountId)
	testutils.AssertNotEmpty(t, result.Data.Region)
	testutils.AssertNotEmpty(t, result.Data.CreatedBy)
	testutils.AssertNotEmpty(t, result.Data.DisplayName)

	t.Cleanup(func() { testutils.DeleteProxy(t, proxyId) })
}
