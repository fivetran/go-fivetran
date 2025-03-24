package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestUserConnectionMembershipsListE2E(t *testing.T) {
	userId := testutils.CreateUser(t)
	ConnectionId := testutils.CreateConnection(t)
	testutils.CreateUserConnection(t, userId, ConnectionId)

	result, err := testutils.Client.NewUserConnectionMembershipsList().UserId(userId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertEqual(t, result.Data.Items[0].ConnectionId, ConnectionId)
	testutils.AssertEqual(t, result.Data.Items[0].Role, "Connection Administrator")
	testutils.AssertNotEmpty(t, result.Data.Items[0].CreatedAt)

	t.Cleanup(func() {
		testutils.DeleteUserConnection(t, userId, ConnectionId)
		testutils.DeleteConnection(t, ConnectionId)
		testutils.DeleteUser(t, userId)
	})
}
