package fivetran_test

import (
	"context"
	"testing"

	"github.com/fivetran/go-fivetran"
	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewPrivateLinkModifyE2E(t *testing.T) {
	//t.Skip("Private links have a strict limit on the number of requests per hour; to test changes in these modules, this Skip must be removed")

	privateLinkId := testutils.CreateTempPrivateLink(t)

	details, err := testutils.Client.NewPrivateLinkModify().PrivateLinkId(privateLinkId).
		Config(fivetran.NewPrivateLinkConfig().
			PrivateConnectionServiceId("test2")).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", details)
		if details.Code != "IllegalState" {
			t.Error(err)
		}
	}

	// Private link in CREATING state is not allowed to perform any operation
	testutils.AssertEqual(t, details.Code, "IllegalState")
	testutils.AssertNotEmpty(t, details.Message)
}