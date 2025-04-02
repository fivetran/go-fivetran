package groups

import (
	"time"

	"github.com/fivetran/go-fivetran/common"
	"github.com/fivetran/go-fivetran/connections"
	"github.com/fivetran/go-fivetran/users"
)

type GroupItem struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type GroupDetailsResponse struct {
	common.CommonResponse
	Data GroupItem `json:"data"`
}

type GroupSshKeyResponse struct {
	common.CommonResponse
	Data struct {
		PublicKey string `json:"public_key"`
	}
}

type GroupServiceAccountResponse struct {
	common.CommonResponse
	Data struct {
		ServiceAccount string `json:"service_account"`
	}
}

type GroupListConnectionsResponse struct {
	common.CommonResponse
	Data struct {
		Items      []connections.DetailsResponseDataCommon `json:"items"`
		NextCursor string                                 `json:"next_cursor"`
	} `json:"data"`
}

type GroupListUsersResponse struct {
	common.CommonResponse
	Data struct {
		Items      []users.UserDetailsData `json:"items"`
		NextCursor string                  `json:"next_cursor"`
	} `json:"data"`
}

type GroupsListResponse struct {
	common.CommonResponse
	Data struct {
		Items      []GroupItem `json:"items"`
		NextCursor string      `json:"next_cursor"`
	} `json:"data"`
}

type groupCreateRequest struct {
	Name *string `json:"name,omitempty"`
}

type groupAddUserRequest struct {
	Email *string `json:"email,omitempty"`
	Role  *string `json:"role,omitempty"`
}

type groupUpdateRequest struct {
	Name *string `json:"name,omitempty"`
}
