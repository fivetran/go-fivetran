package fivetran_test

import (
	"context"
	"testing"
)

func TestNewGroupRemoveUserE2E(t *testing.T) {
	t.Skip("Account has new RBAC model in place and we can't add a user with a new role names. It will be fixed soon")

	userId := CreateUser(t)
	AddUserToGroup(t, PredefinedGroupId, "william_addison.@fivetran.com")

	deleted, err := Client.NewGroupRemoveUser().GroupID(PredefinedGroupId).UserID(userId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}

	AssertEqual(t, deleted.Code, "Success")
	AssertEqual(t, deleted.Message, "User with id '"+userId+"' has been removed from the group")
}
