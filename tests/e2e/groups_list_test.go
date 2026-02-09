package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewGroupsListE2E(t *testing.T) {
	result, err := testutils.Client.NewGroupsList().Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertHasLength(t, result.Data.Items, 1)
	testutils.AssertEqual(t, result.Message, "Groups retrieved successfully")
	testutils.AssertEqual(t, result.Data.Items[0].ID, testutils.PredefinedGroupId)
	testutils.AssertEqual(t, result.Data.Items[0].Name, testutils.PredefinedGroupName)
	testutils.AssertEqual(t, result.Data.Items[0].CreatedAt.IsZero(), false)
	testutils.AssertEmpty(t, result.Data.NextCursor)
}
