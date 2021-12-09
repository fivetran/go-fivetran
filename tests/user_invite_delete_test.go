package tests

import (
	"context"
	"testing"
	"github.com/corbym/gocrest/then"
	"github.com/corbym/gocrest/is"
)

func TestUserInviteDelete(t *testing.T) {
	for _, c := range GetClients() {
		//create
		user, err :=  c.NewUserInvite().
		Email("john.black@testmail.com").
		GivenName("John").
		FamilyName("Black").
		Phone("+19876543210").
		Role("ReadOnly").
		Do(context.Background())

		if err != nil {
			t.Logf("%+v\n", user)
			t.Error(err)
		}

		then.AssertThat(t, user.Code, is.EqualTo("Success"))
		then.AssertThat(t, user.Message, is.EqualTo("User has been invited to the account")) 
		then.AssertThat(t, user.Data.ID, is.Not(is.Empty()))
		then.AssertThat(t, user.Data.Email, is.EqualTo("john.black@testmail.com"))
		then.AssertThat(t, user.Data.GivenName, is.EqualTo("John"))
		then.AssertThat(t, user.Data.FamilyName, is.EqualTo("Black"))
		then.AssertThat(t, *user.Data.Verified, is.False())
		then.AssertThat(t, *user.Data.Invited, is.True())
		then.AssertThat(t, user.Data.Picture, is.Empty())
		then.AssertThat(t, user.Data.Phone, is.EqualTo("+19876543210"))
		then.AssertThat(t, user.Data.LoggedInAt.IsZero(), is.True())
		then.AssertThat(t, user.Data.CreatedAt.IsZero(), is.False())

		//cleanup
		deleted, err :=  c.NewUserDelete().UserID(user.Data.ID).Do(context.Background())

		if err != nil {
			t.Logf("%+v\n", deleted)
			t.Error(err)
		}

		then.AssertThat(t, deleted.Code, is.EqualTo("Success"))
		then.AssertThat(t, deleted.Message, is.EqualTo("User with id '" + user.Data.ID + "' has been deleted"))
	}
}