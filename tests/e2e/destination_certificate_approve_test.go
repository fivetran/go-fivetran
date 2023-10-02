package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewCertificateDestinationCertificateApproveE2E(t *testing.T) {
	destinationId := testutils.CreateTempDestination(t)
	response, err := testutils.Client.NewCertificateDestinationCertificateApprove().
		DestinationID(destinationId).
		Hash(testutils.CertificateHash).
		EncodedCert(testutils.EncodedCertificate).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)
	testutils.AssertEqual(t, response.Data.ValidatedBy, testutils.PredefinedUserGivenName+" "+testutils.PredefinedUserFamilyName)
}
