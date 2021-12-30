package fivetran_test

import (
	"context"
	"testing"
)

func TestNewCertificateDestinationFingerprintApproveE2E(t *testing.T) {
	for version, c := range Clients {
		t.Run(version, func(t *testing.T) {
			destinationId := CreateTempDestination(t)
			response, err := c.NewCertificateDestinationFingerprintApprove().
				DestinationID(destinationId).
				Hash("test_hash").
				PublicKey("test_public_key").
				Do(context.Background())

			if err != nil {
				t.Logf("%+v\n", response)
				t.Error(err)
			}

			AssertEqual(t, response.Code, "Success")
			AssertEqual(t, response.Message, "The fingerprint has been approved")
		})
	}
}
