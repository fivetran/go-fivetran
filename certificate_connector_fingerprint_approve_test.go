package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewCertificateConnectorFingerprintApproveE2E(t *testing.T) {
	connectorId := testutils.CreateTempConnector(t)
	response, err := testutils.Client.NewCertificateConnectorFingerprintApprove().
		ConnectorID(connectorId).
		Hash("test_hash").
		PublicKey("test_public_key").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)
}
