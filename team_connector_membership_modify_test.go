package fivetran_test

import (
    "context"
    "testing"
)

func TestNewTeamConnectorMembershipModifyE2E(t *testing.T) {
    teamId := CreateTeam(t)
    connectorId := CreateConnector(t)
    CreateTeamConnector(t, teamId, connectorId)

    modified, err := Client.NewTeamConnectorMembershipModify().
        TeamId(teamId).
        ConnectorId(connectorId).
        Role("Connector Collaborator").
        Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", modified)
        t.Error(err)
    }

    AssertEqual(t, modified.Code, "Success")
    AssertEqual(t, modified.Message, "Connector membership has been updated")
    
    t.Cleanup(func() { 
        DeleteTeamConnector(t, teamId, connectorId)
        DeleteConnector(t, connectorId)
        DeleteTeam(t, teamId)
    })
}
