package fivetran_test

import (
	"context"
	"testing"
)

func TestNewUserDeleteE2E(t *testing.T) {
	for version, c := range Clients {
		t.Run(version, func(t *testing.T) {
			userId := CreateUser(t)
			deleted, err := c.NewUserDelete().UserID(userId).Do(context.Background())

			if err != nil {
				t.Logf("%+v\n", deleted)
				t.Error(err)
			}
			AssertEqual(t, deleted.Code, "Success")
			AssertEqual(t, deleted.Message, "User with id '"+userId+"' has been deleted")
		})
	}
}
