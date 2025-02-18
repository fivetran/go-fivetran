package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewUserConnectorMebershipDeleteE2E(t *testing.T) {
	userId := testutils.CreateUser(t)
	connectorId := testutils.CreateConnector(t)
	testutils.CreateUserConnector(t, userId, connectorId)

	deleted, err := testutils.Client.NewUserConnectorMembershipDelete().
		UserId(userId).
		ConnectorId(connectorId).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}

	testutils.AssertEqual(t, deleted.Code, "Success")
	testutils.AssertNotEmpty(t, deleted.Message)

	t.Cleanup(func() {
		testutils.DeleteConnector(t, connectorId)
		testutils.DeleteUser(t, userId)
	})
}
