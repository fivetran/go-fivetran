package fivetran_test

import (
	"context"
	"testing"
)

func TestNewGroupAddUserE2E(t *testing.T) {
	for version, c := range Clients {
		t.Run(version, func(t *testing.T) {
			t.Skip("Account has new RBAC model in place and we can't add a user with a new role names. It will be fix soon")

			userId := CreateUser(t)

			created, err := c.NewGroupAddUser().GroupID(PredefinedGroupId).
				Email("william_addison.@fivetran.com").
				Role("Destination Administrator").
				Do(context.Background())

			if err != nil {
				t.Logf("%+v\n", created)
				t.Error(err)
			}

			AssertEqual(t, created.Code, "Success")
			AssertEqual(t, created.Message, "User has been invited to the group")

			t.Cleanup(func() { RemoveUserFromGroup(t, PredefinedGroupId, userId) })
		})
	}
}
