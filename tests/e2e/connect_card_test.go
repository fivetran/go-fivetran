package fivetran_test

import (
    "context"
    "testing"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewConnectCardE2E(t *testing.T) {
    connectorId := testutils.CreateTempConnector(t)

    config := testutils.Client.NewConnectCardConfig().
        RedirectUri("http://test_domain.com").
        HideSetupGuide(false)

    created, err := testutils.Client.NewConnectCard().
        ConnectorId(connectorId).
        Config(config).
        Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", created)
        t.Error(err)
    }

    testutils.AssertEqual(t, created.Code, "Success")
    testutils.AssertNotEmpty(t, created.Message)
    testutils.AssertNotEmpty(t, created.Data.ConnectorId)
    testutils.AssertNotEmpty(t, created.Data.ConnectCard.Token)
    testutils.AssertNotEmpty(t, created.Data.ConnectCard.Uri)
    testutils.AssertEqual(t, created.Data.ConnectCardConfig.RedirectUri, "http://test_domain.com")
    testutils.AssertEqual(t, created.Data.ConnectCardConfig.HideSetupGuide, false)
}