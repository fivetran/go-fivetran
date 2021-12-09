package tests

import (
	"context"
	"testing"
	"github.com/corbym/gocrest/then"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/has"
)
func TestGroupListUsers(t *testing.T) {
	for _, c := range GetClients() {
	
		users, err := c.NewGroupListUsers().GroupID("_moonbeam").Do(context.Background())
		if err != nil {
			t.Logf("%+v\n", users)
			t.Error(err)
		}
		
		then.AssertThat(t, users.Code, is.EqualTo("Success"))
		then.AssertThat(t, users.Data.Items, has.Length(4))
		then.AssertThat(t, users.Data.Items[0].ID, is.EqualTo("_airworthy"))
		then.AssertThat(t, users.Data.Items[0].Email, is.EqualTo("airworthy.moonbeam@gmail.com"))
		then.AssertThat(t, users.Data.Items[0].GivenName, is.EqualTo("Airworthy"))
		then.AssertThat(t, users.Data.Items[0].FamilyName, is.EqualTo("Moonbeam"))
		then.AssertThat(t, *users.Data.Items[0].Verified, is.True())
		then.AssertThat(t, *users.Data.Items[0].Invited, is.False())
		then.AssertThat(t, users.Data.Items[0].Picture, is.Not(is.Empty()))
		then.AssertThat(t, users.Data.Items[0].Phone, is.EqualTo("+1 123 456 7890"))
		then.AssertThat(t,  users.Data.Items[0].LoggedInAt.IsZero(), is.True())
		then.AssertThat(t, users.Data.Items[0].CreatedAt.IsZero(), is.False())
		then.AssertThat(t, users.Data.NextCursor, is.Empty())
    }
}