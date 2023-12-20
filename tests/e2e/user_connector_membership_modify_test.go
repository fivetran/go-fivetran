package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewUserConnectorMembershipModifyE2E(t *testing.T) {
	userId := testutils.CreateUser(t)
	connectorId := testutils.CreateConnector(t)
	testutils.CreateUserConnector(t, userId, connectorId)

	modified, err := testutils.Client.NewUserConnectorMembershipModify().
		UserId(userId).
		ConnectorId(connectorId).
		Role("Connector Collaborator").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", modified)
		t.Error(err)
	}

	testutils.AssertEqual(t, modified.Code, "Success")
	testutils.AssertEqual(t, modified.Message, "Connector membership has been updated")

	t.Cleanup(func() {
		testutils.DeleteUserConnector(t, userId, connectorId)
		testutils.DeleteConnector(t, connectorId)
		testutils.DeleteUser(t, userId)
	})
}
