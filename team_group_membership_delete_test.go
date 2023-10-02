package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewTeamGroupMembershipDeleteE2E(t *testing.T) {
	groupId := testutils.CreateGroup(t)
	teamId := testutils.CreateTeam(t)
	testutils.CreateTeamGroup(t, teamId, groupId)

	deleted, err := testutils.Client.NewTeamGroupMembershipDelete().
		TeamId(teamId).
		GroupId(groupId).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}

	testutils.AssertEqual(t, deleted.Code, "Success")
	testutils.AssertEqual(t, deleted.Message, "Group membership has been deleted")

	t.Cleanup(func() {
		testutils.DeleteGroup(t, groupId)
		testutils.DeleteTeam(t, teamId)
	})
}
