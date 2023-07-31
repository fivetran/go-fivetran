package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

const (
	EXPECTED_USER_MODIFY_RESPONSE_CODE = "Success"
	EXPECTED_USER_MODIFY_USER_ID       = "firewood_martial"
	EXPECTED_USER_MODIFY_GIVEN_NAME    = "John"
	EXPECTED_USER_MODIFY_FAMILY_NAME   = "White"
	EXPECTED_USER_MODIFY_PHONE         = "+123456789"
	EXPECTED_USER_MODIFY_PICTURE       = "http://mycompany.com/avatars/john_white.png"
	EXPECTED_USER_MODIFY_ROLE          = "Account Administrator"
	EXPECTED_USER_MODIFY_EMAIL         = "john.white@mycompany.com"
)

func TestUserModifyServiceDo(t *testing.T) {
	ftClient, mockClient := CreateTestClient()

	handler := mockClient.When(http.MethodPatch, fmt.Sprintf("/v1/users/%s", EXPECTED_USER_MODIFY_USER_ID)).ThenCall(
		func(req *http.Request) (*http.Response, error) {
			responseData := prepareUserModifyResponse()
			response := mock.NewResponse(req, http.StatusOK, responseData)
			return response, nil
		},
	)

	service := ftClient.NewUserModify().
		UserID(EXPECTED_USER_MODIFY_USER_ID).
		GivenName(EXPECTED_USER_MODIFY_GIVEN_NAME).
		FamilyName(EXPECTED_USER_MODIFY_FAMILY_NAME).
		Phone(EXPECTED_USER_MODIFY_PHONE).
		Picture(EXPECTED_USER_MODIFY_PICTURE).
		Role(EXPECTED_USER_MODIFY_ROLE)

	response, err := service.Do(context.Background())

	if err != nil {
		t.Error(err)
	}

	expectedResponse := prepareExpectedUserModifyResponse()
	assertUserModifyResponse(t, response, expectedResponse)

	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)
}

func prepareUserModifyResponse() string {
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
	}`, EXPECTED_USER_MODIFY_RESPONSE_CODE,
		EXPECTED_USER_MODIFY_USER_ID, EXPECTED_USER_MODIFY_EMAIL,
		EXPECTED_USER_MODIFY_GIVEN_NAME,
		EXPECTED_USER_MODIFY_FAMILY_NAME,
		EXPECTED_USER_MODIFY_PICTURE,
		EXPECTED_USER_MODIFY_PHONE,
		EXPECTED_USER_MODIFY_ROLE)
}

func prepareExpectedUserModifyResponse() fivetran.UserModifyResponse {
	var verifyFlag = false
	return fivetran.UserModifyResponse{
		Code:    EXPECTED_USER_MODIFY_RESPONSE_CODE,
		Message: "User has been invited to the account",
		Data: fivetran.UserModifyData{
			ID:         EXPECTED_USER_MODIFY_USER_ID,
			Email:      EXPECTED_USER_MODIFY_EMAIL,
			GivenName:  EXPECTED_USER_MODIFY_GIVEN_NAME,
			FamilyName: EXPECTED_USER_MODIFY_FAMILY_NAME,
			Verified:   &verifyFlag,
			Invited:    &verifyFlag,
			Picture:    EXPECTED_USER_MODIFY_PICTURE,
			Phone:      EXPECTED_USER_MODIFY_PHONE,
			Role:       EXPECTED_USER_MODIFY_ROLE,
		},
	}
}

func assertUserModifyResponse(t *testing.T, actual, expected fivetran.UserModifyResponse) {
	assertEqual(t, actual.Code, expected.Code)
	assertEqual(t, actual.Message, expected.Message)
	assertEqual(t, actual.Data.ID, expected.Data.ID)
	assertEqual(t, actual.Data.Email, expected.Data.Email)
	assertEqual(t, actual.Data.GivenName, expected.Data.GivenName)
	assertEqual(t, actual.Data.FamilyName, expected.Data.FamilyName)
	assertEqual(t, actual.Data.Verified, expected.Data.Verified)
	assertEqual(t, actual.Data.Picture, expected.Data.Picture)
	assertEqual(t, actual.Data.Phone, expected.Data.Phone)
	assertEqual(t, actual.Data.Role, expected.Data.Role)
}
