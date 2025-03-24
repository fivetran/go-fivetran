package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewUserConnectionMembershipCreateE2E(t *testing.T) {
	userId := testutils.CreateUser(t)
	ConnectionId := testutils.CreateConnection(t)

	created, err := testutils.Client.NewUserConnectionMembershipCreate().
		UserId(userId).
		ConnectionId(ConnectionId).
		Role("Connection Administrator").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertNotEmpty(t, created.Message)
	testutils.AssertEqual(t, created.Data.ConnectionId, ConnectionId)
	testutils.AssertEqual(t, created.Data.Role, "Connection Administrator")
	testutils.AssertNotEmpty(t, created.Data.CreatedAt)

	t.Cleanup(func() {
		testutils.DeleteUserConnection(t, userId, ConnectionId)
		testutils.DeleteConnection(t, ConnectionId)
		testutils.DeleteUser(t, userId)
	})
}
