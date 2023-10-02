package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewUserDetailsE2E(t *testing.T) {
	user, err := testutils.Client.NewUserDetails().UserID(testutils.PredefinedUserId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", user)
		t.Error(err)
	}

	testutils.AssertEqual(t, user.Code, "Success")
	testutils.AssertEqual(t, user.Data.ID, testutils.PredefinedUserId)
	testutils.AssertEqual(t, user.Data.Email, testutils.PredefinedUserEmail)
	testutils.AssertEqual(t, user.Data.GivenName, testutils.PredefinedUserGivenName)
	testutils.AssertEqual(t, user.Data.FamilyName, testutils.PredefinedUserFamilyName)
	testutils.AssertEqual(t, *user.Data.Verified, true)
	testutils.AssertEqual(t, *user.Data.Invited, false)
	testutils.AssertEmpty(t, user.Data.Picture)
	testutils.AssertEqual(t, user.Data.Phone, testutils.PredefinedUserPhone)
	testutils.AssertEqual(t, user.Data.Role, "Account Administrator")
	testutils.AssertEqual(t, user.Data.LoggedInAt.IsZero(), false)
	testutils.AssertEqual(t, user.Data.CreatedAt.IsZero(), false)
}
