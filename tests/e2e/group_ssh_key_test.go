package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewGroupSshPublicKeyE2E(t *testing.T) {
	result, err := testutils.Client.NewGroupSshPublicKey().GroupID(testutils.PredefinedGroupId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", result)
		t.Error(err)
	}

	testutils.AssertEqual(t, result.Code, "Success")
	testutils.AssertNotEmpty(t, result.Data.PublicKey)

}
