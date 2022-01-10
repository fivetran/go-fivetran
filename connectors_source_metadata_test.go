package fivetran_test

import (
	"context"
	"testing"
)

func TestNewConnectorsSourceMetadataE2E(t *testing.T) {
	meta, err := Client.NewConnectorsSourceMetadata().Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", meta)
		t.Error(err)
	}

	AssertEqual(t, meta.Code, "Success")
	AssertNotEmpty(t, meta.Data.Items)
	AssertNotEmpty(t, meta.Data.Items[0].ID)
	AssertNotEmpty(t, meta.Data.Items[0].Name)
	AssertNotEmpty(t, meta.Data.Items[0].Type)
	AssertNotEmpty(t, meta.Data.Items[0].Description)
	AssertNotEmpty(t, meta.Data.Items[0].IconURL)
}
