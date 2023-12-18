package users_test

import (
    "context"
    "fmt"
    "net/http"
    "testing"

	"github.com/fivetran/go-fivetran/common"
    
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

const (
	EXPECTED_USER_DELETE_RESPONSE_CODE = "Success"
	EXPECTED_USER_DELETE_USER_ID       = "user_id"
)

func TestUserDeleteServiceDo(t *testing.T) {
	ftClient, mockClient := testutils.CreateTestClient()

	handler := mockClient.When(http.MethodDelete, fmt.Sprintf("/v1/users/%s", EXPECTED_USER_DELETE_USER_ID)).ThenCall(
		func(req *http.Request) (*http.Response, error) {
			responseData := prepareUserDeleteResponse()
			response := mock.NewResponse(req, http.StatusOK, responseData)
			return response, nil
		},
	)

	service := ftClient.NewUserDelete().UserID(EXPECTED_USER_DELETE_USER_ID)

	response, err := service.Do(context.Background())

	if err != nil {
		t.Error(err)
	}

	assertUserDeleteResponse(t, response, "Success", "User has been deleted")

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
}

func TestUserDeleteServiceDoMissingUserID(t *testing.T) {
	ftClient, _ := testutils.CreateTestClient()

	service := ftClient.NewUserDelete()

	_, err := service.Do(context.Background())

	expectedError := fmt.Errorf("missing required userID")
	testutils.AssertEqual(t, err, expectedError)
}

func prepareUserDeleteResponse() string {
	return fmt.Sprintf(`{"code": "%s", "message": "User has been deleted"}`, EXPECTED_USER_DELETE_RESPONSE_CODE)
}

func assertUserDeleteResponse(t *testing.T, actual common.CommonResponse, code string, message string) {
	testutils.AssertEqual(t, actual.Code, code)
	testutils.AssertEqual(t, actual.Message, message)
}
