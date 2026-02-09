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
	DESTINATION_UPDATE_SERVICE            		= "snowflake"
	DESTINATION_UPDATE_ID                 		= "decent_dropsy"
	DESTINATION_UPDATE_REGION             		= "GCP_AUSTRALIA_SOUTHEAST1"
	DESTINATION_UPDATE_TIME_ZONE_OFFSET   		= "+10"
	DESTINATION_UPDATE_DAYLIGHT 		  		= true
    DESTINATION_UPDATE_HYBRIDDEPLOYMENTAGENTID  = "hybrid_deployment_agent_id"
    DESTINATION_UPDATE_PRIVATELINKID            = "private_link_id"
	DESTINATION_UPDATE_PROXY_AGENT_ID           = "proxy_agent_id"
    DESTINATION_UPDATE_NETWORKINGMETHOD         = "Direct"
	DESTINATION_UPDATE_SETUP_STATUS      		= "connected"
	DESTINATION_UPDATE_TEST_TITLE_HOST    		= "Host Connection"
	DESTINATION_UPDATE_TEST_TITLE_DB      		= "Database Connection"
	DESTINATION_UPDATE_TEST_TITLE_PERM    		= "Permission Test"
	DESTINATION_UPDATE_TEST_STATUS_PASSED 		= "PASSED"
	DESTINATION_UPDATE_HOST              		= "your-account.snowflakecomputing.com"
	DESTINATION_UPDATE_PORT               		= "1433"
	DESTINATION_UPDATE_DATABASE          		= "fivetran"
	DESTINATION_UPDATE_AUTH               		= "PASSWORD"
	DESTINATION_UPDATE_USER               		= "fivetran_user"
	DESTINATION_UPDATE_PASSWORD           		= "******"
	DESTINATION_UPDATE_MASKED             		= "******"
)

func TestDestinationUpdateService(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/destinations/decent_dropsy").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertUpdateRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, prepareDestinationUpdateResponse())
			return response, nil
		})

	destinationConfig := fivetran.NewDestinationConfig()
	destinationConfig = destinationConfig.
		Host(DESTINATION_UPDATE_HOST).
		Port(1433).
		Database(DESTINATION_UPDATE_DATABASE).
		Auth(DESTINATION_UPDATE_AUTH).
		User(DESTINATION_UPDATE_USER).
		Password(DESTINATION_UPDATE_PASSWORD)

	service := ftClient.NewDestinationUpdate().
		DestinationID(DESTINATION_UPDATE_ID).
		Region(DESTINATION_UPDATE_REGION).
		TimeZoneOffset(DESTINATION_UPDATE_TIME_ZONE_OFFSET).
		Config(destinationConfig).
		DaylightSavingTimeEnabled(DESTINATION_UPDATE_DAYLIGHT).
        HybridDeploymentAgentId(DESTINATION_UPDATE_HYBRIDDEPLOYMENTAGENTID).
        PrivateLinkId(DESTINATION_UPDATE_PRIVATELINKID).
		ProxyAgentId(DESTINATION_UPDATE_PROXY_AGENT_ID).
        NetworkingMethod(DESTINATION_UPDATE_NETWORKINGMETHOD).
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

	assertDestinationUpdateResponse(t, response)
}

func TestDestinationUpdateCustomService(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPatch, "/v1/destinations/decent_dropsy").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			body := testutils.RequestBodyToJson(t, req)
			assertUpdateRequest(t, body)
			response := mock.NewResponse(req, http.StatusOK, prepareDestinationUpdateResponse())
			return response, nil
		})

	service := ftClient.NewDestinationUpdate().
		DestinationID(DESTINATION_UPDATE_ID).
		Region(DESTINATION_UPDATE_REGION).
		DaylightSavingTimeEnabled(DESTINATION_UPDATE_DAYLIGHT).
		HybridDeploymentAgentId(DESTINATION_UPDATE_HYBRIDDEPLOYMENTAGENTID).
        PrivateLinkId(DESTINATION_UPDATE_PRIVATELINKID).
		ProxyAgentId(DESTINATION_UPDATE_PROXY_AGENT_ID).
        NetworkingMethod(DESTINATION_UPDATE_NETWORKINGMETHOD).
		TimeZoneOffset(DESTINATION_UPDATE_TIME_ZONE_OFFSET).
		ConfigCustom(&map[string]interface{}{
			"host":     DESTINATION_UPDATE_HOST,
			"port":     1433,
			"database": DESTINATION_UPDATE_DATABASE,
			"auth":     DESTINATION_UPDATE_AUTH,
			"user":     DESTINATION_UPDATE_USER,
			"password": DESTINATION_UPDATE_PASSWORD,
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

	assertDestinationUpdateCustomResponse(t, response)
}

func prepareDestinationUpdateResponse() string {
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
			"proxy_agent_id": "%v",
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
		DESTINATION_UPDATE_ID,
		DESTINATION_UPDATE_ID,
		DESTINATION_UPDATE_SERVICE,
		DESTINATION_UPDATE_REGION,
		DESTINATION_UPDATE_DAYLIGHT,
		DESTINATION_UPDATE_HYBRIDDEPLOYMENTAGENTID,
        DESTINATION_UPDATE_PRIVATELINKID,
		DESTINATION_UPDATE_PROXY_AGENT_ID,
        DESTINATION_UPDATE_NETWORKINGMETHOD,
		DESTINATION_UPDATE_TIME_ZONE_OFFSET,
		DESTINATION_UPDATE_SETUP_STATUS,
		DESTINATION_UPDATE_TEST_TITLE_HOST,
		DESTINATION_UPDATE_TEST_STATUS_PASSED,
		DESTINATION_UPDATE_TEST_TITLE_DB,
		DESTINATION_UPDATE_TEST_STATUS_PASSED,
		DESTINATION_UPDATE_TEST_TITLE_PERM,
		DESTINATION_UPDATE_TEST_STATUS_PASSED,
		DESTINATION_UPDATE_HOST,
		DESTINATION_UPDATE_PORT,
		DESTINATION_UPDATE_DATABASE,
		DESTINATION_UPDATE_AUTH,
		DESTINATION_UPDATE_USER,
		DESTINATION_UPDATE_MASKED)
}

func assertDestinationUpdateResponse(t *testing.T, response destinations.DestinationDetailsWithSetupTestsResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Message, "Destination has been updated")
	testutils.AssertEqual(t, response.Data.ID, DESTINATION_UPDATE_ID)
	testutils.AssertEqual(t, response.Data.GroupID, DESTINATION_UPDATE_ID)
	testutils.AssertEqual(t, response.Data.Service, DESTINATION_UPDATE_SERVICE)
	testutils.AssertEqual(t, response.Data.DaylightSavingTimeEnabled, DESTINATION_UPDATE_DAYLIGHT)
	testutils.AssertEqual(t, response.Data.HybridDeploymentAgentId, DESTINATION_UPDATE_HYBRIDDEPLOYMENTAGENTID)
	testutils.AssertEqual(t, response.Data.PrivateLinkId, DESTINATION_UPDATE_PRIVATELINKID)
	testutils.AssertEqual(t, response.Data.ProxyAgentId, DESTINATION_UPDATE_PROXY_AGENT_ID)
	testutils.AssertEqual(t, response.Data.NetworkingMethod, DESTINATION_UPDATE_NETWORKINGMETHOD)
	testutils.AssertEqual(t, response.Data.Region, DESTINATION_UPDATE_REGION)
	testutils.AssertEqual(t, response.Data.TimeZoneOffset, DESTINATION_UPDATE_TIME_ZONE_OFFSET)
	testutils.AssertEqual(t, response.Data.SetupStatus, DESTINATION_UPDATE_SETUP_STATUS)

	// assert setup tests
	testutils.AssertEqual(t, len(response.Data.SetupTests), 3)
	testutils.AssertEqual(t, response.Data.SetupTests[0].Title, DESTINATION_UPDATE_TEST_TITLE_HOST)
	testutils.AssertEqual(t, response.Data.SetupTests[0].Status, DESTINATION_UPDATE_TEST_STATUS_PASSED)
	testutils.AssertEqual(t, response.Data.SetupTests[1].Title, DESTINATION_UPDATE_TEST_TITLE_DB)
	testutils.AssertEqual(t, response.Data.SetupTests[1].Status, DESTINATION_UPDATE_TEST_STATUS_PASSED)
	testutils.AssertEqual(t, response.Data.SetupTests[2].Title, DESTINATION_UPDATE_TEST_TITLE_PERM)
	testutils.AssertEqual(t, response.Data.SetupTests[2].Status, DESTINATION_UPDATE_TEST_STATUS_PASSED)

	// assert config
	testutils.AssertEqual(t, response.Data.Config.Host, DESTINATION_UPDATE_HOST)
	testutils.AssertEqual(t, response.Data.Config.Port, DESTINATION_UPDATE_PORT)
	testutils.AssertEqual(t, response.Data.Config.Database, DESTINATION_UPDATE_DATABASE)
	testutils.AssertEqual(t, response.Data.Config.Auth, DESTINATION_UPDATE_AUTH)
	testutils.AssertEqual(t, response.Data.Config.User, DESTINATION_UPDATE_USER)
	testutils.AssertEqual(t, response.Data.Config.Password, DESTINATION_UPDATE_MASKED)
}

func assertDestinationUpdateCustomResponse(t *testing.T, response destinations.DestinationDetailsWithSetupTestsCustomResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Message, "Destination has been updated")
	testutils.AssertEqual(t, response.Data.ID, DESTINATION_UPDATE_ID)
	testutils.AssertEqual(t, response.Data.GroupID, DESTINATION_UPDATE_ID)
	testutils.AssertEqual(t, response.Data.Service, DESTINATION_UPDATE_SERVICE)
	testutils.AssertEqual(t, response.Data.DaylightSavingTimeEnabled, DESTINATION_UPDATE_DAYLIGHT)
	testutils.AssertEqual(t, response.Data.HybridDeploymentAgentId, DESTINATION_UPDATE_HYBRIDDEPLOYMENTAGENTID)
	testutils.AssertEqual(t, response.Data.PrivateLinkId, DESTINATION_UPDATE_PRIVATELINKID)
	testutils.AssertEqual(t, response.Data.ProxyAgentId, DESTINATION_UPDATE_PROXY_AGENT_ID)
	testutils.AssertEqual(t, response.Data.NetworkingMethod, DESTINATION_UPDATE_NETWORKINGMETHOD)
	testutils.AssertEqual(t, response.Data.Region, DESTINATION_UPDATE_REGION)
	testutils.AssertEqual(t, response.Data.TimeZoneOffset, DESTINATION_UPDATE_TIME_ZONE_OFFSET)
	testutils.AssertEqual(t, response.Data.SetupStatus, DESTINATION_UPDATE_SETUP_STATUS)

	// assert setup tests
	testutils.AssertEqual(t, len(response.Data.SetupTests), 3)
	testutils.AssertEqual(t, response.Data.SetupTests[0].Title, DESTINATION_UPDATE_TEST_TITLE_HOST)
	testutils.AssertEqual(t, response.Data.SetupTests[0].Status, DESTINATION_UPDATE_TEST_STATUS_PASSED)
	testutils.AssertEqual(t, response.Data.SetupTests[1].Title, DESTINATION_UPDATE_TEST_TITLE_DB)
	testutils.AssertEqual(t, response.Data.SetupTests[1].Status, DESTINATION_UPDATE_TEST_STATUS_PASSED)
	testutils.AssertEqual(t, response.Data.SetupTests[2].Title, DESTINATION_UPDATE_TEST_TITLE_PERM)
	testutils.AssertEqual(t, response.Data.SetupTests[2].Status, DESTINATION_UPDATE_TEST_STATUS_PASSED)

	// assert config
	testutils.AssertEqual(t, response.Data.Config["host"], DESTINATION_UPDATE_HOST)
	testutils.AssertEqual(t, response.Data.Config["port"], DESTINATION_UPDATE_PORT)
	testutils.AssertEqual(t, response.Data.Config["database"], DESTINATION_UPDATE_DATABASE)
	testutils.AssertEqual(t, response.Data.Config["auth"], DESTINATION_UPDATE_AUTH)
	testutils.AssertEqual(t, response.Data.Config["user"], DESTINATION_UPDATE_USER)
	testutils.AssertEqual(t, response.Data.Config["password"], DESTINATION_UPDATE_MASKED)
}

func assertUpdateRequest(t *testing.T, request map[string]interface{}) {
	c, ok := request["config"]
	testutils.AssertEqual(t, ok, true)
	config, ok := c.(map[string]interface{})
	testutils.AssertEqual(t, ok, true)

	testutils.AssertKey(t, "host", config, DESTINATION_UPDATE_HOST)
	testutils.AssertKey(t, "port", config, float64(1433))
	testutils.AssertKey(t, "database", config, DESTINATION_UPDATE_DATABASE)
	testutils.AssertKey(t, "auth", config, DESTINATION_UPDATE_AUTH)
	testutils.AssertKey(t, "user", config, DESTINATION_UPDATE_USER)
	testutils.AssertKey(t, "password", config, DESTINATION_UPDATE_PASSWORD)
}
