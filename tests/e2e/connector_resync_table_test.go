package fivetran_test

import (
	"testing"
)

func TestNewConnectorReSyncTableE2E(t *testing.T) {
	t.Skip("Does not supported yet. To test it we need a real connector which supports Schema and send POST /connectors/{connectorId}/schemas/reload before the resync to catch a schema-table information")
}