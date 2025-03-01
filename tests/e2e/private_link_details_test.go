package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestPrivateLinkDetailsE2E(t *testing.T) {
	t.Skip("Passed in previous runs. Private links have a strict limit on the number of requests per hour; to test changes in these modules, this Skip must be removed")

	linkId := testutils.CreatePrivateLink(t)

	result, err := testutils.Client.NewPrivateLinkDetails().PrivateLinkId(linkId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertEqual(t, result.Data.Id, linkId)
	testutils.AssertNotEmpty(t, result.Data.Name)
	testutils.AssertNotEmpty(t, result.Data.AccountId)
	testutils.AssertEqual(t, result.Data.Region, "GCP_US_EAST4")
	testutils.AssertEqual(t, result.Data.Service, "SOURCE_GCP")
	testutils.AssertNotEmpty(t, result.Data.CreatedAt)
	testutils.AssertNotEmpty(t, result.Data.CreatedBy)

	t.Cleanup(func() { testutils.DeletePrivateLink(t, linkId) })
}
