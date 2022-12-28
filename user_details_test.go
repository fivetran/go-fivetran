package fivetran_test

import (
	"context"
	"testing"
)

func TestNewUserDetailsE2E(t *testing.T) {
	user, err := Client.NewUserDetails().UserID(PredefinedUserId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", user)
		t.Error(err)
	}

	AssertEqual(t, user.Code, "Success")
	AssertEqual(t, user.Data.ID, PredefinedUserId)
	AssertEqual(t, user.Data.Email, PredefinedUserEmail)
	AssertEqual(t, user.Data.GivenName, PredefinedUserGivenName)
	AssertEqual(t, user.Data.FamilyName, PredefinedUserFamilyName)
	AssertEqual(t, *user.Data.Verified, true)
	AssertEqual(t, *user.Data.Invited, false)
	AssertEmpty(t, user.Data.Picture)
	AssertEqual(t, user.Data.Phone, PredefinedUserPhone)
	AssertEqual(t, user.Data.Role, "Account Administrator")
	AssertEqual(t, user.Data.LoggedInAt.IsZero(), false)
	AssertEqual(t, user.Data.CreatedAt.IsZero(), false)
}
