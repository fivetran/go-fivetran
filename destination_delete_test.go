package fivetran_test

import (
	"context"
	"testing"
)

func TestNewDestinationDeleteIntegration(t *testing.T) {
	for version, c := range Clients {
		t.Run(version, func(t *testing.T) {
			destinationId := CreateDestination(t)
			deleted, err := c.NewDestinationDelete().DestinationID(destinationId).Do(context.Background())

			if err != nil {
				t.Logf("%+v\n", deleted)
				t.Error(err)
			}

			AssertEqual(t, deleted.Code, "Success")
			AssertEqual(t, deleted.Message, "Destination with id '"+destinationId+"' has been deleted")
		})
	}
}
