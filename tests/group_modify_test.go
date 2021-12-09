package tests

import (
	"context"
	"testing"

	"github.com/corbym/gocrest/then"
	"github.com/corbym/gocrest/is"
)
func TestGroupModify(t *testing.T) {
	for _, c := range GetClients() {
		//arrange
		created, err := c.NewGroupCreate().Name("test").Do(context.Background())
		if err != nil {
			t.Logf("%+v\n", created)
			t.Error(err)
		}

		//act & assert
		result, err :=  c.NewGroupModify().GroupID(created.Data.ID).Name("test_new").Do(context.Background())
		if err != nil {
			t.Logf("%+v\n", result)
			t.Error(err)
		}
		then.AssertThat(t, result.Code, is.EqualTo("Success"))
		then.AssertThat(t, result.Data.ID, is.EqualTo(result.Data.ID))
		then.AssertThat(t, result.Data.Name, is.EqualTo("test_new"))
		then.AssertThat(t, result.Data.CreatedAt, is.Not(is.Nil()))

		//cleanup
		deleted, err := c.NewGroupDelete().GroupID(created.Data.ID).Do(context.Background())
		if err != nil {
			t.Logf("%+v\n", deleted)
			t.Error(err)
		}
    }
}