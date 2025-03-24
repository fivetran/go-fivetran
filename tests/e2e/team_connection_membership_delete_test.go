package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewTeamConnectionMebershipDeleteE2E(t *testing.T) {
	teamId := testutils.CreateTeam(t)
	ConnectionId := testutils.CreateConnection(t)
	testutils.CreateTeamConnection(t, teamId, ConnectionId)

	deleted, err := testutils.Client.NewTeamConnectionMembershipDelete().
		TeamId(teamId).
		ConnectionId(ConnectionId).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}

	testutils.AssertEqual(t, deleted.Code, "Success")
	testutils.AssertNotEmpty(t, deleted.Message)

	t.Cleanup(func() {
		testutils.DeleteConnection(t, ConnectionId)
		testutils.DeleteTeam(t, teamId)
	})
}
