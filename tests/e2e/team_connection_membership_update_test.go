package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewTeamConnectionMembershipUpdateE2E(t *testing.T) {
	teamId := testutils.CreateTeam(t)
	ConnectionId := testutils.CreateConnection(t)
	testutils.CreateTeamConnection(t, teamId, ConnectionId)

	modified, err := testutils.Client.NewTeamConnectionMembershipUpdate().
		TeamId(teamId).
		ConnectionId(ConnectionId).
		Role("Connection Collaborator").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", modified)
		t.Error(err)
	}

	testutils.AssertEqual(t, modified.Code, "Success")
	testutils.AssertNotEmpty(t, modified.Message)

	t.Cleanup(func() {
		testutils.DeleteTeamConnection(t, teamId, ConnectionId)
		testutils.DeleteConnection(t, ConnectionId)
		testutils.DeleteTeam(t, teamId)
	})
}
