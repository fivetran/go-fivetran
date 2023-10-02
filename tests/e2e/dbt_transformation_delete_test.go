package fivetran_test

import (
	"context"
	"strings"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewDbtTransformationDeleteE2E(t *testing.T) {
	t.Skip("Skipping test until we get more info on dbt transformations data which can be used for testing")

	dbtTransformationId := testutils.CreateDbtTransformation(t)
	deleted, err := testutils.Client.NewDestinationDelete().
		DestinationID(dbtTransformationId).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}

	testutils.AssertEqual(t, deleted.Code, "Success")
	testutils.AssertNotEmpty(t, deleted.Message)
	testutils.AssertEqual(t, strings.Contains(deleted.Message, dbtTransformationId), true)
}
