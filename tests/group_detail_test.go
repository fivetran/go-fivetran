package tests

import (
	"context"
	"testing"

	"github.com/corbym/gocrest/then"
	"github.com/corbym/gocrest/is"
)

func TestGetGroupDetails(t *testing.T) {
	for _, c := range GetClients() {
		result, err :=  c.NewGroupDetails().GroupID("_moonbeam").Do(context.Background())
		if err != nil {
			t.Logf("%+v\n", result)
			t.Error(err)
		}
		then.AssertThat(t, result.Code, is.EqualTo("Success"))
		then.AssertThat(t, result.Data.ID, is.EqualTo("_moonbeam"))
		then.AssertThat(t, result.Data.Name, is.EqualTo("Production"))
		then.AssertThat(t, result.Data.CreatedAt, is.Not(is.Nil()))
    }
}