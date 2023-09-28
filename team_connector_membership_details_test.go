package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestTeamConnectorMembershipDetailsE2E(t *testing.T) {
	teamId := testutils.CreateTeam(t)
	connectorId := testutils.CreateConnector(t)
	testutils.CreateTeamConnector(t, teamId, connectorId)

	result, err := testutils.Client.NewTeamConnectorMembershipDetails().
		TeamId(teamId).
		ConnectorId(connectorId).
		Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertEqual(t, result.Data.ConnectorId, connectorId)
	testutils.AssertEqual(t, result.Data.Role, "Connector Administrator")
	testutils.AssertNotEmpty(t, result.Data.CreatedAt)

	t.Cleanup(func() {
		testutils.DeleteTeamConnector(t, teamId, connectorId)
		testutils.DeleteConnector(t, connectorId)
		testutils.DeleteTeam(t, teamId)
	})
}
