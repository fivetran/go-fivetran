package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewGroupModifyE2E(t *testing.T) {
	groupId := testutils.CreateTempGroup(t)

	updated, err := testutils.Client.NewGroupModify().GroupID(groupId).Name("test_updated").Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", updated)
		t.Error(err)
	}

	testutils.AssertEqual(t, updated.Code, "Success")
	testutils.AssertEqual(t, updated.Data.ID, groupId)
	testutils.AssertEqual(t, updated.Data.Name, "test_updated")
	testutils.AssertEqual(t, updated.Data.CreatedAt.IsZero(), false)
}
