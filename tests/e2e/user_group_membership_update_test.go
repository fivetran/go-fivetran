package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewUserGroupMembershipUpdateE2E(t *testing.T) {
	userId := testutils.CreateUser(t)
	groupId := testutils.CreateGroup(t)
	testutils.CreateUserGroup(t, userId, groupId)

	modified, err := testutils.Client.NewUserGroupMembershipUpdate().
		UserId(userId).
		GroupId(groupId).
		Role("Destination Reviewer").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", modified)
		t.Error(err)
	}

	testutils.AssertEqual(t, modified.Code, "Success")
	testutils.AssertEqual(t, modified.Message, "Group membership has been updated")

	t.Cleanup(func() {
		testutils.DeleteUserGroup(t, userId, groupId)
		testutils.DeleteGroup(t, groupId)
		testutils.DeleteUser(t, userId)
	})
}
