package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewUserGroupMembershipDeleteE2E(t *testing.T) {
	groupId := testutils.CreateGroup(t)
	userId := testutils.CreateUser(t)
	testutils.CreateUserGroup(t, userId, groupId)

	deleted, err := testutils.Client.NewUserGroupMembershipDelete().
		UserId(userId).
		GroupId(groupId).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}

	testutils.AssertEqual(t, deleted.Code, "Success")
	testutils.AssertEqual(t, deleted.Message, "Group membership has been deleted")

	t.Cleanup(func() {
		testutils.DeleteGroup(t, groupId)
		testutils.DeleteUser(t, userId)
	})
}
