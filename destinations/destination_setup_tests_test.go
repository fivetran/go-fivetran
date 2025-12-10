package destinations_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/destinations"
	"github.com/fivetran/go-fivetran/tests/mock"
	testutils "github.com/fivetran/go-fivetran/test_utils"
)

const SETUP_TESTS_ID = "decent_dropsy"
const SETUP_TESTS_SERVICE = "snowflake"
const SETUP_TESTS_REGION = "GCP_US_EAST4"
const SETUP_TESTS_TIME_ZONE = "-5"
const SETUP_TESTS_SETUP_STATUS = "connected"
const SETUP_TESTS_DAYLIGHT = true
const SETUP_TESTS_HOST = "your-account.snowflakecomputing.com"
const SETUP_TESTS_PORT = 443
const SETUP_TESTS_DATABASE = "fivetran"
const SETUP_TESTS_AUTH = "PASSWORD"
const SETUP_TESTS_USER = "fivetran_user"
const SETUP_TESTS_MASKED_PASSWORD = "******"
const SETUP_TESTS_TITLE = "Connection Test"
const SETUP_TESTS_STATUS = "PASSED"
const SETUP_TESTS_MESSAGE = "Successfully connected"

func TestDestinationSetupTestsCustomService(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/destinations/"+SETUP_TESTS_ID+"/test").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareSetupTestsResponseWithIntPort())
			return response, nil
		})

	service := ftClient.NewDestinationSetupTests().
		DestinationID(SETUP_TESTS_ID).
		TrustCertificates(true).
		TrustFingerprints(false)

	// act
	response, err := service.DoCustom(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	// assert
	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)

	assertSetupTestsCustomResponse(t, response)
}

func prepareSetupTestsResponseWithIntPort() string {
	return fmt.Sprintf(`{
		"code": "Success",
		"message": "Setup tests have been completed",
		"data": {
			"id": "%s",
			"group_id": "%s",
			"service": "%s",
			"region": "%s",
			"daylight_saving_time_enabled": %v,
			"time_zone_offset": "%s",
			"setup_status": "%s",
			"config": {
				"host": "%s",
				"port": %d,
				"database": "%s",
				"auth": "%s",
				"user": "%s",
				"password": "%s"
			},
			"setup_tests": [
				{
					"title": "%s",
					"status": "%s",
					"message": "%s"
				}
			]
		}
	}`,
		SETUP_TESTS_ID,
		SETUP_TESTS_ID,
		SETUP_TESTS_SERVICE,
		SETUP_TESTS_REGION,
		SETUP_TESTS_DAYLIGHT,
		SETUP_TESTS_TIME_ZONE,
		SETUP_TESTS_SETUP_STATUS,
		SETUP_TESTS_HOST,
		SETUP_TESTS_PORT,
		SETUP_TESTS_DATABASE,
		SETUP_TESTS_AUTH,
		SETUP_TESTS_USER,
		SETUP_TESTS_MASKED_PASSWORD,
		SETUP_TESTS_TITLE,
		SETUP_TESTS_STATUS,
		SETUP_TESTS_MESSAGE)
}

func assertSetupTestsCustomResponse(t *testing.T, response destinations.DestinationDetailsWithSetupTestsCustomResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.ID, SETUP_TESTS_ID)
	testutils.AssertEqual(t, response.Data.GroupID, SETUP_TESTS_ID)
	testutils.AssertEqual(t, response.Data.Service, SETUP_TESTS_SERVICE)
	testutils.AssertEqual(t, response.Data.DaylightSavingTimeEnabled, SETUP_TESTS_DAYLIGHT)
	testutils.AssertEqual(t, response.Data.Region, SETUP_TESTS_REGION)
	testutils.AssertEqual(t, response.Data.TimeZoneOffset, SETUP_TESTS_TIME_ZONE)
	testutils.AssertEqual(t, response.Data.SetupStatus, SETUP_TESTS_SETUP_STATUS)
	testutils.AssertEqual(t, response.Data.Config["host"], SETUP_TESTS_HOST)
	// Port is returned as float64 when unmarshaled from JSON number
	testutils.AssertEqual(t, response.Data.Config["port"], float64(SETUP_TESTS_PORT))
	testutils.AssertEqual(t, response.Data.Config["database"], SETUP_TESTS_DATABASE)
	testutils.AssertEqual(t, response.Data.Config["auth"], SETUP_TESTS_AUTH)
	testutils.AssertEqual(t, response.Data.Config["user"], SETUP_TESTS_USER)
	testutils.AssertEqual(t, response.Data.Config["password"], SETUP_TESTS_MASKED_PASSWORD)
	testutils.AssertEqual(t, len(response.Data.SetupTests), 1)
	testutils.AssertEqual(t, response.Data.SetupTests[0].Title, SETUP_TESTS_TITLE)
	testutils.AssertEqual(t, response.Data.SetupTests[0].Status, SETUP_TESTS_STATUS)
	testutils.AssertEqual(t, response.Data.SetupTests[0].Message, SETUP_TESTS_MESSAGE)
}
