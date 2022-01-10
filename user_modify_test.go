package fivetran_test

import (
	"context"
	"testing"
)

func TestNewUserModifyE2E(t *testing.T) {
	userId := CreateTempUser(t)
	user, err := Client.NewUserModify().
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

	AssertEqual(t, user.Code, "Success")
	AssertEqual(t, user.Message, "User has been updated")
	AssertEqual(t, user.Data.ID, userId)
	AssertEqual(t, user.Data.Email, "william_addison.@fivetran.com")
	AssertEqual(t, user.Data.GivenName, "Steven")
	AssertEqual(t, user.Data.FamilyName, "Gerrard")
	AssertEqual(t, *user.Data.Verified, false)
	AssertEqual(t, *user.Data.Invited, true)
	AssertEqual(t, user.Data.Phone, "+19876543210")
	AssertEqual(t, user.Data.Picture, "http://picture.com")
	//todo: incomment when role field will be mapped then.AssertThat(t, user.Data.Role, is.EqualTo("Owner"))
	AssertEqual(t, user.Data.LoggedInAt.IsZero(), true)
	AssertEqual(t, user.Data.CreatedAt.IsZero(), false)
}
