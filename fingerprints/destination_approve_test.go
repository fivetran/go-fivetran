package fingerprints_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewCertificateDestinationFingerprintApproveE2E(t *testing.T) {
	destinationId := testutils.CreateTempDestination(t)
	response, err := testutils.Client.NewCertificateDestinationFingerprintApprove().
		DestinationID(destinationId).
		Hash("test_hash").
		PublicKey("test_public_key").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)
	testutils.AssertEqual(t, response.Data.ValidatedBy, testutils.PredefinedUserId)
}
