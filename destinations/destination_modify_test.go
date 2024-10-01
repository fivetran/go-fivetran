package destinations_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/fivetran/go-fivetran/destinations"

	"github.com/fivetran/go-fivetran/tests/mock"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

const (
	DESTINATION_MODIFY_SERVICE            		= "snowflake"
	DESTINATION_MODIFY_ID                 		= "decent_dropsy"
	DESTINATION_MODIFY_REGION             		= "GCP_AUSTRALIA_SOUTHEAST1"
	DESTINATION_MODIFY_TIME_ZONE_OFFSET   		= "+10"
	DESTINATION_MODIFY_DAYLIGHT 		  		= true
    DESTINATION_MODIFY_HYBRIDDEPLOYMENTAGENTID  = "hybrid_deployment_agent_id"
    DESTINATION_MODIFY_PRIVATELINKID            = "private_link_id"
    DESTINATION_MODIFY_NETWORKINGMETHOD         = "Direct"
	DESTINATION_MODIFY_SETUP_STATUS      		= "connected"
	DESTINATION_MODIFY_TEST_TITLE_HOST    		= "Host Connection"
	DESTINATION_MODIFY_TEST_TITLE_DB      		= "Database Connection"
	DESTINATION_MODIFY_TEST_TITLE_PERM    		= "Permission Test"
	DESTINATION_MODIFY_TEST_STATUS_PASSED 		= "PASSED"
	DESTINATION_MODIFY_HOST              		= "your-account.snowflakecomputing.com"
	DESTINATION_MODIFY_PORT               		= "1433"
	DESTINATION_MODIFY_DATABASE          		= "fivetran"
	DESTINATION_MODIFY_AUTH               		= "PASSWORD"
	DESTINATION_MODIFY_USER               		= "fivetran_user"
	DESTINATION_MODIFY_PASSWORD           		= "******"
	DESTINATION_MODIFY_MASKED             		= "******"
)

func TestDestinationModifyService(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/destinations/decent_dropsy").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertModifyRequest(t, body)
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
		DaylightSavingTimeEnabled(DESTINATION_MODIFY_DAYLIGHT).
        HybridDeploymentAgentId(DESTINATION_MODIFY_HYBRIDDEPLOYMENTAGENTID).
        PrivateLinkId(DESTINATION_MODIFY_PRIVATELINKID).
        NetworkingMethod(DESTINATION_MODIFY_NETWORKINGMETHOD).
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
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)

	assertDestinationModifyResponse(t, response)
}

func TestDestinationModifyCustomService(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/destinations/decent_dropsy").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertModifyRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, prepareDestinationModifyResponse())
			return response, nil
		})

	service := ftClient.NewDestinationModify().
		DestinationID(DESTINATION_MODIFY_ID).
		Region(DESTINATION_MODIFY_REGION).
		DaylightSavingTimeEnabled(DESTINATION_MODIFY_DAYLIGHT).
		HybridDeploymentAgentId(DESTINATION_MODIFY_HYBRIDDEPLOYMENTAGENTID).
        PrivateLinkId(DESTINATION_MODIFY_PRIVATELINKID).
        NetworkingMethod(DESTINATION_MODIFY_NETWORKINGMETHOD).
		TimeZoneOffset(DESTINATION_MODIFY_TIME_ZONE_OFFSET).
		ConfigCustom(&map[string]interface{}{
			"host":     DESTINATION_MODIFY_HOST,
			"port":     1433,
			"database": DESTINATION_MODIFY_DATABASE,
			"auth":     DESTINATION_MODIFY_AUTH,
			"user":     DESTINATION_MODIFY_USER,
			"password": DESTINATION_MODIFY_PASSWORD,
		}).
		TrustCertificates(TRUST_CERTIFICATES).
		TrustFingerprints(TRUST_FINGERPRINTS).
		RunSetupTests(RUN_SETUP_TESTS)

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

	assertDestinationModifyCustomResponse(t, response)
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
			"daylight_saving_time_enabled": %v,
			"hybrid_deployment_agent_id": "%v",
            "private_link_id": "%v",
            "networking_method": "%v",
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
		DESTINATION_MODIFY_DAYLIGHT,
		DESTINATION_MODIFY_HYBRIDDEPLOYMENTAGENTID,
        DESTINATION_MODIFY_PRIVATELINKID,
        DESTINATION_MODIFY_NETWORKINGMETHOD,
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

func assertDestinationModifyResponse(t *testing.T, response destinations.DestinationDetailsWithSetupTestsResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Message, "Destination has been updated")
	testutils.AssertEqual(t, response.Data.ID, DESTINATION_MODIFY_ID)
	testutils.AssertEqual(t, response.Data.GroupID, DESTINATION_MODIFY_ID)
	testutils.AssertEqual(t, response.Data.Service, DESTINATION_MODIFY_SERVICE)
	testutils.AssertEqual(t, response.Data.DaylightSavingTimeEnabled, DESTINATION_MODIFY_DAYLIGHT)
	testutils.AssertEqual(t, response.Data.HybridDeploymentAgentId, DESTINATION_MODIFY_HYBRIDDEPLOYMENTAGENTID)
	testutils.AssertEqual(t, response.Data.PrivateLinkId, DESTINATION_MODIFY_PRIVATELINKID)
	testutils.AssertEqual(t, response.Data.NetworkingMethod, DESTINATION_MODIFY_NETWORKINGMETHOD)
	testutils.AssertEqual(t, response.Data.Region, DESTINATION_MODIFY_REGION)
	testutils.AssertEqual(t, response.Data.TimeZoneOffset, DESTINATION_MODIFY_TIME_ZONE_OFFSET)
	testutils.AssertEqual(t, response.Data.SetupStatus, DESTINATION_MODIFY_SETUP_STATUS)

	// assert setup tests
	testutils.AssertEqual(t, len(response.Data.SetupTests), 3)
	testutils.AssertEqual(t, response.Data.SetupTests[0].Title, DESTINATION_MODIFY_TEST_TITLE_HOST)
	testutils.AssertEqual(t, response.Data.SetupTests[0].Status, DESTINATION_MODIFY_TEST_STATUS_PASSED)
	testutils.AssertEqual(t, response.Data.SetupTests[1].Title, DESTINATION_MODIFY_TEST_TITLE_DB)
	testutils.AssertEqual(t, response.Data.SetupTests[1].Status, DESTINATION_MODIFY_TEST_STATUS_PASSED)
	testutils.AssertEqual(t, response.Data.SetupTests[2].Title, DESTINATION_MODIFY_TEST_TITLE_PERM)
	testutils.AssertEqual(t, response.Data.SetupTests[2].Status, DESTINATION_MODIFY_TEST_STATUS_PASSED)

	// assert config
	testutils.AssertEqual(t, response.Data.Config.Host, DESTINATION_MODIFY_HOST)
	testutils.AssertEqual(t, response.Data.Config.Port, DESTINATION_MODIFY_PORT)
	testutils.AssertEqual(t, response.Data.Config.Database, DESTINATION_MODIFY_DATABASE)
	testutils.AssertEqual(t, response.Data.Config.Auth, DESTINATION_MODIFY_AUTH)
	testutils.AssertEqual(t, response.Data.Config.User, DESTINATION_MODIFY_USER)
	testutils.AssertEqual(t, response.Data.Config.Password, DESTINATION_MODIFY_MASKED)
}

func assertDestinationModifyCustomResponse(t *testing.T, response destinations.DestinationDetailsWithSetupTestsCustomResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Message, "Destination has been updated")
	testutils.AssertEqual(t, response.Data.ID, DESTINATION_MODIFY_ID)
	testutils.AssertEqual(t, response.Data.GroupID, DESTINATION_MODIFY_ID)
	testutils.AssertEqual(t, response.Data.Service, DESTINATION_MODIFY_SERVICE)
	testutils.AssertEqual(t, response.Data.DaylightSavingTimeEnabled, DESTINATION_MODIFY_DAYLIGHT)
	testutils.AssertEqual(t, response.Data.HybridDeploymentAgentId, DESTINATION_MODIFY_HYBRIDDEPLOYMENTAGENTID)
	testutils.AssertEqual(t, response.Data.PrivateLinkId, DESTINATION_MODIFY_PRIVATELINKID)
	testutils.AssertEqual(t, response.Data.NetworkingMethod, DESTINATION_MODIFY_NETWORKINGMETHOD)
	testutils.AssertEqual(t, response.Data.Region, DESTINATION_MODIFY_REGION)
	testutils.AssertEqual(t, response.Data.TimeZoneOffset, DESTINATION_MODIFY_TIME_ZONE_OFFSET)
	testutils.AssertEqual(t, response.Data.SetupStatus, DESTINATION_MODIFY_SETUP_STATUS)

	// assert setup tests
	testutils.AssertEqual(t, len(response.Data.SetupTests), 3)
	testutils.AssertEqual(t, response.Data.SetupTests[0].Title, DESTINATION_MODIFY_TEST_TITLE_HOST)
	testutils.AssertEqual(t, response.Data.SetupTests[0].Status, DESTINATION_MODIFY_TEST_STATUS_PASSED)
	testutils.AssertEqual(t, response.Data.SetupTests[1].Title, DESTINATION_MODIFY_TEST_TITLE_DB)
	testutils.AssertEqual(t, response.Data.SetupTests[1].Status, DESTINATION_MODIFY_TEST_STATUS_PASSED)
	testutils.AssertEqual(t, response.Data.SetupTests[2].Title, DESTINATION_MODIFY_TEST_TITLE_PERM)
	testutils.AssertEqual(t, response.Data.SetupTests[2].Status, DESTINATION_MODIFY_TEST_STATUS_PASSED)

	// assert config
	testutils.AssertEqual(t, response.Data.Config["host"], DESTINATION_MODIFY_HOST)
	testutils.AssertEqual(t, response.Data.Config["port"], DESTINATION_MODIFY_PORT)
	testutils.AssertEqual(t, response.Data.Config["database"], DESTINATION_MODIFY_DATABASE)
	testutils.AssertEqual(t, response.Data.Config["auth"], DESTINATION_MODIFY_AUTH)
	testutils.AssertEqual(t, response.Data.Config["user"], DESTINATION_MODIFY_USER)
	testutils.AssertEqual(t, response.Data.Config["password"], DESTINATION_MODIFY_MASKED)
}

func assertModifyRequest(t *testing.T, request map[string]interface{}) {
	c, ok := request["config"]
	testutils.AssertEqual(t, ok, true)
	config, ok := c.(map[string]interface{})
	testutils.AssertEqual(t, ok, true)

	testutils.AssertKey(t, "host", config, DESTINATION_MODIFY_HOST)
	testutils.AssertKey(t, "port", config, float64(1433))
	testutils.AssertKey(t, "database", config, DESTINATION_MODIFY_DATABASE)
	testutils.AssertKey(t, "auth", config, DESTINATION_MODIFY_AUTH)
	testutils.AssertKey(t, "user", config, DESTINATION_MODIFY_USER)
	testutils.AssertKey(t, "password", config, DESTINATION_MODIFY_PASSWORD)
}
