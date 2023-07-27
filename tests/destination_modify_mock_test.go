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
	DESTINATION_MODIFY_SERVICE            = "snowflake"
	DESTINATION_MODIFY_ID                 = "decent_dropsy"
	DESTINATION_MODIFY_REGION             = "GCP_AUSTRALIA_SOUTHEAST1"
	DESTINATION_MODIFY_TIME_ZONE_OFFSET   = "+10"
	DESTINATION_MODIFY_SETUP_STATUS       = "connected"
	DESTINATION_MODIFY_TEST_TITLE_HOST    = "Host Connection"
	DESTINATION_MODIFY_TEST_TITLE_DB      = "Database Connection"
	DESTINATION_MODIFY_TEST_TITLE_PERM    = "Permission Test"
	DESTINATION_MODIFY_TEST_STATUS_PASSED = "PASSED"
	DESTINATION_MODIFY_HOST               = "your-account.snowflakecomputing.com"
	DESTINATION_MODIFY_PORT               = "1433"
	DESTINATION_MODIFY_DATABASE           = "fivetran"
	DESTINATION_MODIFY_AUTH               = "PASSWORD"
	DESTINATION_MODIFY_USER               = "fivetran_user"
	DESTINATION_MODIFY_PASSWORD           = "******"
	DESTINATION_MODIFY_MASKED             = "******"
)

func TestDestinationModifyService(t *testing.T) {
	// arrange
	ftClient, mockClient := CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/destinations/decent_dropsy").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareDestinationModifyResponse())
			return response, nil
		})

	destinationConfig := fivetran.NewDestinationConfig()
	destinationConfig = destinationConfig.
		Host(DESTINATION_MODIFY_HOST).
		Port(1433).
		Database(DESTINATION_MODIFY_DATABASE).
		Auth(DESTINATION_MODIFY_AUTH).
		User(DESTINATION_MODIFY_USER).
		Password(DESTINATION_MODIFY_PASSWORD)

	service := ftClient.NewDestinationModify().
		DestinationID(DESTINATION_MODIFY_ID).
		Region(DESTINATION_MODIFY_REGION).
		TimeZoneOffset(DESTINATION_MODIFY_TIME_ZONE_OFFSET).
		Config(destinationConfig).
		TrustCertificates(TRUST_CERTIFICATES).
		TrustFingerprints(TRUST_FINGERPRINTS).
		RunSetupTests(RUN_SETUP_TESTS)

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

	assertDestinationModifyResponse(t, response)
}

func prepareDestinationModifyResponse() string {
	return fmt.Sprintf(`{
		"code": "Success",
		"message": "Destination has been updated",
		"data": {
			"id": "%s",
			"group_id": "%s",
			"service": "%s",
			"region": "%s",
			"time_zone_offset": "%s",
			"setup_status": "%s",
			"setup_tests": [
				{
					"title": "%s",
					"status": "%s",
					"message": ""
				},
				{
					"title": "%s",
					"status": "%s",
					"message": ""
				},
				{
					"title": "%s",
					"status": "%s",
					"message": ""
				}
			],
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
		DESTINATION_MODIFY_ID,
		DESTINATION_MODIFY_ID,
		DESTINATION_MODIFY_SERVICE,
		DESTINATION_MODIFY_REGION,
		DESTINATION_MODIFY_TIME_ZONE_OFFSET,
		DESTINATION_MODIFY_SETUP_STATUS,
		DESTINATION_MODIFY_TEST_TITLE_HOST,
		DESTINATION_MODIFY_TEST_STATUS_PASSED,
		DESTINATION_MODIFY_TEST_TITLE_DB,
		DESTINATION_MODIFY_TEST_STATUS_PASSED,
		DESTINATION_MODIFY_TEST_TITLE_PERM,
		DESTINATION_MODIFY_TEST_STATUS_PASSED,
		DESTINATION_MODIFY_HOST,
		DESTINATION_MODIFY_PORT,
		DESTINATION_MODIFY_DATABASE,
		DESTINATION_MODIFY_AUTH,
		DESTINATION_MODIFY_USER,
		DESTINATION_MODIFY_MASKED)
}

func assertDestinationModifyResponse(t *testing.T, response fivetran.DestinationModifyResponse) {
	assertEqual(t, response.Code, "Success")
	assertEqual(t, response.Message, "Destination has been updated")
	assertEqual(t, response.Data.ID, DESTINATION_MODIFY_ID)
	assertEqual(t, response.Data.GroupID, DESTINATION_MODIFY_ID)
	assertEqual(t, response.Data.Service, DESTINATION_MODIFY_SERVICE)
	assertEqual(t, response.Data.Region, DESTINATION_MODIFY_REGION)
	assertEqual(t, response.Data.TimeZoneOffset, DESTINATION_MODIFY_TIME_ZONE_OFFSET)
	assertEqual(t, response.Data.SetupStatus, DESTINATION_MODIFY_SETUP_STATUS)

	// assert setup tests
	assertEqual(t, len(response.Data.SetupTests), 3)
	assertEqual(t, response.Data.SetupTests[0].Title, DESTINATION_MODIFY_TEST_TITLE_HOST)
	assertEqual(t, response.Data.SetupTests[0].Status, DESTINATION_MODIFY_TEST_STATUS_PASSED)
	assertEqual(t, response.Data.SetupTests[1].Title, DESTINATION_MODIFY_TEST_TITLE_DB)
	assertEqual(t, response.Data.SetupTests[1].Status, DESTINATION_MODIFY_TEST_STATUS_PASSED)
	assertEqual(t, response.Data.SetupTests[2].Title, DESTINATION_MODIFY_TEST_TITLE_PERM)
	assertEqual(t, response.Data.SetupTests[2].Status, DESTINATION_MODIFY_TEST_STATUS_PASSED)

	// assert config
	assertEqual(t, response.Data.Config.Host, DESTINATION_MODIFY_HOST)
	assertEqual(t, response.Data.Config.Port, DESTINATION_MODIFY_PORT)
	assertEqual(t, response.Data.Config.Database, DESTINATION_MODIFY_DATABASE)
	assertEqual(t, response.Data.Config.Auth, DESTINATION_MODIFY_AUTH)
	assertEqual(t, response.Data.Config.User, DESTINATION_MODIFY_USER)
	assertEqual(t, response.Data.Config.Password, DESTINATION_MODIFY_MASKED)
}
