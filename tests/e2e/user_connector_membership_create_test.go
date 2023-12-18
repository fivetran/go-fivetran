package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewUserConnectorMembershipCreateE2E(t *testing.T) {
	userId := testutils.CreateUser(t)
	connectorId := testutils.CreateConnector(t)

	created, err := testutils.Client.NewUserConnectorMembershipCreate().
		UserId(userId).
		ConnectorId(connectorId).
		Role("Connector Administrator").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertEqual(t, created.Message, "Connector membership has been created")
	testutils.AssertEqual(t, created.Data.ConnectorId, connectorId)
	testutils.AssertEqual(t, created.Data.Role, "Connector Administrator")
	testutils.AssertNotEmpty(t, created.Data.CreatedAt)

	t.Cleanup(func() {
		testutils.DeleteUserConnector(t, userId, connectorId)
		testutils.DeleteConnector(t, connectorId)
		testutils.DeleteUser(t, userId)
	})
}
