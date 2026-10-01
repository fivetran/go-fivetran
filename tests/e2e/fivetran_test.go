package fivetran_test

import (
	"testing"
	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func init() {
	testutils.InitE2E()
}

func TestCleanupBeforeE2E(t *testing.T) {
	// Run aggressive cleanup at the start to ensure no leftover resources
	t.Log("Running pre-test cleanup...")
	testutils.CleanupAccount()
	testutils.CleanupAccount()
	testutils.CleanupAccount()
}
