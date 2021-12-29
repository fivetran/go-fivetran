package fivetran_test

import (
	"context"
	"testing"
)

func TestNewGroupAddUserIntegration(t *testing.T) {
	t.Skip("Account has new RBAC model in place and we can't add a user with a new role names. It will be fix soon")

	for version, c := range Clients {
		t.Run(version, func(t *testing.T) {
			userId := CreateUser(t)

			created, err := c.NewGroupAddUser().GroupID("climbed_consulted").
				Email("william_addison.@fivetran.com").
				Role("Destination Administrator").
				Do(context.Background())

			if err != nil {
				t.Logf("%+v\n", created)
				t.Error(err)
			}

			AssertEqual(t, created.Code, "Success")
			AssertEqual(t, created.Message, "User has been invited to the group")

			t.Cleanup(func() { RemoveUserFromGroup(t, "climbed_consulted", userId) })
		})
	}
}
