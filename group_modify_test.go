package fivetran_test

import (
	"context"
	"testing"
)

func TestNewGroupModifyE2E(t *testing.T) {
	groupId := CreateTempGroup(t)

	updated, err := Client.NewGroupModify().GroupID(groupId).Name("test_updated").Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", updated)
		t.Error(err)
	}

	AssertEqual(t, updated.Code, "Success")
	AssertEqual(t, updated.Data.ID, groupId)
	AssertEqual(t, updated.Data.Name, "test_updated")
	AssertEqual(t, updated.Data.CreatedAt.IsZero(), false)
}
