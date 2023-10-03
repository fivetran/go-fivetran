package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewTeamGroupMembershipCreateE2E(t *testing.T) {
	groupId := testutils.CreateGroup(t)
	teamId := testutils.CreateTeam(t)

	created, err := testutils.Client.NewTeamGroupMembershipCreate().
		TeamId(teamId).
		GroupId(groupId).
		Role("Destination Administrator").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertEqual(t, created.Message, "Group membership has been created")
	testutils.AssertEqual(t, created.Data.GroupId, groupId)
	testutils.AssertEqual(t, created.Data.Role, "Destination Administrator")
	testutils.AssertNotEmpty(t, created.Data.CreatedAt)

	t.Cleanup(func() {
		testutils.DeleteTeamGroup(t, teamId, groupId)
		testutils.DeleteGroup(t, groupId)
		testutils.DeleteTeam(t, teamId)
	})
}
