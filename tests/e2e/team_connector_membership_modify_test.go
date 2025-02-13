package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewTeamConnectorMembershipModifyE2E(t *testing.T) {
	teamId := testutils.CreateTeam(t)
	connectorId := testutils.CreateConnector(t)
	testutils.CreateTeamConnector(t, teamId, connectorId)

	modified, err := testutils.Client.NewTeamConnectorMembershipModify().
		TeamId(teamId).
		ConnectorId(connectorId).
		Role("Connector Collaborator").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", modified)
		t.Error(err)
	}

	testutils.AssertEqual(t, modified.Code, "Success")
	testutils.AssertNotEmpty(t, modified.Message)

	t.Cleanup(func() {
		testutils.DeleteTeamConnector(t, teamId, connectorId)
		testutils.DeleteConnector(t, connectorId)
		testutils.DeleteTeam(t, teamId)
	})
}
