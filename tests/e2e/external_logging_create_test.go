package fivetran_test

import (
	"context"
	"testing"

	"github.com/fivetran/go-fivetran"
	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewExternalLoggingCreateE2E(t *testing.T) {
	created, err := testutils.Client.NewExternalLoggingCreate().
		GroupId(testutils.PredefinedGroupId).
		Service("cloudwatch").
		Enabled(true).
		Config(fivetran.NewExternalLoggingConfig().
			RoleArn("arn:aws:iam::123456789012:role/FivetranLogRole").
			Region("us-east-1").
			LogGroupName("fivetran_log")).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertNotEmpty(t, created.Message)
	testutils.AssertEqual(t, created.Data.Id, testutils.PredefinedGroupId)
	testutils.AssertEqual(t, created.Data.Service, "cloudwatch")
	testutils.AssertEqual(t, created.Data.Enabled, true)
	testutils.AssertEqual(t, created.Data.Config.RoleArn, "arn:aws:iam::123456789012:role/FivetranLogRole")
	testutils.AssertEqual(t, created.Data.Config.Region, "us-east-1")

	t.Cleanup(func() { testutils.DeleteExternalLogging(t, testutils.PredefinedGroupId) })
}
