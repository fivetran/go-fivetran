package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewUserUpdateE2E(t *testing.T) {
	userId := testutils.CreateTempUser(t)
	user, err := testutils.Client.NewUserUpdate().
		UserID(userId).
		FamilyName("Gerrard").
		GivenName("Steven").
		Phone("+19876543210").
		Picture("http://picture.com").
		Role("Destination Creator").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", user)
		t.Error(err)
	}

	testutils.AssertEqual(t, user.Code, "Success")
	testutils.AssertNotEmpty(t, user.Message)
	testutils.AssertEqual(t, user.Data.ID, userId)
	testutils.AssertEqual(t, user.Data.Email, "william_addison.@fivetran.com")
	testutils.AssertEqual(t, user.Data.GivenName, "Steven")
	testutils.AssertEqual(t, user.Data.FamilyName, "Gerrard")
	testutils.AssertEqual(t, *user.Data.Verified, false)
	testutils.AssertEqual(t, *user.Data.Invited, true)
	testutils.AssertEqual(t, user.Data.Phone, "+19876543210")
	testutils.AssertEqual(t, user.Data.Picture, "http://picture.com")
	//todo: incomment when role field will be mapped then.testutils.AssertThat(t, user.Data.Role, is.EqualTo("Owner"))
	testutils.AssertEqual(t, user.Data.LoggedInAt.IsZero(), true)
	testutils.AssertEqual(t, user.Data.CreatedAt.IsZero(), false)
}
