package fivetran_test

import (
	"context"
	"testing"
)

func TestNewGroupCreateE2E(t *testing.T) {
	created, err := Client.NewGroupCreate().Name("test").Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	AssertEqual(t, created.Code, "Success")
	AssertNotEmpty(t, created.Data.ID)
	AssertEqual(t, created.Data.Name, "test")

	t.Cleanup(func() { DeleteGroup(t, created.Data.ID) })
}
