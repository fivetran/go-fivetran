package fivetran_test

import (
    "context"
    "testing"
)

func TestNewTeamConnectorMembershipCreateE2E(t *testing.T) {
    teamId := CreateTeam(t)
    connectorId := CreateConnector(t)

    created, err := Client.NewTeamConnectorMembershipCreate().
        TeamId(teamId).
        ConnectorId(connectorId).
        Role("Connector Administrator").
        Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", created)
        t.Error(err)
    }

    AssertEqual(t, created.Code, "Success")
    AssertEqual(t, created.Message, "Connector membership has been created")    
    AssertEqual(t, created.Data.ConnectorId, connectorId)
    AssertEqual(t, created.Data.Role, "Connector Administrator")
    AssertNotEmpty(t, created.Data.CreatedAt)
    
    t.Cleanup(func() { 
        DeleteTeamConnector(t, teamId, connectorId)
        DeleteConnector(t, connectorId)
        DeleteTeam(t, teamId)
    })
}
