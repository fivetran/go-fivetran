package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewPrivateLinkDeleteE2E(t *testing.T) {
	t.Skip("Passed in previous runs. Private links have a strict limit on the number of requests per hour; to test changes in these modules, this Skip must be removed")

	privateLinkId := testutils.CreatePrivateLink(t)
	
	deleted, err := testutils.Client.NewPrivateLinkDelete().PrivateLinkId(privateLinkId).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}

	testutils.AssertEqual(t, deleted.Code, "Success")
	testutils.AssertNotEmpty(t, deleted.Message)
}
