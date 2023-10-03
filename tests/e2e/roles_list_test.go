package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewRolesListE2E(t *testing.T) {
	result, err := testutils.Client.NewRolesList().Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertNotEmpty(t, result.Data.Items[0].Name)
	testutils.AssertNotEmpty(t, result.Data.Items[0].Description)
	testutils.AssertNotEmpty(t, result.Data.Items[0].Scope)
}
