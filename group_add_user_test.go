package fivetran_test

import (
	"context"
	"testing"
)

func TestNewGroupAddUserE2E(t *testing.T) {
	userId := CreateUser(t)

	created, err := Client.NewGroupAddUser().GroupID(PredefinedGroupId).
		Email("william_addison.@fivetran.com").
		Role("Destination Administrator").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	AssertEqual(t, created.Code, "Success")
	AssertNotEmpty(t, created.Message)

	t.Cleanup(func() {
		RemoveUserFromGroup(t, PredefinedGroupId, userId)
		DeleteUser(t, userId)
	})
}
