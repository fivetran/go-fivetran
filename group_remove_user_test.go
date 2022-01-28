package fivetran_test

import (
	"context"
	"testing"
)

func TestNewGroupRemoveUserE2E(t *testing.T) {
	userId := CreateUser(t)
	AddUserToGroup(t, PredefinedGroupId, "william_addison.@fivetran.com")

	deleted, err := Client.NewGroupRemoveUser().GroupID(PredefinedGroupId).UserID(userId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}

	AssertEqual(t, deleted.Code, "Success")
	AssertEqual(t, deleted.Message, "User with id '"+userId+"' has been removed from the group")

	t.Cleanup(func() {
		RemoveUserFromGroup(t, PredefinedGroupId, userId)
		DeleteUser(t, userId)
	})
}
