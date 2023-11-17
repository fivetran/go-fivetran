package fivetran_test

import (
	"context"
	"fmt"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewGroupServiceAccountE2E(t *testing.T) {
	result, err := testutils.Client.NewGroupServiceAccount().GroupID(testutils.PredefinedGroupId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertEqual(t, result.Data.ServiceAccount, fmt.Sprintf("g-%v@fivetran-production.iam.gserviceaccount.com", testutils.PredefinedGroupId))
}
