package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewUserConnectionMebershipDeleteE2E(t *testing.T) {
	userId := testutils.CreateUser(t)
	ConnectionId := testutils.CreateConnection(t)
	testutils.CreateUserConnection(t, userId, ConnectionId)

	deleted, err := testutils.Client.NewUserConnectionMembershipDelete().
		UserId(userId).
		ConnectionId(ConnectionId).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}

	testutils.AssertEqual(t, deleted.Code, "Success")
	testutils.AssertNotEmpty(t, deleted.Message)

	t.Cleanup(func() {
		testutils.DeleteConnection(t, ConnectionId)
		testutils.DeleteUser(t, userId)
	})
}
