package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestNewCertificateConnectorFingerprintApproveMock(t *testing.T) {
	//t.Skip("Endpoints redesigned. Test temporary disabled")
	// arrange
	validatedBy := "user_name"
	validatedDate := "validated_date"
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPost, fmt.Sprintf("/v1/connectors/%v/fingerprints", TEST_CONNECTOR_ID)).ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := requestBodyToJson(t, req)

			assertEqual(t, len(body), 2)
			assertEqual(t, body["hash"], TEST_HASH)
			assertEqual(t, body["public_key"], TEST_PUBLIC_KEY)

			response := mock.NewResponse(req, http.StatusCreated, fmt.Sprintf(`
				{
					"code": "Success", 
					"message": "The fingerprint has been approved",
					"data": {
						"hash": "%v",
						"public_key": "%v",
						"validated_by": "%v",
						"validated_date": "%v"
					} 
				}
				`, TEST_HASH, TEST_PUBLIC_KEY, validatedBy, validatedDate))

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
	assertEqual(t, response.Data.Hash, TEST_HASH)
	assertEqual(t, response.Data.PublicKey, TEST_PUBLIC_KEY)
	assertEqual(t, response.Data.ValidatedBy, validatedBy)
	assertEqual(t, response.Data.ValidatedDate, validatedDate)
	assertNotEmpty(t, response.Message)
}
