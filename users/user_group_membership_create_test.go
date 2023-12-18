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
	TEAM_GROUP_ROLE = "Destination Administrator"
)

func TestNewUserGroupCreate(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/users/user_id/groups").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertUserGroupCreateRequest(t, body)
			response := mock.NewResponse(req, http.StatusCreated, prepareUserGroupCreateResponse())
			return response, nil
		})

	// act
	response, err := ftClient.NewUserGroupMembershipCreate().
		UserId("user_id").
		GroupId("group_id").
		Role(TEAM_GROUP_ROLE).
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

	assertUserGroupCreateResponse(t, response)
}

func prepareUserGroupCreateResponse() string {
	return fmt.Sprintf(
		`{
            "code": "Created",
            "message": "Group membership has been created",
            "data": {
                "id": "group_id",
                "role": "%v",
                "created_at": "2021-09-29T10:50:51.397153Z"
            }
        }`,
		TEAM_GROUP_ROLE,
	)
}

func assertUserGroupCreateRequest(t *testing.T, request map[string]interface{}) {
	testutils.AssertKey(t, "id", request, "group_id")
	testutils.AssertKey(t, "role", request, TEAM_GROUP_ROLE)
}

func assertUserGroupCreateResponse(t *testing.T, response users.UserGroupMembershipCreateResponse) {
	testutils.AssertEqual(t, response.Code, "Created")
	testutils.AssertNotEmpty(t, response.Message)
	testutils.AssertEqual(t, response.Data.GroupId, "group_id")
	testutils.AssertEqual(t, response.Data.Role, TEAM_GROUP_ROLE)
	testutils.AssertNotEmpty(t, response.Data.CreatedAt)
}
