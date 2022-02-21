package fivetran_test

import (
	"context"
	"strings"
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
	AssertNotEmpty(t, deleted.Message)
	AssertEqual(t, strings.Contains(deleted.Message, userId), true)

	t.Cleanup(func() {
		DeleteUser(t, userId)
	})
}
