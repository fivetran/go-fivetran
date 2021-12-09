package tests

import (
	"context"
	"testing"
	"github.com/corbym/gocrest/then"
	"github.com/corbym/gocrest/is"
)

func TestCertificateDestinationFingerprintApprove(t *testing.T) {
	for _, c := range GetClients() {
		response, err := c.NewCertificateDestinationFingerprintApprove().
		DestinationID("_test").
		Hash("test_hash").
		PublicKey("test_public_key").
		Do(context.Background())
		
		if err != nil {
			t.Logf("%+v\n", response)
			t.Error(err)
		}

		then.AssertThat(t, response.Code, is.EqualTo("Success"))
		then.AssertThat(t, response.Message, is.EqualTo("The fingerprint has been approved"))
	}
}