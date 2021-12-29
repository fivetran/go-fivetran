package fivetran_test

import (
	"context"
	"testing"
)

func TestNewGroupRemoveUserIntegration(t *testing.T) {
	t.Skip("Account has new RBAC model in place and we can't add a user with a new role names. It will be fix soon")

	for version, c := range Clients {
		t.Run(version, func(t *testing.T) {
			userId := CreateUser(t)
			AddUserToGroup(t, "climbed_consulted", "william_addison.@fivetran.com")

			deleted, err := c.NewGroupRemoveUser().GroupID("climbed_consulted").UserID(userId).Do(context.Background())
			if err != nil {
				t.Logf("%+v\n", deleted)
				t.Error(err)
			}

			AssertEqual(t, deleted.Code, "Success")
			AssertEqual(t, deleted.Message, "User with id '"+userId+"' has been removed from the group")
		})
	}
}
