package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestTeamUserMembershipDetailsE2E(t *testing.T) {
	teamId := testutils.CreateTeam(t)
	testutils.CreateTeamUser(t, teamId, testutils.PredefinedUserId)

	result, err := testutils.Client.NewTeamUserMembershipDetails().
		TeamId(teamId).
		UserId(testutils.PredefinedUserId).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertEqual(t, result.Data.UserId, testutils.PredefinedUserId)
	testutils.AssertEqual(t, result.Data.Role, "Team Member")

	t.Cleanup(func() {
		testutils.DeleteTeamUser(t, teamId, testutils.PredefinedUserId)
		testutils.DeleteTeam(t, teamId)
	})
}
