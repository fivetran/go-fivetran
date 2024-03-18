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

const DESTINATION_DETAILS_SERVICE = "snowflake"
const DESTINATION_DETAILS_ID = "decent_dropsy"
const DESTINATION_DETAILS_REGION = "GCP_US_EAST4"
const DESTINATION_DETAILS_TIME_ZONE = "-5"
const DESTINATION_DETAILS_SETUP_STATUS = "connected"
const DESTINATION_DETAILS_DAYLIGHT = true
const DESTINATION_DETAILS_HOST = "your-account.snowflakecomputing.com"
const DESTINATION_DETAILS_PORT = "443"
const DESTINATION_DETAILS_DATABASE = "fivetran"
const DESTINATION_DETAILS_AUTH = "PASSWORD"
const DESTINATION_DETAILS_USER = "fivetran_user"
const DESTINATION_DETAILS_MASKED_PASSWORD = "******"

func TestDestinationDetailsService(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/destinations/"+ID).ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareDestinationDetailsResponse())
			return response, nil
		})

	service := ftClient.NewDestinationDetails().DestinationID(ID)

	// act
	response, err := service.Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	// assert
	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)

	assertDestinationDetailsResponse(t, response)
}

func TestDestinationDetailsCustomService(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodGet, "/v1/destinations/"+ID).ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareDestinationDetailsResponse())
			return response, nil
		})

	service := ftClient.NewDestinationDetails().DestinationID(ID)

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

	assertDestinationDetailsCustomResponse(t, response)
}

func prepareDestinationDetailsResponse() string {
	return fmt.Sprintf(`{
		"code": "Success",
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
				"port": "%s",
				"database": "%s",
				"auth": "%s",
				"user": "%s",
				"password": "%s"
			}
		}
	}`,
		DESTINATION_DETAILS_ID,
		DESTINATION_DETAILS_ID,
		DESTINATION_DETAILS_SERVICE,
		DESTINATION_DETAILS_REGION,
		DESTINATION_DETAILS_DAYLIGHT,
		DESTINATION_DETAILS_TIME_ZONE,
		DESTINATION_DETAILS_SETUP_STATUS,
		DESTINATION_DETAILS_HOST,
		DESTINATION_DETAILS_PORT,
		DESTINATION_DETAILS_DATABASE,
		DESTINATION_DETAILS_AUTH,
		DESTINATION_DETAILS_USER,
		DESTINATION_DETAILS_MASKED_PASSWORD)
}

func assertDestinationDetailsResponse(t *testing.T, response destinations.DestinationDetailsResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.ID, DESTINATION_DETAILS_ID)
	testutils.AssertEqual(t, response.Data.GroupID, DESTINATION_DETAILS_ID)
	testutils.AssertEqual(t, response.Data.Service, DESTINATION_DETAILS_SERVICE)
	testutils.AssertEqual(t, response.Data.Region, DESTINATION_DETAILS_REGION)
	testutils.AssertEqual(t, response.Data.DaylightSavingTimeEnabled, DESTINATION_DETAILS_DAYLIGHT)
	testutils.AssertEqual(t, response.Data.TimeZoneOffset, DESTINATION_DETAILS_TIME_ZONE)
	testutils.AssertEqual(t, response.Data.SetupStatus, DESTINATION_DETAILS_SETUP_STATUS)
	testutils.AssertEqual(t, response.Data.Config.Host, DESTINATION_DETAILS_HOST)
	testutils.AssertEqual(t, response.Data.Config.Port, DESTINATION_DETAILS_PORT)
	testutils.AssertEqual(t, response.Data.Config.Database, DESTINATION_DETAILS_DATABASE)
	testutils.AssertEqual(t, response.Data.Config.Auth, DESTINATION_DETAILS_AUTH)
	testutils.AssertEqual(t, response.Data.Config.User, DESTINATION_DETAILS_USER)
	testutils.AssertEqual(t, response.Data.Config.Password, DESTINATION_DETAILS_MASKED_PASSWORD)
}

func assertDestinationDetailsCustomResponse(t *testing.T, response destinations.DestinationDetailsCustomResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.ID, DESTINATION_DETAILS_ID)
	testutils.AssertEqual(t, response.Data.GroupID, DESTINATION_DETAILS_ID)
	testutils.AssertEqual(t, response.Data.Service, DESTINATION_DETAILS_SERVICE)
	testutils.AssertEqual(t, response.Data.DaylightSavingTimeEnabled, DESTINATION_DETAILS_DAYLIGHT)
	testutils.AssertEqual(t, response.Data.Region, DESTINATION_DETAILS_REGION)
	testutils.AssertEqual(t, response.Data.TimeZoneOffset, DESTINATION_DETAILS_TIME_ZONE)
	testutils.AssertEqual(t, response.Data.SetupStatus, DESTINATION_DETAILS_SETUP_STATUS)
	testutils.AssertEqual(t, response.Data.Config["host"], DESTINATION_DETAILS_HOST)
	testutils.AssertEqual(t, response.Data.Config["port"], DESTINATION_DETAILS_PORT)
	testutils.AssertEqual(t, response.Data.Config["database"], DESTINATION_DETAILS_DATABASE)
	testutils.AssertEqual(t, response.Data.Config["auth"], DESTINATION_DETAILS_AUTH)
	testutils.AssertEqual(t, response.Data.Config["user"], DESTINATION_DETAILS_USER)
	testutils.AssertEqual(t, response.Data.Config["password"], DESTINATION_DETAILS_MASKED_PASSWORD)
}
