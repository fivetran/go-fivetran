package tests

import (
	"context"
	"testing"

	"github.com/corbym/gocrest/then"
	"github.com/corbym/gocrest/is"
)


func TestGroupAddRemoveUser(t *testing.T) {
	for _, c := range GetClients() {
		//add
		created, err := c.NewGroupAddUser().GroupID("_test").Email("accountworthy.moonbeam@gmail.com").Role("ReadOnly").Do(context.Background())
		if err != nil {
			t.Logf("%+v\n", created)
			t.Error(err)
		}

		then.AssertThat(t, created.Code, is.EqualTo("Success"))
		then.AssertThat(t, created.Message, is.EqualTo("User has been invited to the group"))

		//delete
		deleted, err := c.NewGroupRemoveUser().GroupID("_test").UserID("_accountworthy").Do(context.Background())
		if err != nil {
			t.Logf("%+v\n", deleted)
			t.Error(err)
		}
		then.AssertThat(t, deleted.Code, is.EqualTo("Success"))
		then.AssertThat(t, deleted.Message, is.EqualTo("User with id '_accountworthy' has been removed from the group"))
	}
}