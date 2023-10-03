package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewGroupAddUserE2E(t *testing.T) {
	userId := testutils.CreateUser(t)

	created, err := testutils.Client.NewGroupAddUser().GroupID(testutils.PredefinedGroupId).
		Email("william_addison.@fivetran.com").
		Role("Destination Administrator").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertNotEmpty(t, created.Message)

	t.Cleanup(func() {
		testutils.RemoveUserFromGroup(t, testutils.PredefinedGroupId, userId)
		testutils.DeleteUser(t, userId)
	})
}
