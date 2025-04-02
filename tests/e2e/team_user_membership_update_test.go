package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewTeamUserMembershipUpdateE2E(t *testing.T) {
	teamId := testutils.CreateTeam(t)
	testutils.CreateTeamUser(t, teamId, testutils.PredefinedUserId)

	modified, err := testutils.Client.NewTeamUserMembershipUpdate().
		TeamId(teamId).
		UserId(testutils.PredefinedUserId).
		Role("Team Manager").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", modified)
		t.Error(err)
	}

	testutils.AssertEqual(t, modified.Code, "Success")
	testutils.AssertEqual(t, modified.Message, "User role has been updated")

	t.Cleanup(func() {
		testutils.DeleteTeamUser(t, teamId, testutils.PredefinedUserId)
		testutils.DeleteTeam(t, teamId)
	})
}
