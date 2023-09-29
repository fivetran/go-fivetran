package users

import (
	"time"

	"github.com/fivetran/go-fivetran/common"
)

type UserDetailsData struct {
	ID         string    `json:"id"`
	Email      string    `json:"email"`
	GivenName  string    `json:"given_name"`
	FamilyName string    `json:"family_name"`
	Verified   *bool     `json:"verified"`
	Invited    *bool     `json:"invited"`
	Picture    string    `json:"picture"`
	Phone      string    `json:"phone"`
	LoggedInAt time.Time `json:"logged_in_at"`
	CreatedAt  time.Time `json:"created_at"`
	Role       string    `json:"role"`
	Active     *bool     `json:"active"`
}

type UserDetailsResponse struct {
	common.CommonResponse
	Data UserDetailsData `json:"data"`
}

type UsersListResponse struct {
	common.CommonResponse
	Data struct {
		Items      []UserDetailsData `json:"items"`
		NextCursor string            `json:"next_cursor"`
	} `json:"data"`
}
