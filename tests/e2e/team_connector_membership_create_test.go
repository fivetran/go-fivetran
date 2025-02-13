package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewTeamConnectorMembershipCreateE2E(t *testing.T) {
	teamId := testutils.CreateTeam(t)
	connectorId := testutils.CreateConnector(t)

	created, err := testutils.Client.NewTeamConnectorMembershipCreate().
		TeamId(teamId).
		ConnectorId(connectorId).
		Role("Connector Administrator").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertNotEmpty(t, created.Message)
	testutils.AssertEqual(t, created.Data.ConnectorId, connectorId)
	testutils.AssertEqual(t, created.Data.Role, "Connector Administrator")
	testutils.AssertNotEmpty(t, created.Data.CreatedAt)

	t.Cleanup(func() {
		testutils.DeleteTeamConnector(t, teamId, connectorId)
		testutils.DeleteConnector(t, connectorId)
		testutils.DeleteTeam(t, teamId)
	})
}
