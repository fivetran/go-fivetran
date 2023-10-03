package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewUsersListE2E(t *testing.T) {
	users, err := testutils.Client.NewUsersList().Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", users)
		t.Error(err)
	}

	testutils.AssertEqual(t, users.Code, "Success")
	testutils.AssertHasLength(t, users.Data.Items, 1)
	testutils.AssertEqual(t, users.Data.Items[0].ID, testutils.PredefinedUserId)
	testutils.AssertEqual(t, users.Data.Items[0].Email, testutils.PredefinedUserEmail)
	testutils.AssertEqual(t, users.Data.Items[0].GivenName, testutils.PredefinedUserGivenName)
	testutils.AssertEqual(t, users.Data.Items[0].FamilyName, testutils.PredefinedUserFamilyName)
	testutils.AssertEqual(t, *users.Data.Items[0].Verified, true)
	testutils.AssertEqual(t, *users.Data.Items[0].Invited, false)
	testutils.AssertEmpty(t, users.Data.Items[0].Picture)
	testutils.AssertEqual(t, users.Data.Items[0].Phone, testutils.PredefinedUserPhone)
	testutils.AssertEqual(t, users.Data.Items[0].Role, "Account Administrator")
	testutils.AssertEqual(t, users.Data.Items[0].LoggedInAt.IsZero(), false)
	testutils.AssertEqual(t, users.Data.Items[0].CreatedAt.IsZero(), false)
	testutils.AssertEmpty(t, users.Data.NextCursor)
}
