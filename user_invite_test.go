package fivetran_test

import (
	"context"
	"testing"
)

func TestNewUserInviteE2E(t *testing.T) {
	user, err := Client.NewUserInvite().
		Email("william_addison.@fivetran.com").
		GivenName("William").
		FamilyName("Addison").
		Phone("+19876543210").
		Role("Account Reviewer").
		Picture("http://picture.com").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", user)
		t.Error(err)
	}

	AssertEqual(t, user.Code, "Success")
	AssertEqual(t, user.Message, "User has been invited to the account")
	AssertNotEmpty(t, user.Data.ID)
	AssertEqual(t, user.Data.Email, "william_addison.@fivetran.com")
	AssertEqual(t, user.Data.GivenName, "William")
	AssertEqual(t, user.Data.FamilyName, "Addison")
	AssertEqual(t, *user.Data.Verified, false)
	AssertEqual(t, *user.Data.Invited, true)
	AssertEqual(t, user.Data.Picture, "http://picture.com")
	AssertEqual(t, user.Data.Phone, "+19876543210")
	AssertEqual(t, user.Data.LoggedInAt.IsZero(), true)
	AssertEqual(t, user.Data.CreatedAt.IsZero(), false)
	t.Cleanup(func() { DeleteUser(t, user.Data.ID) })
}
