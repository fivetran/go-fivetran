package fivetran_test

import (
	"context"
	"testing"
)

func TestNewGroupDeleteIntegration(t *testing.T) {
	for version, c := range Clients {
		t.Run(version, func(t *testing.T) {
			groupId := CreateGroup(t)
			deleted, err := c.NewGroupDelete().GroupID(groupId).Do(context.Background())
			if err != nil {
				t.Logf("%+v\n", deleted)
				t.Error(err)
			}
			AssertEqual(t, deleted.Code, "Success")
			AssertEqual(t, deleted.Message, "Group with id '"+groupId+"' has been deleted")
		})
	}
}
