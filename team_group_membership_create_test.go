package fivetran_test

import (
    "context"
    "testing"
)

func TestNewTeamGroupMembershipCreateE2E(t *testing.T) {
    groupId := CreateGroup(t)
    teamId := CreateTeam(t)

    created, err := Client.NewTeamGroupMembershipCreate().
        TeamId(teamId).
        GroupId(groupId).
        Role("Destination Administrator").
        Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", created)
        t.Error(err)
    }

    AssertEqual(t, created.Code, "Success")
    AssertEqual(t, created.Message, "Group membership has been created")
    AssertEqual(t, created.Data.GroupId, groupId)
    AssertEqual(t, created.Data.Role, "Destination Administrator")
    AssertNotEmpty(t, created.Data.CreatedAt)
    
    t.Cleanup(func() { 
        DeleteTeamGroup(t, teamId, groupId)
        DeleteGroup(t, groupId)
        DeleteTeam(t, teamId)
    })
}
