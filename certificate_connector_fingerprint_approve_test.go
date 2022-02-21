package fivetran_test

import (
	"context"
	"testing"
)

func TestNewCertificateConnectorFingerprintApproveE2E(t *testing.T) {
	connectorId := CreateTempConnector(t)
	response, err := Client.NewCertificateConnectorFingerprintApprove().
		ConnectorID(connectorId).
		Hash("test_hash").
		PublicKey("test_public_key").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	AssertEqual(t, response.Code, "Success")
	AssertNotEmpty(t, response.Message)
}
