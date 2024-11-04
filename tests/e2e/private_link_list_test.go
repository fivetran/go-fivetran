package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestPrivateLinkListE2E(t *testing.T) {
	//t.Skip("Private links have a strict limit on the number of requests per hour; to test changes in these modules, this Skip must be removed")

	linkId := testutils.CreatePrivateLink(t)

	result, err := testutils.Client.NewPrivateLinkList().Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertEqual(t, result.Data.Items[0].Id, linkId)
	testutils.AssertNotEmpty(t, result.Data.Items[0].AccountId)
	testutils.AssertNotEmpty(t, result.Data.Items[0].Region)
	testutils.AssertNotEmpty(t, result.Data.Items[0].Service)
	testutils.AssertNotEmpty(t, result.Data.Items[0].CreatedAt)
	testutils.AssertNotEmpty(t, result.Data.Items[0].CreatedBy)
	testutils.AssertNotEmpty(t, result.Data.Items[0].Name)

	t.Cleanup(func() { testutils.DeletePrivateLink(t, linkId) })
}
