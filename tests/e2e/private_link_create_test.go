package fivetran_test

import (
	"context"
	"testing"

	"github.com/fivetran/go-fivetran"
	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewPrivateLinkCreateE2E(t *testing.T) {
	//t.Skip("Private links have a strict limit on the number of requests per hour; to test changes in these modules, this Skip must be removed")

	created, err := testutils.Client.NewPrivateLinkCreate().
		Name("sdk_private_link_test").
		Service("SOURCE_GCP").
		Region("GCP_US_EAST4").
		Config(fivetran.NewPrivateLinkConfig().
			PrivateConnectionServiceId("test")).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertNotEmpty(t, created.Message)
	testutils.AssertEqual(t, created.Data.Name, "sdk_private_link_test")
	testutils.AssertNotEmpty(t, created.Data.Service)
	testutils.AssertNotEmpty(t, created.Data.Region)
	testutils.AssertNotEmpty(t, created.Data.CloudProvider)
	testutils.AssertEqual(t, created.Data.Config.PrivateConnectionServiceId, "test")

	t.Cleanup(func() { 
		testutils.DeletePrivateLink(t, created.Data.Id) 
	})
}

func TestNewPrivateLinkCustomCreateE2E(t *testing.T) {
	//t.Skip("Private links have a strict limit on the number of requests per hour; to test changes in these modules, this Skip must be removed")

	created, err := testutils.Client.NewPrivateLinkCreate().
		Name("sdk_private_link_test").
		Service("SOURCE_GCP").
		Region("GCP_US_EAST4").
		ConfigCustom(&map[string]interface{}{
			"private_connection_service_id": "test",
		}).
		DoCustom(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertNotEmpty(t, created.Message)
	testutils.AssertEqual(t, created.Data.Name, "sdk_private_link_test")
	testutils.AssertNotEmpty(t, created.Data.Service)
	testutils.AssertNotEmpty(t, created.Data.Region)
	testutils.AssertNotEmpty(t, created.Data.CloudProvider)
	testutils.AssertEqual(t, created.Data.Config["private_connection_service_id"], "test")

	t.Cleanup(func() { 
		testutils.DeletePrivateLink(t, created.Data.Id) 
	})
}