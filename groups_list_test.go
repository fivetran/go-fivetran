package fivetran_test

import (
	"context"
	"testing"
)

func TestNewGroupsListE2E(t *testing.T) {
	result, err := Client.NewGroupsList().Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	AssertEqual(t, result.Code, "Success")
	AssertHasLength(t, result.Data.Items, 1)
	AssertEmpty(t, result.Message)
	AssertEqual(t, result.Data.Items[0].ID, PredefinedGroupId)
	AssertEqual(t, result.Data.Items[0].Name, "Warehouse")
	AssertEqual(t, result.Data.Items[0].CreatedAt.IsZero(), false)
	AssertEmpty(t, result.Data.NextCursor)
}
