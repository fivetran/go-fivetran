package fingerprints_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
	"github.com/fivetran/go-fivetran/tests"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestNewCertificateDestinationFingerprintApproveMock(t *testing.T) {
	// arrange
	validatedBy := "user_name"
	validatedDate := "validated_date"
	testDestinationId := "destination_id"
	testHash := "hash"
	testPublicKey := "public_key"

	ftClient, mockClient := tests.CreateTestClient()
	handler := mockClient.When(http.MethodPost, fmt.Sprintf("/v1/destinations/%v/fingerprints", testDestinationId)).ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := tests.RequestBodyToJson(t, req)

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
	response, err := ftClient.NewCertificateDestinationFingerprintApprove().
		DestinationID(testDestinationId).
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
