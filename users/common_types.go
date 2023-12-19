package users

import (
	"time"

	"github.com/fivetran/go-fivetran/common"
)

/* Requests */
type userInviteRequest struct {
	Email      *string `json:"email,omitempty"`
	GivenName  *string `json:"given_name,omitempty"`
	FamilyName *string `json:"family_name,omitempty"`
	Phone      *string `json:"phone,omitempty"`
	Picture    *string `json:"picture,omitempty"`
	Role       *string `json:"role,omitempty"`
}

type userModifyRequest struct {
	GivenName  *string                `json:"given_name,omitempty"`
	FamilyName *string                `json:"family_name,omitempty"`
	Phone      *common.NullableString `json:"phone,omitempty"`
	Picture    *common.NullableString `json:"picture,omitempty"`
	Role       *string                `json:"role,omitempty"`
}

type userGroupMembershipCreateRequest struct {
	GroupId *string `json:"id,omitempty"`
	Role    *string `json:"role,omitempty"`
}

type userGroupMembershipModifyRequest struct {
	Role *string `json:"role,omitempty"`
}

type userConnectorMembershipCreateRequest struct {
	ConnectorId *string `json:"id,omitempty"`
	Role        *string `json:"role,omitempty"`
}

type userConnectorMembershipModifyRequest struct {
	Role *string `json:"role,omitempty"`
}

/* Responses */

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

type UserConnectorMembershipCreateResponse struct {
	common.CommonResponse
	Data UserConnectorMembership `json:"data"`
}


type UserConnectorMembershipDetailsResponse struct {
	Code string                        `json:"code"`
	Data UserConnectorMembership       `json:"data"`
}


type UserConnectorMembershipsListResponse struct {
	Code string `json:"code"`
	Data struct {
		Items      []UserConnectorMembership `json:"items"`
		NextCursor string                    `json:"next_cursor"`
	} `json:"data"`
}

type UserGroupMembershipCreateResponse struct {
	common.CommonResponse
	Data UserGroupMembership `json:"data"`
}

type UserGroupMembershipDetailsResponse struct {
	Code string                    `json:"code"`
	Data UserGroupMembership 	   `json:"data"`
}

type UserGroupMembershipsListResponse struct {
	Code string `json:"code"`
	Data struct {
		Items      []UserGroupMembership `json:"items"`
		NextCursor string                `json:"next_cursor"`
	} `json:"data"`
}

type UserMembership struct {
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

type UserConnectorMembership struct {
	ConnectorId string `json:"id"`
	UserMembership
}

type UserGroupMembership struct {
	GroupId string `json:"id"`
	UserMembership
}