package fivetran_test

import (
	"context"
	"testing"
)

func TestNewCertificateConnectorCertificateApproveE2E(t *testing.T) {
	connectorId := CreateTempConnector(t)
	response, err := Client.NewCertificateConnectorCertificateApprove().
		ConnectorID(connectorId).
		Hash(CertificateHash).
		EncodedCert(EncodedCertificate).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	AssertEqual(t, response.Code, "Success")
	AssertEqual(t, response.Message, "The certificate has been approved")
}
