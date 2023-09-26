package fivetran_test

import (
    "context"
    "testing"
)

func TestNewTeamRoleDeleteE2E(t *testing.T) {
    teamId := CreateTeam(t)

    deleted, err := Client.NewTeamsDeleteRoleInAccount().TeamId(teamId).Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", deleted)
        t.Error(err)
    }

    AssertEqual(t, deleted.Code, "Success")
    AssertEqual(t, deleted.Message, "Team role in account has been removed")

    t.Cleanup(func() { DeleteTeam(t, teamId) })
}
