package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewTeamsCreateE2E(t *testing.T) {
	created, err := testutils.Client.NewTeamsCreate().
		Name("test_team").
		Description("test_description").
		Role("Account Reviewer").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertNotEmpty(t, created.Message)
	testutils.AssertNotEmpty(t, created.Data.Id)
	testutils.AssertEqual(t, created.Data.Name, "test_team")
	testutils.AssertEqual(t, created.Data.Description, "test_description")
	testutils.AssertEqual(t, created.Data.Role, "Account Reviewer")

	t.Cleanup(func() { testutils.DeleteTeam(t, created.Data.Id) })
}
