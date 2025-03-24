package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewTeamGroupMembershipUpdateE2E(t *testing.T) {
	teamId := testutils.CreateTeam(t)
	groupId := testutils.CreateGroup(t)
	testutils.CreateTeamGroup(t, teamId, groupId)

	modified, err := testutils.Client.NewTeamGroupMembershipUpdate().
		TeamId(teamId).
		GroupId(groupId).
		Role("Destination Reviewer").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", modified)
		t.Error(err)
	}

	testutils.AssertEqual(t, modified.Code, "Success")
	testutils.AssertEqual(t, modified.Message, "Group membership has been updated")

	t.Cleanup(func() {
		testutils.DeleteTeamGroup(t, teamId, groupId)
		testutils.DeleteGroup(t, groupId)
		testutils.DeleteTeam(t, teamId)
	})
}
