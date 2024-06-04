package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewProxyCreateE2E(t *testing.T) {
	created, err := testutils.Client.NewProxyCreate().
		DisplayName("go_sdk_test_proxy").
		GroupRegion("GCP_US_EAST4").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertNotEmpty(t, created.Message)
	testutils.AssertNotEmpty(t, created.Data.AgentId)
	testutils.AssertNotEmpty(t, created.Data.AuthToken)
	testutils.AssertNotEmpty(t, created.Data.ProxyServerUri)

	t.Cleanup(func() { 
		testutils.DeleteProxy(t, created.Data.AgentId) 
	})
}
