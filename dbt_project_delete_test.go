package fivetran_test

import (
	"context"
	"strings"
	"testing"
)

func TestNewProjectDeleteE2E(t *testing.T) {
	groupId := CreateProject(t)
	deleted, err := Client.NewGroupDelete().GroupID(groupId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}
	AssertEqual(t, deleted.Code, "Success")
	AssertNotEmpty(t, deleted.Message)
	AssertEqual(t, strings.Contains(deleted.Message, groupId), true)
}
