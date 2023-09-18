package fivetran_test

import (
    "context"
    "testing"
)

func TestNewTeamsModifyE2E(t *testing.T) {
    teamId := CreateTeam(t)

    modified, err := Client.NewTeamsModify().
        TeamId(teamId).
        Name("test_team").
        Description("test_description_2").
        Role("Account Reviewer").
        Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", modified)
        t.Error(err)
    }

    AssertEqual(t, modified.Code, "Success")
    AssertNotEmpty(t, modified.Message)
    AssertNotEmpty(t, modified.Data.Id)
    AssertEqual(t, modified.Data.Name, "test_team")
    AssertEqual(t, modified.Data.Description, "test_description_2")
    AssertEqual(t, modified.Data.Role, "Account Reviewer")
    
    t.Cleanup(func() { DeleteTeam(t, modified.Data.Id) })
}
