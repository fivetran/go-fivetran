package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestProxyListE2E(t *testing.T) {
	proxyId := testutils.CreateProxy(t)

	result, err := testutils.Client.NewProxyList().Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertEqual(t, result.Data.Items[0].Id, proxyId)
	testutils.AssertNotEmpty(t, result.Data.Items[0].AccountId)
	testutils.AssertNotEmpty(t, result.Data.Items[0].Region)
	testutils.AssertNotEmpty(t, result.Data.Items[0].Token)
	testutils.AssertNotEmpty(t, result.Data.Items[0].Salt)
	testutils.AssertNotEmpty(t, result.Data.Items[0].CreatedBy)
	testutils.AssertNotEmpty(t, result.Data.Items[0].DisplayName)

	t.Cleanup(func() { testutils.DeleteProxy(t, proxyId) })
}
