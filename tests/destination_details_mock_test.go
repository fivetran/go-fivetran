package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/tests/mock"
)

const DESTINATION_DETAILS_SERVICE = "snowflake"
const DESTINATION_DETAILS_ID = "decent_dropsy"
const DESTINATION_DETAILS_REGION = "GCP_US_EAST4"
const DESTINATION_DETAILS_TIME_ZONE = "-5"
const DESTINATION_DETAILS_SETUP_STATUS = "connected"
const DESTINATION_DETAILS_HOST = "your-account.snowflakecomputing.com"
const DESTINATION_DETAILS_PORT = 443
const DESTINATION_DETAILS_DATABASE = "fivetran"
const DESTINATION_DETAILS_AUTH = "PASSWORD"
const DESTINATION_DETAILS_USER = "fivetran_user"
const DESTINATION_DETAILS_MASKED_PASSWORD = "******"

func TestDestinationDetailsService(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
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
	assertEqual(t, len(interactions), 1)
	assertEqual(t, interactions[0].Handler, handler)
	assertEqual(t, handler.Interactions, 1)

	assertDestinationDetailsResponse(t, response)
}

func prepareDestinationDetailsResponse() string {
	var responce = fmt.Sprintf(`{
		"code": "Success",
		"data": {
			"id": "%s",
			"group_id": "%s",
			"service": "%s",
			"region": "%s",
			"time_zone_offset": "%s",
			"setup_status": "%s",
			"config": {
				"host": "%s",
				"port": %d,
				"database": "%s",
				"auth": "%s",
				"user": "%s",
				"password": "%s"
			}
		}
	}`, DESTINATION_DETAILS_ID, DESTINATION_DETAILS_ID, DESTINATION_DETAILS_SERVICE, DESTINATION_DETAILS_REGION, DESTINATION_DETAILS_TIME_ZONE, DESTINATION_DETAILS_SETUP_STATUS, DESTINATION_DETAILS_HOST, DESTINATION_DETAILS_PORT, DESTINATION_DETAILS_DATABASE, DESTINATION_DETAILS_AUTH, DESTINATION_DETAILS_USER, DESTINATION_DETAILS_MASKED_PASSWORD)
	return responce
}

func assertDestinationDetailsResponse(t *testing.T, response fivetran.DestinationDetailsResponse) {
	assertEqual(t, response.Code, "Success")
	assertEqual(t, response.Data.ID, DESTINATION_DETAILS_ID)
	assertEqual(t, response.Data.GroupID, DESTINATION_DETAILS_ID)
	assertEqual(t, response.Data.Service, DESTINATION_DETAILS_SERVICE)
	assertEqual(t, response.Data.Region, DESTINATION_DETAILS_REGION)
	assertEqual(t, response.Data.TimeZoneOffset, DESTINATION_DETAILS_TIME_ZONE)
	assertEqual(t, response.Data.SetupStatus, DESTINATION_DETAILS_SETUP_STATUS)
	assertEqual(t, response.Data.Config.Host, DESTINATION_DETAILS_HOST)
	assertEqual(t, response.Data.Config.Port, DESTINATION_DETAILS_PORT)
	assertEqual(t, response.Data.Config.Database, DESTINATION_DETAILS_DATABASE)
	assertEqual(t, response.Data.Config.Auth, DESTINATION_DETAILS_AUTH)
	assertEqual(t, response.Data.Config.User, DESTINATION_DETAILS_USER)
	assertEqual(t, response.Data.Config.Password, DESTINATION_DETAILS_MASKED_PASSWORD)
}
