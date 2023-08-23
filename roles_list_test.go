package fivetran_test

import (
	"context"
	"testing"
)

func TestNewRolesListE2E(t *testing.T) {
	result, err := Client.NewRolesList().Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	AssertEqual(t, result.Code, "Success")
	AssertNotEmpty(t, result.Data.Items[0].Name)
	AssertNotEmpty(t, result.Data.Items[0].Description)
	AssertNotEmpty(t, result.Data.Items[0].Scope)
}