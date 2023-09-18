package fivetran_test

import (
    "context"
    "testing"
)

func TestTeamConnectorMembershipDetailsE2E(t *testing.T) {
    teamId := CreateTeam(t)
    connectorId := CreateConnector(t)
    CreateTeamConnector(t, teamId, connectorId)

    result, err := Client.NewTeamConnectorMembershipDetails().
        TeamId(teamId).
        ConnectorId(connectorId).
        Do(context.Background())
    if err != nil {
        t.Logf("%+v\n", result)
        t.Error(err)
    }

    AssertEqual(t, result.Code, "Success")
    AssertEqual(t, result.Data.ConnectorId, connectorId)
    AssertEqual(t, result.Data.Role, "Connector Administrator")
    AssertNotEmpty(t, result.Data.CreatedAt)

    t.Cleanup(func() { 
        DeleteTeamConnector(t, teamId, connectorId)
        DeleteConnector(t, connectorId)
        DeleteTeam(t, teamId)
    })
}
