package fivetran_test

import (
	"context"
	"testing"
)

func TestNewCertificateDestinationCertificateApproveE2E(t *testing.T) {
	for version, c := range Clients {
		t.Run(version, func(t *testing.T) {
			destinationId := CreateTempDestination(t)
			response, err := c.NewCertificateDestinationCertificateApprove().
				DestinationID(destinationId).
				Hash(CertificateHash).
				EncodedCert(EncodedCertificate).
				Do(context.Background())

			if err != nil {
				t.Logf("%+v\n", response)
				t.Error(err)
			}

			AssertEqual(t, response.Code, "Success")
			AssertEqual(t, response.Message, "The certificate has been approved")
		})
	}
}
