package fivetran_test

import (
    "context"
    "testing"
)

func TestNewTeamConnectorMebershipDeleteE2E(t *testing.T) {
    teamId := CreateTeam(t)
    connectorId := CreateConnector(t)
    CreateTeamConnector(t, teamId, connectorId)

    deleted, err := Client.NewTeamConnectorMembershipDelete().
        TeamId(teamId).
        ConnectorId(connectorId).
        Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", deleted)
        t.Error(err)
    }

    AssertEqual(t, deleted.Code, "Success")
    AssertEqual(t, deleted.Message, "Connector membership has been deleted")

    t.Cleanup(func() { 
        DeleteConnector(t, connectorId)
        DeleteTeam(t, teamId)
    })
}
