package fivetran_test

import (
    "context"
    "testing"
)

func TestNewTeamsDeleteE2E(t *testing.T) {
    teamId := CreateTeam(t)

    deleted, err := Client.NewTeamsDelete().TeamId(teamId).Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", deleted)
        t.Error(err)
    }

    AssertEqual(t, deleted.Code, "Success")
    AssertEqual(t, deleted.Message, "Team has been deleted")
}
