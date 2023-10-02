package fivetran_test

import (
	"context"
	"strings"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewProjectDeleteE2E(t *testing.T) {
	projectId := testutils.CreateDbtProject(t)
	deleted, err := testutils.Client.NewDbtProjectDelete().DbtProjectID(projectId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}
	testutils.AssertEqual(t, deleted.Code, "Success")
	testutils.AssertNotEmpty(t, deleted.Message)
	testutils.AssertEqual(t, strings.Contains(deleted.Message, projectId), true)

	resp, err := testutils.Client.NewDbtProjectDetails().DbtProjectID(projectId).Do(context.Background())

	testutils.AssertEqual(t, err.Error(), "status code: 404; expected: 200")
	testutils.AssertEqual(t, strings.HasPrefix(resp.Code, "NotFound"), true)

	t.Cleanup(func() {
		testutils.CleanupDbtProjects()
		testutils.DeleteDbtDestination()
	})
}
