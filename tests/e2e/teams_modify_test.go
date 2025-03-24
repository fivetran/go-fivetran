package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewTeamsUpdateE2E(t *testing.T) {
	teamId := testutils.CreateTeam(t)

	modified, err := testutils.Client.NewTeamsUpdate().
		TeamId(teamId).
		Name("test_team").
		Description("test_description_2").
		Role("Account Reviewer").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", modified)
		t.Error(err)
	}

	testutils.AssertEqual(t, modified.Code, "Success")
	testutils.AssertNotEmpty(t, modified.Message)
	testutils.AssertNotEmpty(t, modified.Data.Id)
	testutils.AssertEqual(t, modified.Data.Name, "test_team")
	testutils.AssertEqual(t, modified.Data.Description, "test_description_2")
	testutils.AssertEqual(t, modified.Data.Role, "Account Reviewer")

	t.Cleanup(func() { testutils.DeleteTeam(t, modified.Data.Id) })
}
