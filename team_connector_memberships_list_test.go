package fivetran_test

import (
    "context"
    "testing"
)

func TestTeamConnectorMembershipsListE2E(t *testing.T) {
    teamId := CreateTeam(t)
    connectorId := CreateConnector(t)
    CreateTeamConnector(t, teamId, connectorId)

    result, err := Client.NewTeamConnectorMembershipsList().TeamId(teamId).Do(context.Background())
    if err != nil {
        t.Logf("%+v\n", result)
        t.Error(err)
    }

    AssertEqual(t, result.Code, "Success")
    AssertEqual(t, result.Data.Items[0].ConnectorId, connectorId)
    AssertEqual(t, result.Data.Items[0].Role, "Connector Administrator")
    AssertNotEmpty(t, result.Data.Items[0].CreatedAt)

    t.Cleanup(func() { 
        DeleteTeamConnector(t, teamId, connectorId)
        DeleteConnector(t, connectorId)
        DeleteTeam(t, teamId)
    })
}
