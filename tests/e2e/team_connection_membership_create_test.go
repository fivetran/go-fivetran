package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewTeamConnectionMembershipCreateE2E(t *testing.T) {
	teamId := testutils.CreateTeam(t)
	ConnectionId := testutils.CreateConnection(t)

	created, err := testutils.Client.NewTeamConnectionMembershipCreate().
		TeamId(teamId).
		ConnectionId(ConnectionId).
		Role("Connector Administrator").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertNotEmpty(t, created.Message)
	testutils.AssertEqual(t, created.Data.ConnectionId, ConnectionId)
	testutils.AssertEqual(t, created.Data.Role, "Connector Administrator")
	testutils.AssertNotEmpty(t, created.Data.CreatedAt)

	t.Cleanup(func() {
		testutils.DeleteTeamConnection(t, teamId, ConnectionId)
		testutils.DeleteConnection(t, ConnectionId)
		testutils.DeleteTeam(t, teamId)
	})
}
