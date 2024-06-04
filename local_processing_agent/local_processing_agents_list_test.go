package localprocessingagent_test

import (
    "context"
    "net/http"
    "testing"
		"github.com/fivetran/go-fivetran/local_processing_agent"
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestLocalProcessingAgentListServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/local-processing-agents").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareLocalProcessingAgentListResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewLocalProcessingAgentList().
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertLocalProcessingAgentListResponse(t, response)
}

func prepareLocalProcessingAgentListResponse() string {
	return `{
					    "code": "Success",
    					"data": {
					        "items": [
					          {
          					  "id": "id",
            					"display_name": "display_name",
            					"group_id": "group_id",
            					"registered_at": "1970-01-01T00:00:00.000000Z",
            					"usage": [
              					{
                					"connection_id": "connection_id",
                					"schema": "schema",
                					"service": "service"
              					}
            					]
          					}
        					],
        					"next_cursor": null
    					}
					}`
}

func assertLocalProcessingAgentListResponse(t *testing.T, response localprocessingagent.LocalProcessingAgentListResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.Items[0].Id, "id")
	testutils.AssertEqual(t, response.Data.Items[0].DisplayName, "display_name")
	testutils.AssertEqual(t, response.Data.Items[0].GroupId, "group_id")
	testutils.AssertEqual(t, response.Data.Items[0].RegisteredAt, "1970-01-01T00:00:00.000000Z")
	testutils.AssertEqual(t, response.Data.Items[0].Usage[0].ConnectionId, "connection_id")
	testutils.AssertEqual(t, response.Data.Items[0].Usage[0].Schema, "schema")
	testutils.AssertEqual(t, response.Data.Items[0].Usage[0].Service, "service")
}