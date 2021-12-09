package tests

import (
	"context"
	"testing"
	"github.com/corbym/gocrest/then"
	"github.com/corbym/gocrest/is"
)

func TestDestinationSetupTest(t *testing.T) {
	for _, c := range GetClients() {
		response, err := c.NewDestinationSetupTests().DestinationID("_test").
		TrustCertificates(true).
		TrustFingerprints(true).
		Do(context.Background())
		
		if err != nil {
			t.Logf("%+v\n", response)
			t.Error(err)
		}

		then.AssertThat(t, response.Code, is.EqualTo("Success"))
		then.AssertThat(t, response.Message, is.EqualTo("Setup tests have been completed"))
	}
}
