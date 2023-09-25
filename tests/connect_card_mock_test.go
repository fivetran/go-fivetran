package tests

import (
    "context"
    "fmt"
    "net/http"
    "testing"

    "github.com/fivetran/go-fivetran"
    "github.com/fivetran/go-fivetran/tests/mock"
)

func TestNewConnectCard(t *testing.T) {

    // arrange
    ftClient, mockClient := CreateTestClient()
    handler := mockClient.When(http.MethodPost, "/v1/connectors/connectorId/connect-card").ThenCall(

        func(req *http.Request) (*http.Response, error) {
            requestBodyToJson(t, req)
            response := mock.NewResponse(req, http.StatusOK, prepareConnectCardResponse())
            return response, nil
        })

    config := ftClient.NewConnectCardConfig().
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
    assertEqual(t, len(interactions), 1)
    assertEqual(t, interactions[0].Handler, handler)
    assertEqual(t, handler.Interactions, 1)
    
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

func assertConnectCardResponse(t *testing.T, response fivetran.ConnectCardResponse) {
    assertEqual(t, response.Code, "Success")
    assertNotEmpty(t, response.Message)
    assertNotEmpty(t, response.Data.ConnectorId)
    assertNotEmpty(t, response.Data.ConnectCard.Token)
    assertNotEmpty(t, response.Data.ConnectCard.Uri)
    assertEqual(t, response.Data.ConnectCardConfig.RedirectUri, "http://test_domain.com")
    assertEqual(t, response.Data.ConnectCardConfig.HideSetupGuide, false)
}