package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewTeamConnectorMebershipDeleteE2E(t *testing.T) {
	teamId := testutils.CreateTeam(t)
	connectorId := testutils.CreateConnector(t)
	testutils.CreateTeamConnector(t, teamId, connectorId)

	deleted, err := testutils.Client.NewTeamConnectorMembershipDelete().
		TeamId(teamId).
		ConnectorId(connectorId).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}

	testutils.AssertEqual(t, deleted.Code, "Success")
	testutils.AssertEqual(t, deleted.Message, "Connector membership has been deleted")

	t.Cleanup(func() {
		testutils.DeleteConnector(t, connectorId)
		testutils.DeleteTeam(t, teamId)
	})
}
