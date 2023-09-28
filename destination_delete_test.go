package fivetran_test

import (
	"context"
	"strings"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewDestinationDeleteE2E(t *testing.T) {
	destinationId := testutils.CreateDestination(t)
	deleted, err := testutils.Client.NewDestinationDelete().DestinationID(destinationId).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}

	testutils.AssertEqual(t, deleted.Code, "Success")
	testutils.AssertNotEmpty(t, deleted.Message)
	testutils.AssertEqual(t, strings.Contains(deleted.Message, destinationId), true)
}
