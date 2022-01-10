package fivetran_test

import (
	"context"
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
	AssertEqual(t, deleted.Message, "Group with id '"+groupId+"' has been deleted")
}
