package fivetran_test

import (
    "context"
    "testing"
)

func TestTeamsListE2E(t *testing.T) {
    teamId := CreateTeam(t)

    result, err := Client.NewTeamsList().Do(context.Background())
    if err != nil {
        t.Logf("%+v\n", result)
        t.Error(err)
    }

    AssertEqual(t, result.Code, "Success")
    AssertEqual(t, result.Data.Items[0].Id, teamId)
    AssertEqual(t, result.Data.Items[0].Name, "test_team")
    AssertEqual(t, result.Data.Items[0].Description, "test_description")
    AssertEqual(t, result.Data.Items[0].Role, "Account Reviewer")

    t.Cleanup(func() { DeleteTeam(t, teamId) })
}
