package fivetran_test

import (
	"context"
	"testing"
)

func TestNewGroupDetailsE2E(t *testing.T) {
	result, err := Client.NewGroupDetails().GroupID(PredefinedGroupId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	AssertEqual(t, result.Code, "Success")
	AssertEqual(t, result.Data.ID, PredefinedGroupId)
	AssertEqual(t, result.Data.Name, "Warehouse")
	AssertEqual(t, result.Data.CreatedAt.IsZero(), false)
}
