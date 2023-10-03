package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewGroupDetailsE2E(t *testing.T) {
	result, err := testutils.Client.NewGroupDetails().GroupID(testutils.PredefinedGroupId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertEqual(t, result.Data.ID, testutils.PredefinedGroupId)
	testutils.AssertEqual(t, result.Data.Name, testutils.PredefinedGroupName)
	testutils.AssertEqual(t, result.Data.CreatedAt.IsZero(), false)
}
