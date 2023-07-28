package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

const (
	EXPECTED_USER_RESPONSE_CODE = "Success"
	EXPECTED_USER_ID            = "nozzle_eat"
	EXPECTED_USER_EMAIL         = "john@mycompany.com"
	EXPECTED_USER_GIVEN_NAME    = "John"
	EXPECTED_USER_FAMILY_NAME   = "White"
	EXPECTED_USER_VERIFIED      = true
	EXPECTED_USER_INVITED       = false
	EXPECTED_USER_PICTURE       = "null"
	EXPECTED_USER_PHONE         = "null"
	EXPECTED_USER_LOGGED_IN_AT  = "2019-01-03T08:44:45.369Z"
	EXPECTED_USER_CREATED_AT    = "2018-01-15T11:00:27.329220Z"
	EXPECTED_USER_ROLE          = "Account Reviewer"
)

func TestUserDetailsServiceDo(t *testing.T) {
	ftClient, mockClient := CreateTestClient()

	handler := mockClient.When(http.MethodGet, "/v1/users/user_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			responseData := prepareUserDetailsResponse()
			response := mock.NewResponse(req, http.StatusOK, responseData)
			return response, nil
		},
	)

	service := ftClient.NewUserDetails().UserID("user_id")

	response, err := service.Do(context.Background())

	if err != nil {
		t.Error(err)
	}

	expectedResponse := prepareExpectedUserDetailsResponse()
	assertUserDetailsResponse(t, response, expectedResponse, "Success", "")

	interactions := mockClient.Interactions()
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)
}

func TestUserDetailsServiceDoMissingUserID(t *testing.T) {
	ftClient, _ := CreateTestClient()

	service := ftClient.NewUserDetails()

	_, err := service.Do(context.Background())

	expectedError := fmt.Errorf("missing required UserId")
	assertEqual(t, err, expectedError)
}

func prepareUserDetailsResponse() string {
	return fmt.Sprintf(`{
		"code": "%s",
		"data": {
			"id": "%s",
			"email": "%s",
			"given_name": "%s",
			"family_name": "%s",
			"verified": %t,
			"invited": %t,
			"picture": "%s",
			"phone": "%s",
			"role": "%s",
			"logged_in_at": "%s",
			"created_at": "%s",
			"active": %t
		}
	}`,
		EXPECTED_USER_RESPONSE_CODE,
		EXPECTED_USER_ID,
		EXPECTED_USER_EMAIL,
		EXPECTED_USER_GIVEN_NAME,
		EXPECTED_USER_FAMILY_NAME,
		EXPECTED_USER_VERIFIED,
		EXPECTED_USER_INVITED,
		EXPECTED_USER_PICTURE,
		EXPECTED_USER_PHONE,
		EXPECTED_USER_ROLE,
		EXPECTED_USER_LOGGED_IN_AT,
		EXPECTED_USER_CREATED_AT,
		true)
}

func prepareExpectedUserDetailsResponse() fivetran.UserDetailsResponse {
	return fivetran.UserDetailsResponse{
		Code:    EXPECTED_USER_RESPONSE_CODE,
		Message: "",
		Data: fivetran.UserDetailsData{
			ID:         EXPECTED_USER_ID,
			Email:      EXPECTED_USER_EMAIL,
			GivenName:  EXPECTED_USER_GIVEN_NAME,
			FamilyName: EXPECTED_USER_FAMILY_NAME,
			Verified:   boolPtr(EXPECTED_USER_VERIFIED),
			Invited:    boolPtr(EXPECTED_USER_INVITED),
			Picture:    EXPECTED_USER_PICTURE,
			Phone:      EXPECTED_USER_PHONE,
			LoggedInAt: parseTime(EXPECTED_USER_LOGGED_IN_AT),
			CreatedAt:  parseTime(EXPECTED_USER_CREATED_AT),
			Role:       EXPECTED_USER_ROLE,
		},
	}
}

func assertUserDetailsResponse(t *testing.T,
	actual fivetran.UserDetailsResponse,
	expected fivetran.UserDetailsResponse,
	code string,
	massage string) {
	assertEqual(t, actual.Code, code)
	assertEqual(t, actual.Message, massage)
	assertEqual(t, actual.Data.ID, expected.Data.ID)
	assertEqual(t, actual.Data.Email, expected.Data.Email)
	assertEqual(t, actual.Data.GivenName, expected.Data.GivenName)
	assertEqual(t, actual.Data.FamilyName, expected.Data.FamilyName)
	assertEqual(t, actual.Data.Verified, expected.Data.Verified)
	assertEqual(t, actual.Data.Invited, expected.Data.Invited)
	assertEqual(t, actual.Data.Picture, expected.Data.Picture)
	assertEqual(t, actual.Data.Phone, expected.Data.Phone)
	assertEqual(t, actual.Data.LoggedInAt, expected.Data.LoggedInAt)
	assertEqual(t, actual.Data.CreatedAt, expected.Data.CreatedAt)
	assertEqual(t, actual.Data.Role, expected.Data.Role)
}

func boolPtr(value bool) *bool {
	return &value
}

func parseTime(timeStr string) time.Time {
	parsedTime, _ := time.Parse(time.RFC3339Nano, timeStr)
	return parsedTime
}
