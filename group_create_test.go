package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewGroupCreateE2E(t *testing.T) {
	created, err := testutils.Client.NewGroupCreate().Name("test").Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertNotEmpty(t, created.Data.ID)
	testutils.AssertEqual(t, created.Data.Name, "test")

	t.Cleanup(func() { testutils.DeleteGroup(t, created.Data.ID) })
}
