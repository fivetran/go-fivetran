package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewUserConnectionMembershipUpdateE2E(t *testing.T) {
	userId := testutils.CreateUser(t)
	ConnectionId := testutils.CreateConnection(t)
	testutils.CreateUserConnection(t, userId, ConnectionId)

	modified, err := testutils.Client.NewUserConnectionMembershipUpdate().
		UserId(userId).
		ConnectionId(ConnectionId).
		Role("Connection Collaborator").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", modified)
		t.Error(err)
	}

	testutils.AssertEqual(t, modified.Code, "Success")
	testutils.AssertNotEmpty(t, modified.Message)

	t.Cleanup(func() {
		testutils.DeleteUserConnection(t, userId, ConnectionId)
		testutils.DeleteConnection(t, ConnectionId)
		testutils.DeleteUser(t, userId)
	})
}
