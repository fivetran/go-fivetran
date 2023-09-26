package fivetran_test

import (
    "context"
    "testing"
)

func TestNewTeamGroupMembershipDeleteE2E(t *testing.T) {
    groupId := CreateGroup(t)
    teamId := CreateTeam(t)
    CreateTeamGroup(t, teamId,groupId)

    deleted, err := Client.NewTeamGroupMembershipDelete().
        TeamId(teamId).
        GroupId(groupId).
        Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", deleted)
        t.Error(err)
    }

    AssertEqual(t, deleted.Code, "Success")
    AssertEqual(t, deleted.Message, "Group membership has been deleted")

    t.Cleanup(func() { 
        DeleteGroup(t, groupId)
        DeleteTeam(t, teamId)
    })
}
