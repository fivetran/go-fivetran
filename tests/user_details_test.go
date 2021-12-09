package tests

import (
	"context"
	"testing"
	"github.com/corbym/gocrest/then"
	"github.com/corbym/gocrest/is"
)


func TestGetUserDetails(t *testing.T) {
	for _, c := range GetClients() {
		user, err :=  c.NewUserDetails().UserID("_accountworthy").Do(context.Background())
		if err != nil {
			t.Logf("%+v\n", user)
			t.Error(err)
		}

		then.AssertThat(t, user.Code, is.EqualTo("Success"))
		then.AssertThat(t, user.Data.ID, is.EqualTo("_accountworthy"))
		then.AssertThat(t, user.Data.Email, is.EqualTo("accountworthy.moonbeam@gmail.com"))
		then.AssertThat(t, user.Data.GivenName, is.EqualTo("Accountworthy"))
		then.AssertThat(t, user.Data.FamilyName, is.EqualTo("Moonbeam"))
		then.AssertThat(t, *user.Data.Verified, is.True())
		then.AssertThat(t, *user.Data.Invited, is.False())
		then.AssertThat(t, user.Data.Picture, is.Not(is.Empty()))
		then.AssertThat(t, user.Data.Phone, is.Empty())
		then.AssertThat(t, user.Data.LoggedInAt.IsZero(), is.True())
		then.AssertThat(t, user.Data.CreatedAt.IsZero(), is.False())
	}
}