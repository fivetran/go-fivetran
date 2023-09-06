package fivetran_test

import (
	"context"
	"strings"
	"testing"
)

func TestNewProjectDeleteE2E(t *testing.T) {
	t.Skip("Skip test until api is ready")
	projectId := CreateDbtProject(t)
	deleted, err := Client.NewDbtProjectDelete().DbtProjectID(projectId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}
	AssertEqual(t, deleted.Code, "Success")
	AssertNotEmpty(t, deleted.Message)
	AssertEqual(t, strings.Contains(deleted.Message, projectId), true)
}
