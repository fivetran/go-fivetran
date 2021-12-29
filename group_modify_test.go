package fivetran_test

import (
	"context"
	"testing"
)

func TestNewGroupModifyIntegration(t *testing.T) {
	for version, c := range Clients {
		t.Run(version, func(t *testing.T) {
			groupId := CreateTempGroup(t)

			updated, err := c.NewGroupModify().GroupID(groupId).Name("test_updated").Do(context.Background())
			if err != nil {
				t.Logf("%+v\n", updated)
				t.Error(err)
			}

			AssertEqual(t, updated.Code, "Success")
			AssertEqual(t, updated.Data.ID, groupId)
			AssertEqual(t, updated.Data.Name, "test_updated")
			AssertEqual(t, updated.Data.CreatedAt.IsZero(), false)
		})
	}
}
