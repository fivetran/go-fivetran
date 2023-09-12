package fivetran_test

import (
	"context"
	"strings"
	"testing"
)

func TestNewProjectDeleteE2E(t *testing.T) {
	projectId := CreateDbtProject(t)
	deleted, err := Client.NewDbtProjectDelete().DbtProjectID(projectId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}
	AssertEqual(t, deleted.Code, "Success")
	AssertNotEmpty(t, deleted.Message)
	AssertEqual(t, strings.Contains(deleted.Message, projectId), true)

	resp, err := Client.NewDbtProjectDetails().DbtProjectID(projectId).Do(context.Background())

	AssertEqual(t, err.Error(), "status code: 404; expected: 200")
	AssertEqual(t, strings.HasPrefix(resp.Code, "NotFound"), true)

	t.Cleanup(func() { cleanupDbtProjects() })
}
