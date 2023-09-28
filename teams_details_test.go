package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestTeamsDetailsE2E(t *testing.T) {
	teamId := testutils.CreateTeam(t)

	result, err := testutils.Client.NewTeamsDetails().TeamId(teamId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertEqual(t, result.Data.Id, teamId)
	testutils.AssertEqual(t, result.Data.Name, "test_team")
	testutils.AssertEqual(t, result.Data.Description, "test_description")
	testutils.AssertEqual(t, result.Data.Role, "Account Reviewer")

	t.Cleanup(func() { testutils.DeleteTeam(t, teamId) })
}
