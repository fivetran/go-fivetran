package destinations_test

import (
	"context"
	"strconv"
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/destinations"

	"github.com/fivetran/go-fivetran/tests/mock"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestDestinationSetupTestsDetailsService(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/destinations/"+ID+"/test").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareDestinationDetailsResponse())
			return response, nil
		})

	service := ftClient.NewDestinationSetupTests().DestinationID(ID)

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

	assertDestinationDetailsWithSetupTestsResponse(t, response)
}

func TestDestinationSetupTestsDetailsCustomService(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/destinations/"+ID+"/test").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareDestinationDetailsResponse())
			return response, nil
		})

	service := ftClient.NewDestinationSetupTests().DestinationID(ID)

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

	assertDestinationDetailsWithSetupTestsCustomResponse(t, response)
}

func TestDestinationSetupTestsDetailsCustomServiceIntPort(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/destinations/"+ID+"/test").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			response := mock.NewResponse(req, http.StatusOK, prepareDestinationDetailsResponseIntPort())
			return response, nil
		})

	service := ftClient.NewDestinationSetupTests().DestinationID(ID)

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

	assertDestinationDetailsWithSetupTestsCustomResponseIntPort(t, response)
}

func assertDestinationDetailsWithSetupTestsResponse(t *testing.T, response destinations.DestinationDetailsWithSetupTestsResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.ID, DESTINATION_DETAILS_ID)
	testutils.AssertEqual(t, response.Data.GroupID, DESTINATION_DETAILS_ID)
	testutils.AssertEqual(t, response.Data.Service, DESTINATION_DETAILS_SERVICE)
	testutils.AssertEqual(t, response.Data.Region, DESTINATION_DETAILS_REGION)
	testutils.AssertEqual(t, response.Data.DaylightSavingTimeEnabled, DESTINATION_DETAILS_DAYLIGHT)
	testutils.AssertEqual(t, response.Data.HybridDeploymentAgentId, DESTINATION_DETAILS_HYBRIDDEPLOYMENTAGENTID)
	testutils.AssertEqual(t, response.Data.PrivateLinkId, DESTINATION_DETAILS_PRIVATELINKID)
	testutils.AssertEqual(t, response.Data.NetworkingMethod, DESTINATION_DETAILS_NETWORKINGMETHOD)
	testutils.AssertEqual(t, response.Data.TimeZoneOffset, DESTINATION_DETAILS_TIME_ZONE)
	testutils.AssertEqual(t, response.Data.SetupStatus, DESTINATION_DETAILS_SETUP_STATUS)
	testutils.AssertEqual(t, response.Data.Config.Host, DESTINATION_DETAILS_HOST)
	testutils.AssertEqual(t, response.Data.Config.Port, DESTINATION_DETAILS_PORT)
	testutils.AssertEqual(t, response.Data.Config.Database, DESTINATION_DETAILS_DATABASE)
	testutils.AssertEqual(t, response.Data.Config.Auth, DESTINATION_DETAILS_AUTH)
	testutils.AssertEqual(t, response.Data.Config.User, DESTINATION_DETAILS_USER)
	testutils.AssertEqual(t, response.Data.Config.Password, DESTINATION_DETAILS_MASKED_PASSWORD)
}

func assertDestinationDetailsWithSetupTestsCustomResponse(t *testing.T, response destinations.DestinationDetailsWithSetupTestsCustomResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.ID, DESTINATION_DETAILS_ID)
	testutils.AssertEqual(t, response.Data.GroupID, DESTINATION_DETAILS_ID)
	testutils.AssertEqual(t, response.Data.Service, DESTINATION_DETAILS_SERVICE)
	testutils.AssertEqual(t, response.Data.DaylightSavingTimeEnabled, DESTINATION_DETAILS_DAYLIGHT)
	testutils.AssertEqual(t, response.Data.HybridDeploymentAgentId, DESTINATION_DETAILS_HYBRIDDEPLOYMENTAGENTID)
	testutils.AssertEqual(t, response.Data.PrivateLinkId, DESTINATION_DETAILS_PRIVATELINKID)
	testutils.AssertEqual(t, response.Data.NetworkingMethod, DESTINATION_DETAILS_NETWORKINGMETHOD)
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

func assertDestinationDetailsWithSetupTestsCustomResponseIntPort(t *testing.T, response destinations.DestinationDetailsWithSetupTestsCustomResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertEqual(t, response.Data.ID, DESTINATION_DETAILS_ID)
	testutils.AssertEqual(t, response.Data.GroupID, DESTINATION_DETAILS_ID)
	testutils.AssertEqual(t, response.Data.Service, DESTINATION_DETAILS_SERVICE)
	testutils.AssertEqual(t, response.Data.DaylightSavingTimeEnabled, DESTINATION_DETAILS_DAYLIGHT)
	testutils.AssertEqual(t, response.Data.HybridDeploymentAgentId, DESTINATION_DETAILS_HYBRIDDEPLOYMENTAGENTID)
	testutils.AssertEqual(t, response.Data.PrivateLinkId, DESTINATION_DETAILS_PRIVATELINKID)
	testutils.AssertEqual(t, response.Data.NetworkingMethod, DESTINATION_DETAILS_NETWORKINGMETHOD)
	testutils.AssertEqual(t, response.Data.Region, DESTINATION_DETAILS_REGION)
	testutils.AssertEqual(t, response.Data.TimeZoneOffset, DESTINATION_DETAILS_TIME_ZONE)
	testutils.AssertEqual(t, response.Data.SetupStatus, DESTINATION_DETAILS_SETUP_STATUS)
	testutils.AssertEqual(t, response.Data.Config["host"], DESTINATION_DETAILS_HOST)
	port, _ := strconv.ParseFloat(DESTINATION_DETAILS_PORT, 64)
	testutils.AssertEqual(t, response.Data.Config["port"], port)
	testutils.AssertEqual(t, response.Data.Config["database"], DESTINATION_DETAILS_DATABASE)
	testutils.AssertEqual(t, response.Data.Config["auth"], DESTINATION_DETAILS_AUTH)
	testutils.AssertEqual(t, response.Data.Config["user"], DESTINATION_DETAILS_USER)
	testutils.AssertEqual(t, response.Data.Config["password"], DESTINATION_DETAILS_MASKED_PASSWORD)
}