package fivetran_test

import (
    "context"
    "testing"
)

func TestTeamUserMembershipDetailsE2E(t *testing.T) {
    teamId := CreateTeam(t)
    CreateTeamUser(t, teamId, PredefinedUserId)

    result, err := Client.NewTeamUserMembershipDetails().
        TeamId(teamId).
        UserId(PredefinedUserId).
        Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", result)
        t.Error(err)
    }

    AssertEqual(t, result.Code, "Success")
    AssertEqual(t, result.Data.UserId, PredefinedUserId)
    AssertEqual(t, result.Data.Role, "Team Member")

    t.Cleanup(func() { 
        DeleteTeamUser(t, teamId, PredefinedUserId)
        DeleteTeam(t, teamId)
    })
}