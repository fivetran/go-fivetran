package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewGroupListPrivateLinksE2E(t *testing.T) {
	t.Skip("Private links have a strict limit on the number of requests per hour; to test changes in these modules, this Skip must be removed")

	privateLinkId, groupId := testutils.CreateTempPrivateLink(t)
	privateLinks, err := testutils.Client.NewGroupListPrivateLinks().GroupID(groupId).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", privateLinks)
		t.Error(err)
	}

	testutils.AssertEqual(t, privateLinks.Code, "Success")
	testutils.AssertHasLength(t, privateLinks.Data.Items, 1)

	testutils.AssertEqual(t, privateLinks.Data.Items[0].Name, "test")
	testutils.AssertEqual(t, privateLinks.Data.Items[0].Id, privateLinkId)
	testutils.AssertEqual(t, privateLinks.Data.Items[0].GroupId, groupId)
}