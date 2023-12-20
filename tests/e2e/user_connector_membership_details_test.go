package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestUserConnectorMembershipDetailsE2E(t *testing.T) {
	userId := testutils.CreateUser(t)
	connectorId := testutils.CreateConnector(t)
	testutils.CreateUserConnector(t, userId, connectorId)

	result, err := testutils.Client.NewUserConnectorMembershipDetails().
		UserId(userId).
		ConnectorId(connectorId).
		Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertEqual(t, result.Data.ConnectorId, connectorId)
	testutils.AssertEqual(t, result.Data.Role, "Connector Administrator")
	testutils.AssertNotEmpty(t, result.Data.CreatedAt)

	t.Cleanup(func() {
		testutils.DeleteUserConnector(t, userId, connectorId)
		testutils.DeleteConnector(t, connectorId)
		testutils.DeleteUser(t, userId)
	})
}
