package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewUserGroupMembershipCreateE2E(t *testing.T) {
	groupId := testutils.CreateGroup(t)
	userId := testutils.CreateUser(t)

	created, err := testutils.Client.NewUserGroupMembershipCreate().
		UserId(userId).
		GroupId(groupId).
		Role("Destination Analyst").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertEqual(t, created.Message, "Group membership has been created")
	testutils.AssertEqual(t, created.Data.GroupId, groupId)
	testutils.AssertEqual(t, created.Data.Role, "Destination Analyst")
	testutils.AssertNotEmpty(t, created.Data.CreatedAt)

	t.Cleanup(func() {
		testutils.DeleteUserGroup(t, userId, groupId)
		testutils.DeleteGroup(t, groupId)
		testutils.DeleteUser(t, userId)
	})
}
