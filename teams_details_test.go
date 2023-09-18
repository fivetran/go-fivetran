package fivetran_test

import (
	"context"
	"testing"
)

func TestTeamsDetailsE2E(t *testing.T) {
	teamId := CreateTeam(t)

	result, err := Client.NewTeamsDetails().TeamId(teamId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	AssertEqual(t, result.Code, "Success")
    AssertEqual(t, result.Data.Id, teamId)
    AssertEqual(t, result.Data.Name, "test_team")
    AssertEqual(t, result.Data.Description, "test_description")
    AssertEqual(t, result.Data.Role, "Account Reviewer")

    t.Cleanup(func() { DeleteTeam(t, teamId) })
}