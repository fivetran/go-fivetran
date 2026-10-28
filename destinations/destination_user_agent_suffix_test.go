package destinations_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	fivetran "github.com/fivetran/go-fivetran"

	"github.com/fivetran/go-fivetran/tests/mock"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

const USER_AGENT_SUFFIX = "fivetran_destination_v2"

func expectedUserAgent() string {
	return fmt.Sprintf("%s terraform-provider-fivetran/1.9.37 %s", fivetran.DefaultUserAgent, USER_AGENT_SUFFIX)
}

func TestDestinationDetailsServiceDoCustomWithUserAgentSuffix(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	ftClient.CustomUserAgent("terraform-provider-fivetran/1.9.37")
	handler := mockClient.When(http.MethodGet, "/v1/destinations/"+ID).ThenCall(
		func(req *http.Request) (*http.Response, error) {
			testutils.AssertEqual(t, req.Header.Get("User-Agent"), expectedUserAgent())
			return mock.NewResponse(req, http.StatusOK, prepareDestinationDetailsResponse()), nil
		})

	// act
	response, err := ftClient.NewDestinationDetails().
		DestinationID(ID).
		DoCustomWithUserAgentSuffix(context.Background(), USER_AGENT_SUFFIX)

	// assert
	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertDestinationDetailsCustomResponse(t, response)
}

func TestDestinationDetailsServiceDoCustomWithUserAgentSuffixMissingDestinationID(t *testing.T) {
	// arrange
	ftClient, _ := testutils.CreateTestClient()

	// act
	_, err := ftClient.NewDestinationDetails().DoCustomWithUserAgentSuffix(context.Background(), USER_AGENT_SUFFIX)

	// assert
	testutils.AssertEqual(t, err, fmt.Errorf("missing required destinationID"))
}

func TestDestinationCreateServiceDoCustomWithUserAgentSuffix(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	ftClient.CustomUserAgent("terraform-provider-fivetran/1.9.37")
	handler := mockClient.When(http.MethodPost, "/v1/destinations").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			testutils.AssertEqual(t, req.Header.Get("User-Agent"), expectedUserAgent())
			body := testutils.RequestBodyToJson(t, req)
			assertRequest(t, body)
			return mock.NewResponse(req, http.StatusCreated, prepareDestinationResponse()), nil
		})

	// act
	response, err := ftClient.NewDestinationCreate().
		Service(SERVICE).
		Region(REGION).
		TimeZoneOffset(TIME_ZONE).
		TrustCertificates(TRUST_CERTIFICATES).
		TrustFingerprints(TRUST_FINGERPRINTS).
		RunSetupTests(RUN_SETUP_TESTS).
		DaylightSavingTimeEnabled(DAYLIGHTSAVINGTIMEENABLED).
		HybridDeploymentAgentId(HYBRIDDEPLOYMENTAGENTID).
		PrivateLinkId(PRIVATELINKID).
		ProxyAgentId(PROXY_AGENT_ID).
		NetworkingMethod(NETWORKINGMETHOD).
		GroupID(ID).
		ConfigCustom(prepareConfigCustom()).
		DoCustomWithUserAgentSuffix(context.Background(), USER_AGENT_SUFFIX)

	// assert
	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	asserResponseCustom(t, response)
}

func TestDestinationUpdateServiceDoCustomWithUserAgentSuffix(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	ftClient.CustomUserAgent("terraform-provider-fivetran/1.9.37")
	handler := mockClient.When(http.MethodPatch, "/v1/destinations/"+DESTINATION_UPDATE_ID).ThenCall(
		func(req *http.Request) (*http.Response, error) {
			testutils.AssertEqual(t, req.Header.Get("User-Agent"), expectedUserAgent())
			body := testutils.RequestBodyToJson(t, req)
			assertUpdateRequest(t, body)
			return mock.NewResponse(req, http.StatusOK, prepareDestinationUpdateResponse()), nil
		})

	// act
	response, err := ftClient.NewDestinationUpdate().
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
		RunSetupTests(RUN_SETUP_TESTS).
		DoCustomWithUserAgentSuffix(context.Background(), USER_AGENT_SUFFIX)

	// assert
	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
	assertDestinationUpdateCustomResponse(t, response)
}

func TestDestinationUpdateServiceDoCustomWithUserAgentSuffixMissingDestinationID(t *testing.T) {
	// arrange
	ftClient, _ := testutils.CreateTestClient()

	// act
	_, err := ftClient.NewDestinationUpdate().DoCustomWithUserAgentSuffix(context.Background(), USER_AGENT_SUFFIX)

	// assert
	testutils.AssertEqual(t, err, fmt.Errorf("missing required DestinationID"))
}

func TestDestinationDeleteServiceDoWithUserAgentSuffix(t *testing.T) {
	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	ftClient.CustomUserAgent("terraform-provider-fivetran/1.9.37")
	handler := mockClient.When(http.MethodDelete, "/v1/destinations/destination_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			testutils.AssertEqual(t, req.Header.Get("User-Agent"), expectedUserAgent())
			return mock.NewResponse(req, http.StatusOK, `{"code": "Success", "message": "Destination deleted"}`), nil
		})

	// act
	response, err := ftClient.NewDestinationDelete().
		DestinationID("destination_id").
		DoWithUserAgentSuffix(context.Background(), USER_AGENT_SUFFIX)

	// assert
	if err != nil {
		t.Error(err)
	}

	assertDestinationDeleteResponse(t, response, "Success", "Destination deleted")

	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)
}

func TestDestinationDeleteServiceDoWithUserAgentSuffixMissingDestinationID(t *testing.T) {
	// arrange
	ftClient, _ := testutils.CreateTestClient()

	// act
	_, err := ftClient.NewDestinationDelete().DoWithUserAgentSuffix(context.Background(), USER_AGENT_SUFFIX)

	// assert
	testutils.AssertEqual(t, err, fmt.Errorf("missing required destinationID"))
}
