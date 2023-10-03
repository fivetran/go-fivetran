package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewTeamUserMembershipCreateE2E(t *testing.T) {
	teamId := testutils.CreateTeam(t)

	created, err := testutils.Client.NewTeamUserMembershipCreate().
		TeamId(teamId).
		UserId(testutils.PredefinedUserId).
		Role("Team Member").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertEqual(t, created.Message, "User has been added to the team")
	testutils.AssertEqual(t, created.Data.UserId, testutils.PredefinedUserId)
	testutils.AssertEqual(t, created.Data.Role, "Team Member")

	t.Cleanup(func() {
		testutils.DeleteTeamUser(t, teamId, testutils.PredefinedUserId)
		testutils.DeleteTeam(t, teamId)
	})
}
