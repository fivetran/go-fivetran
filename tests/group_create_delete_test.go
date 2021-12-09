package tests

import (
	"context"
	"testing"

	"github.com/corbym/gocrest/then"
	"github.com/corbym/gocrest/is"
)

func TestGroupCreateDelete(t *testing.T) {
	for _, c := range GetClients() {
		//create
		created, err := c.NewGroupCreate().Name("test").Do(context.Background())
		if err != nil {
			t.Logf("%+v\n", created)
			t.Error(err)
		}

		then.AssertThat(t, created.Code, is.EqualTo("Success"))
		then.AssertThat(t, created.Data.ID, is.Not(is.Empty()))
		then.AssertThat(t, created.Data.Name, is.EqualTo("test"))

		//delete
		deleted, err := c.NewGroupDelete().GroupID(created.Data.ID).Do(context.Background())
		if err != nil {
			t.Logf("%+v\n", deleted)
			t.Error(err)
		}
		then.AssertThat(t, deleted.Code, is.EqualTo("Success"))
		then.AssertThat(t, deleted.Message, is.EqualTo("Group with id '" + created.Data.ID + "' has been deleted"))
    }
}