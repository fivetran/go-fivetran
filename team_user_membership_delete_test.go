package fivetran_test

import (
    "context"
    "testing"
)

func TestNewTeamUserMembershipDeleteE2E(t *testing.T) {
    teamId := CreateTeam(t)
    CreateTeamUser(t, teamId, PredefinedUserId)

    deleted, err := Client.NewTeamUserMembershipDelete().
        TeamId(teamId).
        UserId(PredefinedUserId).
        Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", deleted)
        t.Error(err)
    }

    AssertEqual(t, deleted.Code, "Success")
    AssertEqual(t, deleted.Message, "User has been removed from the team")

    t.Cleanup(func() { 
        DeleteTeam(t, teamId)
    })
}