package fivetran_test

import (
    "context"
    "testing"
)

func TestNewTeamUserMembershipCreateE2E(t *testing.T) {
    teamId := CreateTeam(t)

    created, err := Client.NewTeamUserMembershipCreate().
        TeamId(teamId).
        UserId(PredefinedUserId).
        Role("Team Member").
        Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", created)
        t.Error(err)
    }

    AssertEqual(t, created.Code, "Success")
    AssertEqual(t, created.Message, "User has been added to the team")
    AssertEqual(t, created.Data.UserId, PredefinedUserId)
    AssertEqual(t, created.Data.Role, "Team Member")
    
    t.Cleanup(func() { 
        DeleteTeamUser(t, teamId, PredefinedUserId)
        DeleteTeam(t, teamId)
    })
}
