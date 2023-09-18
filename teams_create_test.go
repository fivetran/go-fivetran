package fivetran_test

import (
    "context"
    "testing"
)

func TestNewTeamsCreateE2E(t *testing.T) {
    created, err := Client.NewTeamsCreate().
        Name("test_team").
        Description("test_description").
        Role("Account Reviewer").
        Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", created)
        t.Error(err)
    }

    AssertEqual(t, created.Code, "Success")
    AssertNotEmpty(t, created.Message)
    AssertNotEmpty(t, created.Data.Id)
    AssertEqual(t, created.Data.Name, "test_team")
    AssertEqual(t, created.Data.Description, "test_description")
    AssertEqual(t, created.Data.Role, "Account Reviewer")
    
    t.Cleanup(func() { DeleteTeam(t, created.Data.Id) })
}
