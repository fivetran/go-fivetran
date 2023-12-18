package users_test

import (
    "context"
    "fmt"
    "net/http"
    "testing"
    "time"

	"github.com/fivetran/go-fivetran/common"
	"github.com/fivetran/go-fivetran/users"
    
    "github.com/fivetran/go-fivetran/tests/mock"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

const (
	EXPECTED_USER_RESPONSE_CODE = "Success"
	EXPECTED_USER_ID            = "nozzle_eat"
	EXPECTED_USER_EMAIL         = "john@mycompany.com"
	EXPECTED_USER_GIVEN_NAME    = "John"
	EXPECTED_USER_FAMILY_NAME   = "White"
	EXPECTED_USER_VERIFIED      = true
	EXPECTED_USER_INVITED       = true
	EXPECTED_USER_PICTURE       = "null"
	EXPECTED_USER_PHONE         = "null"
	EXPECTED_USER_LOGGED_IN_AT  = "2019-01-03T08:44:45.369Z"
	EXPECTED_USER_CREATED_AT    = "2018-01-15T11:00:27.329220Z"
	EXPECTED_USER_ROLE          = "Account Reviewer"
)

func TestUserDetailsServiceDo(t *testing.T) {
	ftClient, mockClient := testutils.CreateTestClient()

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
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
}

func TestUserDetailsServiceDoMissingUserID(t *testing.T) {
	ftClient, _ := testutils.CreateTestClient()

	service := ftClient.NewUserDetails()

	_, err := service.Do(context.Background())

	expectedError := fmt.Errorf("missing required userID")
	testutils.AssertEqual(t, err, expectedError)
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

func prepareExpectedUserDetailsResponse() users.UserDetailsResponse {
	var flag = true
	return users.UserDetailsResponse{
		CommonResponse: common.CommonResponse{
			Code:    EXPECTED_USER_RESPONSE_CODE,
			Message: "",
		},
		Data: users.UserDetailsData{
			ID:         EXPECTED_USER_ID,
			Email:      EXPECTED_USER_EMAIL,
			GivenName:  EXPECTED_USER_GIVEN_NAME,
			FamilyName: EXPECTED_USER_FAMILY_NAME,
			Verified:   &flag,
			Invited:    &flag,
			Picture:    EXPECTED_USER_PICTURE,
			Phone:      EXPECTED_USER_PHONE,
			LoggedInAt: parseTime(EXPECTED_USER_LOGGED_IN_AT),
			CreatedAt:  parseTime(EXPECTED_USER_CREATED_AT),
			Role:       EXPECTED_USER_ROLE,
		},
	}
}

func assertUserDetailsResponse(t *testing.T,
	actual users.UserDetailsResponse,
	expected users.UserDetailsResponse,
	code string,
	massage string) {
	testutils.AssertEqual(t, actual.Code, code)
	testutils.AssertEqual(t, actual.Message, massage)
	testutils.AssertEqual(t, actual.Data.ID, expected.Data.ID)
	testutils.AssertEqual(t, actual.Data.Email, expected.Data.Email)
	testutils.AssertEqual(t, actual.Data.GivenName, expected.Data.GivenName)
	testutils.AssertEqual(t, actual.Data.FamilyName, expected.Data.FamilyName)
	testutils.AssertEqual(t, actual.Data.Verified, expected.Data.Verified)
	testutils.AssertEqual(t, actual.Data.Invited, expected.Data.Invited)
	testutils.AssertEqual(t, actual.Data.Picture, expected.Data.Picture)
	testutils.AssertEqual(t, actual.Data.Phone, expected.Data.Phone)
	testutils.AssertEqual(t, actual.Data.LoggedInAt, expected.Data.LoggedInAt)
	testutils.AssertEqual(t, actual.Data.CreatedAt, expected.Data.CreatedAt)
	testutils.AssertEqual(t, actual.Data.Role, expected.Data.Role)
}

func parseTime(timeStr string) time.Time {
	parsedTime, _ := time.Parse(time.RFC3339Nano, timeStr)
	return parsedTime
}
