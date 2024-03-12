package fivetran_test

import (
	"context"
	"testing"

	"github.com/fivetran/go-fivetran"
	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewPrivateLinkCreateE2E(t *testing.T) {
	t.Skip("Private links have a strict limit on the number of requests per hour; to test changes in these modules, this Skip must be removed")

	plGroupId := testutils.CreatePrivateLinkGroup(t)
	plDestinationId := testutils.CreatePrivateLinkDestination(t, plGroupId)

	created, err := testutils.Client.NewPrivateLinksCreate().
		Name("test").
		GroupId(plGroupId).
		Config(fivetran.NewPrivateLinksConfig().
			ConnectionServiceId("1")).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertNotEmpty(t, created.Message)
	testutils.AssertEqual(t, created.Data.Name, "test")
	testutils.AssertEqual(t, created.Data.GroupId, plGroupId)
	testutils.AssertEqual(t, created.Data.Config.ConnectionServiceId, "1")

	t.Cleanup(func() { 
		testutils.DeletePrivateLink(t, created.Data.Id) 
		testutils.DeleteDestination(t, plDestinationId) 
		testutils.DeleteGroup(t, plGroupId) 
	})
}
