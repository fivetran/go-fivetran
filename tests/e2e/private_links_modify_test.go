package fivetran_test

import (
	"context"
	"testing"

	"github.com/fivetran/go-fivetran"
	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewPrivateLinkModifyE2E(t *testing.T) {
	t.Skip("Private links have a strict limit on the number of requests per hour; to test changes in these modules, this Skip must be removed")

	privateLinkId, plGroupId := testutils.CreateTempPrivateLink(t)
	details, err := testutils.Client.NewPrivateLinksModify().PrivateLinkId(privateLinkId).
		Config(fivetran.NewPrivateLinksConfig().
			ConnectionServiceId("2")).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", details)
		t.Error(err)
	}

	testutils.AssertEqual(t, details.Code, "Success")
	testutils.AssertNotEmpty(t, details.Message)

	testutils.AssertEqual(t, details.Data.GroupId, plGroupId)
	testutils.AssertEqual(t, details.Data.Config.ConnectionServiceId, "2")
}