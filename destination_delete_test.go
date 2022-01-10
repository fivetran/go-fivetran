package fivetran_test

import (
	"context"
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
	AssertEqual(t, deleted.Message, "Destination with id '"+destinationId+"' has been deleted")
}
