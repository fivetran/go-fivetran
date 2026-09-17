package connections_test

import (
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

// Example demonstrates the expected usage of FetchSourceColumnsService
func TestFetchSourceColumnsServiceUsageExample(t *testing.T) {
	ftClient, _ := testutils.CreateTestClient()

	// This is how the service will be used in terraform-provider-fivetran
	service := ftClient.NewFetchSourceColumnsService().
		ConnectionId("connection_id_123").
		Schema("public").
		Tables([]string{"users", "orders", "products"})

	// Verify the service is created correctly
	if service == nil {
		t.Error("FetchSourceColumnsService should not be nil")
	}
}
