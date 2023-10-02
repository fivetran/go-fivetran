package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewUserInviteE2E(t *testing.T) {
	user, err := testutils.Client.NewUserInvite().
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

	testutils.AssertEqual(t, user.Code, "Success")
	testutils.AssertNotEmpty(t, user.Message)
	testutils.AssertNotEmpty(t, user.Data.ID)
	testutils.AssertEqual(t, user.Data.Email, "william_addison.@fivetran.com")
	testutils.AssertEqual(t, user.Data.GivenName, "William")
	testutils.AssertEqual(t, user.Data.FamilyName, "Addison")
	testutils.AssertEqual(t, *user.Data.Verified, false)
	testutils.AssertEqual(t, *user.Data.Invited, true)
	testutils.AssertEqual(t, user.Data.Picture, "http://picture.com")
	testutils.AssertEqual(t, user.Data.Phone, "+19876543210")
	testutils.AssertEqual(t, user.Data.LoggedInAt.IsZero(), true)
	testutils.AssertEqual(t, user.Data.CreatedAt.IsZero(), false)
	t.Cleanup(func() { testutils.DeleteUser(t, user.Data.ID) })
}
