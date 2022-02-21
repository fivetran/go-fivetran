package fivetran_test

import (
	"context"
	"strings"
	"testing"
)

func TestNewGroupDeleteE2E(t *testing.T) {
	groupId := CreateGroup(t)
	deleted, err := Client.NewGroupDelete().GroupID(groupId).Do(context.Background())
	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}
	AssertEqual(t, deleted.Code, "Success")
	AssertNotEmpty(t, deleted.Message)
	AssertEqual(t, strings.Contains(deleted.Message, groupId), true)
}
