package fivetran_test

import (
    "context"
    "testing"
)

func TestTeamGroupMembershipDetailsE2E(t *testing.T) {
    groupId := CreateGroup(t)
    teamId := CreateTeam(t)
    CreateTeamGroup(t, teamId, groupId)

    result, err := Client.NewTeamGroupMembershipDetails().
        TeamId(teamId).
        GroupId(groupId).
        Do(context.Background())
    if err != nil {
        t.Logf("%+v\n", result)
        t.Error(err)
    }

    AssertEqual(t, result.Code, "Success")
    AssertEqual(t, result.Data.GroupId, groupId)
    AssertEqual(t, result.Data.Role, "Destination Analyst")
    AssertNotEmpty(t, result.Data.CreatedAt)

    t.Cleanup(func() { 
        DeleteTeamGroup(t, teamId, groupId)
        DeleteGroup(t, groupId)
        DeleteTeam(t, teamId)
    })
}
