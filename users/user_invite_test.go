package users_test

import (
    "context"
    "fmt"
    "net/http"
    "testing"

	"github.com/fivetran/go-fivetran/users"
	"github.com/fivetran/go-fivetran/common"
    
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

const (
	EXPECTED_USER_UPDATE_RESPONSE_CODE = "Success"
	EXPECTED_USER_UPDATE_USER_ID       = "firewood_martial"
	EXPECTED_USER_UPDATE_GIVEN_NAME    = "John"
	EXPECTED_USER_UPDATE_FAMILY_NAME   = "White"
	EXPECTED_USER_UPDATE_PHONE         = "+123456789"
	EXPECTED_USER_UPDATE_PICTURE       = "http://mycompany.com/avatars/john_white.png"
	EXPECTED_USER_UPDATE_ROLE          = "Account Administrator"
	EXPECTED_USER_UPDATE_EMAIL         = "john.white@mycompany.com"
)

func TestUserUpdateServiceDo(t *testing.T) {
	ftClient, mockClient := testutils.CreateTestClient()

	handler := mockClient.When(http.MethodPatch, fmt.Sprintf("/v1/users/%s", EXPECTED_USER_UPDATE_USER_ID)).ThenCall(
		func(req *http.Request) (*http.Response, error) {
			responseData := prepareUserUpdateResponse()
			response := mock.NewResponse(req, http.StatusOK, responseData)
			return response, nil
		},
	)

	service := ftClient.NewUserUpdate().
		UserID(EXPECTED_USER_UPDATE_USER_ID).
		GivenName(EXPECTED_USER_UPDATE_GIVEN_NAME).
		FamilyName(EXPECTED_USER_UPDATE_FAMILY_NAME).
		Phone(EXPECTED_USER_UPDATE_PHONE).
		Picture(EXPECTED_USER_UPDATE_PICTURE).
		Role(EXPECTED_USER_UPDATE_ROLE)

	response, err := service.Do(context.Background())

	if err != nil {
		t.Error(err)
	}

	expectedResponse := prepareExpectedUserUpdateResponse()
	assertUserUpdateResponse(t, response, expectedResponse)

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
}

func prepareUserUpdateResponse() string {
	return fmt.Sprintf(`{
		"code": "%s",
		"message": "User has been invited to the account",
		"data": {
			"id": "%s",
			"email": "%s",
			"given_name": "%s",
			"family_name": "%s",
			"verified": false,
			"invited": true,
			"picture": "%s",
			"phone": "%s",
			"logged_in_at": null,
			"created_at": "2019-01-20T16:03:36.786936Z",
			"active": true,
			"role": "%s"
		}
	}`, EXPECTED_USER_UPDATE_RESPONSE_CODE,
		EXPECTED_USER_UPDATE_USER_ID, EXPECTED_USER_UPDATE_EMAIL,
		EXPECTED_USER_UPDATE_GIVEN_NAME,
		EXPECTED_USER_UPDATE_FAMILY_NAME,
		EXPECTED_USER_UPDATE_PICTURE,
		EXPECTED_USER_UPDATE_PHONE,
		EXPECTED_USER_UPDATE_ROLE)
}

func prepareExpectedUserUpdateResponse() users.UserDetailsResponse {
	var verifyFlag = false
	return users.UserDetailsResponse{
		CommonResponse: common.CommonResponse{
			Code:    EXPECTED_USER_UPDATE_RESPONSE_CODE,
			Message: "User has been invited to the account",
		},
		Data: users.UserDetailsData{
			ID:         EXPECTED_USER_UPDATE_USER_ID,
			Email:      EXPECTED_USER_UPDATE_EMAIL,
			GivenName:  EXPECTED_USER_UPDATE_GIVEN_NAME,
			FamilyName: EXPECTED_USER_UPDATE_FAMILY_NAME,
			Verified:   &verifyFlag,
			Invited:    &verifyFlag,
			Picture:    EXPECTED_USER_UPDATE_PICTURE,
			Phone:      EXPECTED_USER_UPDATE_PHONE,
			Role:       EXPECTED_USER_UPDATE_ROLE,
		},
	}
}

func assertUserUpdateResponse(t *testing.T, actual, expected users.UserDetailsResponse) {
	testutils.AssertEqual(t, actual.Code, expected.Code)
	testutils.AssertEqual(t, actual.Message, expected.Message)
	testutils.AssertEqual(t, actual.Data.ID, expected.Data.ID)
	testutils.AssertEqual(t, actual.Data.Email, expected.Data.Email)
	testutils.AssertEqual(t, actual.Data.GivenName, expected.Data.GivenName)
	testutils.AssertEqual(t, actual.Data.FamilyName, expected.Data.FamilyName)
	testutils.AssertEqual(t, actual.Data.Verified, expected.Data.Verified)
	testutils.AssertEqual(t, actual.Data.Picture, expected.Data.Picture)
	testutils.AssertEqual(t, actual.Data.Phone, expected.Data.Phone)
	testutils.AssertEqual(t, actual.Data.Role, expected.Data.Role)
}
