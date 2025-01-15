package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestPrivateLinkListE2E(t *testing.T) {
	t.Skip("Passed in previous runs. Private links have a strict limit on the number of requests per hour; to test changes in these modules, this Skip must be removed")

	linkId := testutils.CreatePrivateLink(t)

	result, err := testutils.Client.NewPrivateLinkList().Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	for _, v := range result.Data.Items {
		if v.Id == linkId {
			testutils.AssertNotEmpty(t, v.AccountId)
			testutils.AssertEqual(t, v.Region, "GCP_US_EAST4")
			testutils.AssertEqual(t, v.Service, "SOURCE_GCP")
			testutils.AssertNotEmpty(t, v.CreatedAt)
			testutils.AssertNotEmpty(t, v.CreatedBy)
			testutils.AssertNotEmpty(t, v.Name)
		}
    }

	t.Cleanup(func() { testutils.DeletePrivateLink(t, linkId) })
}
