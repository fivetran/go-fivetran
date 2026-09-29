package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewCertificateDestinationFingerprintApproveE2E(t *testing.T) {
	destinationId := testutils.CreateTempDestination(t)
	response, err := testutils.Client.NewCertificateDestinationFingerprintApprove().
		DestinationID(destinationId).
		Hash(testutils.CertificateHash).
		PublicKey(testutils.EncodedCertificate).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)
	testutils.AssertEqual(t, response.Data.ValidatedBy, testutils.PredefinedUserGivenName+" "+testutils.PredefinedUserFamilyName)
}
