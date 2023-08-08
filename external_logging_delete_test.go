package fivetran_test

import (
	"context"
	"strings"
	"testing"
)

func TestNewDestinationDeleteE2E(t *testing.T) {
	destinationId := CreateDestination(t)
	deleted, err := Client.NewDestinationDelete().DestinationID(destinationId).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}

	AssertEqual(t, deleted.Code, "Success")
	AssertNotEmpty(t, deleted.Message)
	AssertEqual(t, strings.Contains(deleted.Message, destinationId), true)
}
