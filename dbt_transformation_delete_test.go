package fivetran_test

import (
	"context"
	"strings"
	"testing"
)

func TestNewDbtTransformationDeleteE2E(t *testing.T) {
	t.Skip("Skipping test until we get more info on dbt transformations data which can be used for testing")

	dbtTransformationId := CreateDbtTransformation(t)
	deleted, err := Client.NewDestinationDelete().
		DestinationID(dbtTransformationId).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}

	AssertEqual(t, deleted.Code, "Success")
	AssertNotEmpty(t, deleted.Message)
	AssertEqual(t, strings.Contains(deleted.Message, dbtTransformationId), true)
}
