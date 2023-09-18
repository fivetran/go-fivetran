package fivetran_test

import (
    "context"
    "testing"
)

func TestNewTeamUserMembershipModifyE2E(t *testing.T) {
    teamId := CreateTeam(t)
    CreateTeamUser(t, teamId, PredefinedUserId)

    modified, err := Client.NewTeamUserMembershipModify().
        TeamId(teamId).
        UserId(PredefinedUserId).
        Role("Team Manager").
        Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", modified)
        t.Error(err)
    }

    AssertEqual(t, modified.Code, "Success")
    AssertEqual(t, modified.Message, "User role has been updated")
    
    t.Cleanup(func() { 
        DeleteTeamUser(t, teamId, PredefinedUserId)
        DeleteTeam(t, teamId)
    })
}