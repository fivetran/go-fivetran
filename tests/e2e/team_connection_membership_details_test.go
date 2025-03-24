package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestTeamConnectionMembershipDetailsE2E(t *testing.T) {
	teamId := testutils.CreateTeam(t)
	ConnectionId := testutils.CreateConnection(t)
	testutils.CreateTeamConnection(t, teamId, ConnectionId)

	result, err := testutils.Client.NewTeamConnectionMembershipDetails().
		TeamId(teamId).
		ConnectionId(ConnectionId).
		Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertEqual(t, result.Data.ConnectionId, ConnectionId)
	testutils.AssertEqual(t, result.Data.Role, "Connection Administrator")
	testutils.AssertNotEmpty(t, result.Data.CreatedAt)

	t.Cleanup(func() {
		testutils.DeleteTeamConnection(t, teamId, ConnectionId)
		testutils.DeleteConnection(t, ConnectionId)
		testutils.DeleteTeam(t, teamId)
	})
}
