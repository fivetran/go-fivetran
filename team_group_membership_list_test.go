package fivetran_test

import (
    "context"
    "testing"
)

func TestTeamGroupMembershipsistE2E(t *testing.T) {
    groupId := CreateGroup(t)
    teamId := CreateTeam(t)
    CreateTeamGroup(t, teamId, groupId)

    result, err := Client.NewTeamGroupMembershipsList().TeamId(teamId).Do(context.Background())
    if err != nil {
        t.Logf("%+v\n", result)
        t.Error(err)
    }

    AssertEqual(t, result.Code, "Success")
    AssertEqual(t, result.Data.Items[0].GroupId, groupId)
    AssertEqual(t, result.Data.Items[0].Role, "Destination Analyst")
    AssertNotEmpty(t, result.Data.Items[0].CreatedAt)

    t.Cleanup(func() { 
        DeleteTeamGroup(t, teamId, groupId)
        DeleteGroup(t, groupId)
        DeleteTeam(t, teamId)
    })
}
