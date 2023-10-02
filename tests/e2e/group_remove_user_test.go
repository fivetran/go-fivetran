package fivetran_test

import (
	"context"
	"strings"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewGroupRemoveUserE2E(t *testing.T) {
	userId := testutils.CreateUser(t)
	testutils.AddUserToGroup(t, testutils.PredefinedGroupId, "william_addison.@fivetran.com")

	deleted, err := testutils.Client.NewGroupRemoveUser().GroupID(testutils.PredefinedGroupId).UserID(userId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}

	testutils.AssertEqual(t, deleted.Code, "Success")
	testutils.AssertNotEmpty(t, deleted.Message)
	testutils.AssertEqual(t, strings.Contains(deleted.Message, userId), true)

	t.Cleanup(func() {
		testutils.DeleteUser(t, userId)
	})
}
