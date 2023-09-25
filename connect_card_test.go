package fivetran_test

import (
    "context"
    "testing"
)

func TestNewConnectCardE2E(t *testing.T) {
    connectorId := CreateTempConnector(t)

    config := Client.NewConnectCardConfig().
        RedirectUri("http://test_domain.com").
        HideSetupGuide(false)

    created, err := Client.NewConnectCard().
        ConnectorId(connectorId).
        Config(config).
        Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", created)
        t.Error(err)
    }

    AssertEqual(t, created.Code, "Success")
    AssertNotEmpty(t, created.Message)
    AssertNotEmpty(t, created.Data.ConnectorId)
    AssertNotEmpty(t, created.Data.ConnectCard.Token)
    AssertNotEmpty(t, created.Data.ConnectCard.Uri)
    AssertEqual(t, created.Data.ConnectCardConfig.RedirectUri, "http://test_domain.com")
    AssertEqual(t, created.Data.ConnectCardConfig.HideSetupGuide, false)
}