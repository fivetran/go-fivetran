package fingerprints_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
	
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestNewConnectorFingerprintDetailsMock(t *testing.T) {
	// arrange
	validatedBy := "user_name"
	validatedDate := "validated_date"
	testConnectorId := "connector_id"
	testHash := "hash"

	testPublicKey := "test_public_key"

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, fmt.Sprintf("/v1/connectors/%v/fingerprints/%v", testConnectorId, testHash)).ThenCall(

		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, fmt.Sprintf(`
				{
					"code": "Success", 
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
	response, err := ftClient.NewConnectorFingerprintDetails().
		ConnectorID(testConnectorId).
		Hash(testHash).
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
	testutils.AssertEmpty(t, response.Message)
}
