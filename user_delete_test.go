package fivetran_test

import (
	"context"
	"strings"
	"testing"
)

func TestNewUserDeleteE2E(t *testing.T) {
	userId := CreateUser(t)
	deleted, err := Client.NewUserDelete().UserID(userId).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}
	AssertEqual(t, deleted.Code, "Success")
	AssertNotEmpty(t, deleted.Message)
	AssertEqual(t, strings.Contains(deleted.Message, userId), true)
}
