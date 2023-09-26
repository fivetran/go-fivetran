package fivetran_test

import (
    "context"
    "testing"
)

func TestNewTeamGroupMembershipModifyE2E(t *testing.T) {
    teamId := CreateTeam(t)
    groupId := CreateGroup(t)
    CreateTeamGroup(t, teamId, groupId)

    modified, err := Client.NewTeamGroupMembershipModify().
        TeamId(teamId).
        GroupId(groupId).
        Role("Destination Reviewer").
        Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", modified)
        t.Error(err)
    }

    AssertEqual(t, modified.Code, "Success")
    AssertEqual(t, modified.Message, "Group membership has been updated")
    
    t.Cleanup(func() { 
        DeleteTeamGroup(t, teamId, groupId)
        DeleteGroup(t, groupId)
        DeleteTeam(t, teamId)
    })
}
