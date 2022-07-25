package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestNewCertificateConnectorFingerprintApproveMock(t *testing.T) {

	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/fingerprints").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := requestBodyToJson(t, req)

			assertEqual(t, len(body), 3)
			assertEqual(t, body["connector_id"], TEST_CONNECTOR_ID)
			assertEqual(t, body["hash"], TEST_HASH)
			assertEqual(t, body["public_key"], TEST_PUBLIC_KEY)

			response := mock.NewResponse(req, http.StatusOK,
				`{"code": "Success", "message": "The fingerprint has been approved"}`)

			return response, nil
		})

	// act & assert
	response, err := ftClient.NewCertificateConnectorFingerprintApprove().
		ConnectorID(TEST_CONNECTOR_ID).
		Hash(TEST_HASH).
		PublicKey(TEST_PUBLIC_KEY).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)

	assertEqual(t, response.Code, "Success")
	assertNotEmpty(t, response.Message)
}
