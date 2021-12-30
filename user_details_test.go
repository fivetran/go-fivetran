package fivetran_test

import (
	"context"
	"testing"
)

func TestNewUserDetailsE2E(t *testing.T) {
	for version, c := range Clients {
		t.Run(version, func(t *testing.T) {
			user, err := c.NewUserDetails().UserID(PredefinedUserId).Do(context.Background())
			if err != nil {
				t.Logf("%+v\n", user)
				t.Error(err)
			}

			AssertEqual(t, user.Code, "Success")
			AssertEqual(t, user.Data.ID, PredefinedUserId)
			AssertEqual(t, user.Data.Email, "testingfivetran@gmail.com")
			AssertEqual(t, user.Data.GivenName, "Andrey")
			AssertEqual(t, user.Data.FamilyName, "Markov")
			AssertEqual(t, *user.Data.Verified, true)
			AssertEqual(t, *user.Data.Invited, false)
			AssertEmpty(t, user.Data.Picture)
			AssertEqual(t, user.Data.Phone, "89534322340")
			//AssertEqual(t, user.Data.Role, "account_administrator")
			AssertEqual(t, user.Data.LoggedInAt.IsZero(), true)
			AssertEqual(t, user.Data.CreatedAt.IsZero(), false)
		})
	}
}
