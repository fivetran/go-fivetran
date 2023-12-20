package fingerprints_test

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
	
	"github.com/fivetran/go-fivetran/tests/mock"
)

func TestNewDestinationFingerprintsListMock(t *testing.T) {
	// arrange
	validatedBy := "user_name"
	validatedDate := "validated_date"
	testDestinationId := "destination_id"
	testHash := "hash"

	testPublicKey := "test_public_key"
	nextCursor := "next_cursor"
	cursor := "cursor"
	limit := 1

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, fmt.Sprintf("/v1/destinations/%v/fingerprints", testDestinationId)).ThenCall(
		func(req *http.Request) (*http.Response, error) {
			var query = req.URL.Query()
			testutils.AssertEqual(t, query.Get("cursor"), cursor)
			testutils.AssertEqual(t, query.Get("limit"), strconv.Itoa(limit))
			response := mock.NewResponse(req, http.StatusOK, fmt.Sprintf(`
				{
					"code": "Success", 
					"message": "",
					"data": {
						"items": [
							{
								"hash": "%v",
								"public_key": "%v",
								"validated_by": "%v",
								"validated_date": "%v"
							}
						],
						"next_cursor": "%v"
					} 
				}
				`, testHash, testPublicKey, validatedBy, validatedDate, nextCursor))

			return response, nil
		})

	// act & assert
	response, err := ftClient.NewDestinationFingerprintsList().
		DestinationID(testDestinationId).
		Cursor(cursor).
		Limit(limit).
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
	testutils.AssertEqual(t, response.Data.Items[0].Hash, testHash)
	testutils.AssertEqual(t, response.Data.Items[0].PublicKey, testPublicKey)
	testutils.AssertEqual(t, response.Data.Items[0].ValidatedBy, validatedBy)
	testutils.AssertEqual(t, response.Data.Items[0].ValidatedDate, validatedDate)
	testutils.AssertEqual(t, response.Data.NextCursor, nextCursor)
	testutils.AssertEmpty(t, response.Message)
}
