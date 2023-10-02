package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestTeamGroupMembershipDetailsE2E(t *testing.T) {
	groupId := testutils.CreateGroup(t)
	teamId := testutils.CreateTeam(t)
	testutils.CreateTeamGroup(t, teamId, groupId)

	result, err := testutils.Client.NewTeamGroupMembershipDetails().
		TeamId(teamId).
		GroupId(groupId).
		Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertEqual(t, result.Data.GroupId, groupId)
	testutils.AssertEqual(t, result.Data.Role, "Destination Analyst")
	testutils.AssertNotEmpty(t, result.Data.CreatedAt)

	t.Cleanup(func() {
		testutils.DeleteTeamGroup(t, teamId, groupId)
		testutils.DeleteGroup(t, groupId)
		testutils.DeleteTeam(t, teamId)
	})
}
