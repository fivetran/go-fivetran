package tests

import (
	"context"
	"testing"
	"github.com/corbym/gocrest/then"
	"github.com/corbym/gocrest/is"
)


func TestModifyUser(t *testing.T) {
	for _, c := range GetClients() {
		user, err :=  c.NewUserModify().
		UserID("_moon_acc_reader").
		FamilyName("Gerrard").
		GivenName("Steven").
		Phone("+19876543210").
		Picture("http://picture.com").
		Role("Owner").
		Do(context.Background())

		if err != nil {
			t.Logf("%+v\n", user)
			t.Error(err)
		}

		then.AssertThat(t, user.Code, is.EqualTo("Success"))
		then.AssertThat(t, user.Message, is.EqualTo("User has been updated")) 
		then.AssertThat(t, user.Data.ID, is.EqualTo("_moon_acc_reader")) 
		then.AssertThat(t, user.Data.Email, is.EqualTo("moon.acc.reader@gmail.com")) 
		then.AssertThat(t, user.Data.GivenName, is.EqualTo("Steven")) 
		then.AssertThat(t, user.Data.FamilyName, is.EqualTo("Gerrard")) 
		then.AssertThat(t, *user.Data.Verified, is.False())
		then.AssertThat(t, *user.Data.Invited, is.False())
		then.AssertThat(t, user.Data.Phone, is.EqualTo("+19876543210")) 
		then.AssertThat(t, user.Data.Picture, is.EqualTo("http://picture.com")) 
		//todo: incomment when role field will be mapped then.AssertThat(t, user.Data.Role, is.EqualTo("Owner"))
		then.AssertThat(t, user.Data.LoggedInAt.IsZero(), is.True())
		then.AssertThat(t, user.Data.CreatedAt.IsZero(), is.False())
	}
}