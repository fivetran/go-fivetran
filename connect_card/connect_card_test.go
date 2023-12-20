package connectcard_test

import (
    "context"
    "net/http"
    "testing"
    "fmt"

	"github.com/fivetran/go-fivetran"
    "github.com/fivetran/go-fivetran/connect_card"
    testutils "github.com/fivetran/go-fivetran/test_utils"
    
    "github.com/fivetran/go-fivetran/tests/mock"
)

func TestNewConnectCard(t *testing.T) {

	// arrange
	ftClient, mockClient := testutils.CreateTestClient()
	handler := mockClient.When(http.MethodPost, "/v1/connectors/connectorId/connect-card").ThenCall(

		func(req *http.Request) (*http.Response, error) {
			testutils.RequestBodyToJson(t, req)
			response := mock.NewResponse(req, http.StatusOK, prepareConnectCardResponse())
			return response, nil
		})

	config := fivetran.NewConnectCardConfig().
		RedirectUri("http://test_domain.com").
		HideSetupGuide(false)

	// act
	response, err := ftClient.NewConnectCard().
		ConnectorId("connectorId").
		Config(config).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	// assert
	interactions := mockClient.Interactions()
	testutils.AssertEqual(t, len(interactions), 1)
	testutils.AssertEqual(t, interactions[0].Handler, handler)
	testutils.AssertEqual(t, handler.Interactions, 1)

	assertConnectCardResponse(t, response)
}

func prepareConnectCardResponse() string {
	return fmt.Sprintf(
		`{
            "code": "Success",
            "message": "Connector Connect Card has been created",
            "data": {
                "connector_id": "connectorId",
                "connect_card": {
                    "token": "token",
                    "uri": "uri"
                },
                "connect_card_config": {
                    "redirect_uri": "http://test_domain.com",
                    "hide_setup_guide": false
                }
            }
        }`)
}

func assertConnectCardResponse(t *testing.T, response connectcard.ConnectCardResponse) {
	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)
	testutils.AssertNotEmpty(t, response.Data.ConnectorId)
	testutils.AssertNotEmpty(t, response.Data.ConnectCard.Token)
	testutils.AssertNotEmpty(t, response.Data.ConnectCard.Uri)
	testutils.AssertEqual(t, response.Data.ConnectCardConfig.RedirectUri, "http://test_domain.com")
	testutils.AssertEqual(t, response.Data.ConnectCardConfig.HideSetupGuide, false)
}
