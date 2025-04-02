package fingerprints_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
	
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestNewCertificateConnectionFingerprintApproveMock(t *testing.T) {
	// arrange
	validatedBy := "user_name"
	validatedDate := "validated_date"
	testConnectionId := "connection_id"
	testHash := "hash"
	testPublicKey := "public_key"

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, fmt.Sprintf("/v1/connections/%v/fingerprints", testConnectionId)).ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)

			testutils.AssertEqual(t, len(body), 2)
			testutils.AssertEqual(t, body["hash"], testHash)
			testutils.AssertEqual(t, body["public_key"], testPublicKey)

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
				`, testHash, testPublicKey, validatedBy, validatedDate))

			return response, nil
		})

	// act & assert
	response, err := ftClient.NewCertificateConnectionFingerprintApprove().
		ConnectionID(testConnectionId).
		Hash(testHash).
		PublicKey(testPublicKey).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)

	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.Hash, testHash)
	testutils.AssertEqual(t, response.Data.PublicKey, testPublicKey)
	testutils.AssertEqual(t, response.Data.ValidatedBy, validatedBy)
	testutils.AssertEqual(t, response.Data.ValidatedDate, validatedDate)
	testutils.AssertNotEmpty(t, response.Message)
}
