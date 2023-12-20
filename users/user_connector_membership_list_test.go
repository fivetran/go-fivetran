package users_test

import (
    "context"
    "fmt"
    "net/http"
    "testing"

	"github.com/fivetran/go-fivetran/users"
    
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestUserConnectorListServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/users/user_id/connectors").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareUserConnectorListResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewUserConnectorMembershipsList().
		UserId("user_id").
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertUserConnectorListResponse(t, response)
}

func prepareUserConnectorListResponse() string {
	return fmt.Sprintf(`{
    		"code": "Success",
    		"data": {
      		"items": [
        		{
          			"id": "connector_id_1",
          			"role": "Connector Administrator",
          			"created_at": "2020-05-25T15:26:47.306509Z"
        		},
        		{
          			"id": "connector_id_2",
          			"role": "Connector Reviewer",
          			"created_at": "2020-05-25T15:26:47.306509Z"
        		}
      		],
      		"next_cursor": null
    		}
		}`)
}

func assertUserConnectorListResponse(t *testing.T, response users.UserConnectorMembershipsListResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.Items[0].ConnectorId, "connector_id_1")
	testutils.AssertEqual(t, response.Data.Items[0].Role, "Connector Administrator")
	testutils.AssertEqual(t, response.Data.Items[0].CreatedAt, "2020-05-25T15:26:47.306509Z")

	testutils.AssertEqual(t, response.Data.Items[1].ConnectorId, "connector_id_2")
	testutils.AssertEqual(t, response.Data.Items[1].Role, "Connector Reviewer")
	testutils.AssertEqual(t, response.Data.Items[1].CreatedAt, "2020-05-25T15:26:47.306509Z")
}