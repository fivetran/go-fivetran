package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewTeamUserMembershipDeleteE2E(t *testing.T) {
	teamId := testutils.CreateTeam(t)
	testutils.CreateTeamUser(t, teamId, testutils.PredefinedUserId)

	deleted, err := testutils.Client.NewTeamUserMembershipDelete().
		TeamId(teamId).
		UserId(testutils.PredefinedUserId).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}

	testutils.AssertEqual(t, deleted.Code, "Success")
	testutils.AssertEqual(t, deleted.Message, "User has been removed from the team")

	t.Cleanup(func() {
		testutils.DeleteTeam(t, teamId)
	})
}
