package tests

import (
	"context"
	"testing"

	"github.com/corbym/gocrest/then"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/has"
)

func TestGetGroupsList(t *testing.T) {
	for _, c := range GetClients() {
		result, err :=  c.NewGroupsList().Do(context.Background())
		if err != nil {
			t.Logf("%+v\n", result)
			t.Error(err)
		}
		
		then.AssertThat(t, result.Code, is.EqualTo("Success"))
		then.AssertThat(t, result.Data.Items, has.Length(3))
		then.AssertThat(t, result.Message, is.Empty())
		then.AssertThat(t, result.Data.Items[0].ID, is.EqualTo("_moonbeam"))
		then.AssertThat(t, result.Data.Items[0].Name, is.EqualTo("Production"))
		then.AssertThat(t, result.Data.Items[0].CreatedAt, is.Not(is.Nil()))
		then.AssertThat(t, result.Data.Items[1].ID, is.EqualTo("_moonbeam_bright"))
		then.AssertThat(t, result.Data.Items[1].Name, is.EqualTo("Shine Bright"))
		then.AssertThat(t, result.Data.Items[1].CreatedAt, is.Not(is.Nil()))
		then.AssertThat(t, result.Data.Items[2].ID, is.EqualTo("_test"))
		then.AssertThat(t, result.Data.Items[2].Name, is.EqualTo("Staging"))
		then.AssertThat(t, result.Data.Items[2].CreatedAt, is.Not(is.Nil()))
		then.AssertThat(t, result.Data.NextCursor, is.Empty())
    }
}


