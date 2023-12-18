package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestUserGroupMembershipsistE2E(t *testing.T) {
	groupId := testutils.CreateGroup(t)
	userId := testutils.CreateUser(t)
	testutils.CreateUserGroup(t, userId, groupId)

	result, err := testutils.Client.NewUserGroupMembershipsList().UserId(userId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertEqual(t, result.Data.Items[0].GroupId, groupId)
	testutils.AssertEqual(t, result.Data.Items[0].Role, "Destination Analyst")
	testutils.AssertNotEmpty(t, result.Data.Items[0].CreatedAt)

	t.Cleanup(func() {
		testutils.DeleteUserGroup(t, userId, groupId)
		testutils.DeleteGroup(t, groupId)
		testutils.DeleteUser(t, userId)
	})
}
