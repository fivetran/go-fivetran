package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewCertificateConnectionCertificateApproveE2E(t *testing.T) {
	ConnectionId := testutils.CreateTempConnection(t)
	response, err := testutils.Client.NewCertificateConnectionCertificateApprove().
		ConnectionID(ConnectionId).
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
