package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewConnectorsSourceMetadataE2E(t *testing.T) {
	meta, err := testutils.Client.NewConnectorsSourceMetadata().Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", meta)
		t.Error(err)
	}

	testutils.AssertEqual(t, meta.Code, "Success")
	testutils.AssertNotEmpty(t, meta.Data.Items)
	testutils.AssertNotEmpty(t, meta.Data.Items[0].ID)
	testutils.AssertNotEmpty(t, meta.Data.Items[0].Name)
	testutils.AssertNotEmpty(t, meta.Data.Items[0].Type)
	testutils.AssertNotEmpty(t, meta.Data.Items[0].Description)
	testutils.AssertNotEmpty(t, meta.Data.Items[0].IconURL)
}
