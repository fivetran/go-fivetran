package fivetran_test

import (
    "context"
    "testing"
)

func TestTeamUserMembershipsListE2E(t *testing.T) {
    teamId := CreateTeam(t)
    CreateTeamUser(t, teamId, PredefinedUserId)

    result, err := Client.NewTeamUserMembershipsList().
        TeamId(teamId).
        Do(context.Background())
    if err != nil {
        t.Logf("%+v\n", result)
        t.Error(err)
    }

    AssertEqual(t, result.Code, "Success")
    AssertEqual(t, result.Data.Items[0].UserId, PredefinedUserId)
    AssertEqual(t, result.Data.Items[0].Role, "Team Member")

    t.Cleanup(func() { 
        DeleteTeamUser(t, teamId, PredefinedUserId)
        DeleteTeam(t, teamId)
    })
}
