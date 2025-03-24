package certificates_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestNewCertificateConnectionCertificateApproveMock(t *testing.T) {
	// arrange
	validatedBy := "user_name"
	validatedDate := "validated_date"
	testConnectionId := "connection_id"
	testHash := "hash"
	testEncodedCert := "encoded_cert"

	testPublicKey := "test_public_key"
	testName := "name"
	testType := "type"

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, fmt.Sprintf("/v1/connections/%v/certificates", testConnectionId)).ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)

			testutils.AssertEqual(t, len(body), 2)
			testutils.AssertEqual(t, body["hash"], testHash)
			testutils.AssertEqual(t, body["encoded_cert"], testEncodedCert)

			response := mock.NewResponse(req, http.StatusCreated, fmt.Sprintf(`
				{
					"code": "Success", 
					"message": "The certificate has been approved",
					"data": {
						"hash": "%v",
						"public_key": "%v",
						"name": "%v",
						"type": "%v",
						"validated_by": "%v",
						"validated_date": "%v"
					} 
				}
				`, testHash, testPublicKey, testName, testType, validatedBy, validatedDate))

			return response, nil
		})

	// act & assert
	response, err := ftClient.NewCertificateConnectionCertificateApprove().
		ConnectionID(testConnectionId).
		Hash(testHash).
		EncodedCert(testEncodedCert).
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
	testutils.AssertEqual(t, response.Data.Name, testName)
	testutils.AssertEqual(t, response.Data.Type, testType)
	testutils.AssertEqual(t, response.Data.ValidatedDate, validatedDate)
	testutils.AssertNotEmpty(t, response.Message)
}
