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

const (
	TEAM_CONNECTION_ROLE = "Connection Collaborator"
)

func TestNewUserConnectionCreate(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/users/user_id/connections").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertUserConnectionCreateRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareUserConnectionCreateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewUserConnectionMembershipCreate().
		UserId("user_id").
		ConnectionId("connection_id").
		Role(TEAM_CONNECTION_ROLE).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	// assert
	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)

	assertUserConnectionCreateResponse(t, response)
}

func prepareUserConnectionCreateResponse() string {
	return fmt.Sprintf(
		`{
            "code": "Created",
            "message": "Connection membership has been created",
            "data": {
                "id": "connection_id",
                "role": "%v",
                "created_at": "2021-09-29T10:50:51.397153Z"
            }
        }`,
		TEAM_CONNECTION_ROLE,
	)
}

func assertUserConnectionCreateRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "id", request, "connection_id")
	testutils.AssertKey(t, "role", request, TEAM_CONNECTION_ROLE)
}

func assertUserConnectionCreateResponse(t *testing.T, response users.UserConnectionMembershipCreateResponse) {
	testutils.AssertEqual(t, response.Code, "Created")
	testutils.AssertNotEmpty(t, response.Message)
	testutils.AssertEqual(t, response.Data.ConnectionId, "connection_id")
	testutils.AssertEqual(t, response.Data.Role, TEAM_CONNECTION_ROLE)
	testutils.AssertNotEmpty(t, response.Data.CreatedAt)
}
