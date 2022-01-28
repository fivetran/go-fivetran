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
	AssertEqual(t, created.Message, "User has been invited to the group")

	t.Cleanup(func() { RemoveUserFromGroup(t, PredefinedGroupId, userId) })
}
