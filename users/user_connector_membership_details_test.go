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

func TestUserConnectorDetailsServiceDo(t *testing.T) {
	// arrange

	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/users/user_id/connectors/connector_id").
		ThenCall(func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareUserConnectorDetailsResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewUserConnectorMembershipDetails().
		UserId("user_id").
		ConnectorId("connector_id").
		Do(context.Background())

	// assert
	if err != nil {
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertUserConnectorDetailsResponse(t, response)
}

func prepareUserConnectorDetailsResponse() string {
	return fmt.Sprintf(`{
    		"code": "Success",
    		"data": {
          		"id": "connector_id",
          		"role": "Connector Administrator",
          		"created_at": "2020-05-25T15:26:47.306509Z"
    		}
		}`)
}

func assertUserConnectorDetailsResponse(t *testing.T, response users.UserConnectorMembershipDetailsResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.ConnectorId, "connector_id")
	testutils.AssertEqual(t, response.Data.Role, "Connector Administrator")
	testutils.AssertEqual(t, response.Data.CreatedAt, "2020-05-25T15:26:47.306509Z")
}