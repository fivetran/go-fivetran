package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewGroupListPrivateLinksE2E(t *testing.T) {
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
	testutils.AssertEqual(t, privateLinks.Data.Items[0].Service, "big_query") // It's a bug in PL service - in field service returned service from destination
}