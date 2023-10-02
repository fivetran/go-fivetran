package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestTeamsListE2E(t *testing.T) {
	teamId := testutils.CreateTeam(t)

	result, err := testutils.Client.NewTeamsList().Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertEqual(t, result.Data.Items[0].Id, teamId)
	testutils.AssertEqual(t, result.Data.Items[0].Name, "test_team")
	testutils.AssertEqual(t, result.Data.Items[0].Description, "test_description")
	testutils.AssertEqual(t, result.Data.Items[0].Role, "Account Reviewer")

	t.Cleanup(func() { testutils.DeleteTeam(t, teamId) })
}
