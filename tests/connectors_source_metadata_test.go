package tests

import (
	"context"
	"testing"

	"github.com/corbym/gocrest/has"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
)


func TestConnectorsSourceMetadata(t *testing.T) {
	for v, c := range GetClients() {
		//supported only for v1 version
		if v != "v1" {
			return;
		}

		meta, err := c.NewConnectorsSourceMetadata().Limit(3).Do(context.Background())

		if err != nil {
			t.Logf("%+v\n", meta)
			t.Error(err)
		}

		then.AssertThat(t, meta.Code, is.EqualTo("Success"))
		then.AssertThat(t, meta.Data.Items, has.Length(3))
		then.AssertThat(t, meta.Data.Items[0].ID, is.Not(is.Empty()))
		then.AssertThat(t, meta.Data.Items[0].Name, is.Not(is.Empty()))
		then.AssertThat(t, meta.Data.Items[0].Type, is.Not(is.Empty()))
		then.AssertThat(t, meta.Data.Items[0].Description, is.Not(is.Empty()))
		then.AssertThat(t, meta.Data.Items[0].IconURL, is.Not(is.Empty()))
	}
}