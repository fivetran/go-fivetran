package fivetran_test

import (
	"context"
	"testing"
)

func TestNewGroupListUsersE2E(t *testing.T) {
	users, err := Client.NewGroupListUsers().GroupID(PredefinedGroupId).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", users)
		t.Error(err)
	}

	AssertEqual(t, users.Code, "Success")
	AssertHasLength(t, users.Data.Items, 1)
	AssertEqual(t, users.Data.Items[0].ID, PredefinedUserId)
	AssertEqual(t, users.Data.Items[0].Email, PredefinedUserEmail)
	AssertEqual(t, users.Data.Items[0].GivenName, PredefinedUserGivenName)
	AssertEqual(t, users.Data.Items[0].FamilyName, PredefinedUserFamilyName)
	AssertEqual(t, *users.Data.Items[0].Verified, true)
	AssertEqual(t, *users.Data.Items[0].Invited, false)
	AssertEmpty(t, users.Data.Items[0].Picture)
	AssertEqual(t, users.Data.Items[0].Phone, PredefinedUserPhone)
	AssertEqual(t, users.Data.Items[0].LoggedInAt.IsZero(), false)
	AssertEqual(t, users.Data.Items[0].CreatedAt.IsZero(), false)
	AssertEmpty(t, users.Data.NextCursor)
}
