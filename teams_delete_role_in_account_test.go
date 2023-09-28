package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewTeamRoleDeleteE2E(t *testing.T) {
	teamId := testutils.CreateTeam(t)

	deleted, err := testutils.Client.NewTeamsDeleteRoleInAccount().TeamId(teamId).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}

	testutils.AssertEqual(t, deleted.Code, "Success")
	testutils.AssertEqual(t, deleted.Message, "Team role in account has been removed")

	t.Cleanup(func() { testutils.DeleteTeam(t, teamId) })
}
