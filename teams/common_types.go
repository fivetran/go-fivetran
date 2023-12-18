package teams

import (
	"github.com/fivetran/go-fivetran/common"
)

type TeamData struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Role        string `json:"role"`
}

type TeamMembership struct {
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

type TeamConnectorMembership struct {
	ConnectorId string `json:"id"`
	TeamMembership
}

type TeamGroupMembership struct {
	GroupId string `json:"id"`
	TeamMembership
}

type TeamUserMembership struct {
	UserId string `json:"user_id"`
	TeamMembership
}

type teamsCreateRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Role        *string `json:"role,omitempty"`
}

type TeamsCreateResponse struct {
	common.CommonResponse
	Data TeamData `json:"data"`
}

type teamConnectorMembershipCreateRequest struct {
	ConnectorId *string `json:"id,omitempty"`
	Role        *string `json:"role,omitempty"`
}

type TeamConnectorMembershipCreateResponse struct {
	common.CommonResponse
	Data TeamConnectorMembership `json:"data"`
}

type TeamConnectorMembershipDetailsResponse struct {
	Code string                        `json:"code"`
	Data TeamConnectorMembership `json:"data"`
}

type TeamConnectorMembershipsListResponse struct {
	Code string `json:"code"`
	Data struct {
		Items      []TeamConnectorMembership `json:"items"`
		NextCursor string                          `json:"next_cursor"`
	} `json:"data"`
}

type teamConnectorMembershipModifyRequest struct {
	Role *string `json:"role,omitempty"`
}

type teamGroupMembershipCreateRequest struct {
	GroupId *string `json:"id,omitempty"`
	Role    *string `json:"role,omitempty"`
}

type TeamGroupMembershipCreateResponse struct {
	common.CommonResponse
	Data TeamGroupMembership `json:"data"`
}

type TeamGroupMembershipDetailsResponse struct {
	Code string                    `json:"code"`
	Data TeamGroupMembership `json:"data"`
}

type TeamGroupMembershipsListResponse struct {
	Code string `json:"code"`
	Data struct {
		Items      []TeamGroupMembership `json:"items"`
		NextCursor string                      `json:"next_cursor"`
	} `json:"data"`
}

type teamGroupMembershipModifyRequest struct {
	Role *string `json:"role,omitempty"`
}

type teamUserMembershipCreateRequest struct {
	UserId *string `json:"user_id,omitempty"`
	Role   *string `json:"role,omitempty"`
}

type TeamUserMembershipCreateResponse struct {
	common.CommonResponse
	Data TeamUserMembership `json:"data"`
}

type TeamUserMembershipDetailsResponse struct {
	Code string                   `json:"code"`
	Data TeamUserMembership `json:"data"`
}

type TeamUserMembershipsListResponse struct {
	Code string `json:"code"`
	Data struct {
		Items      []TeamUserMembership `json:"items"`
		NextCursor string                     `json:"next_cursor"`
	} `json:"data"`
}

type teamUserMembershipModifyRequest struct {
	Role *string `json:"role,omitempty"`
}

type TeamsDetailsResponse struct {
	Code string         `json:"code"`
	Data TeamData `json:"data"`
}

type TeamsListResponse struct {
	Code string `json:"code"`
	Data struct {
		Items      []TeamData `json:"items"`
		NextCursor string           `json:"next_cursor"`
	} `json:"data"`
}

type teamsModifyRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Role        *string `json:"role,omitempty"`
}

type TeamsModifyResponse struct {
	common.CommonResponse
	Data TeamData `json:"data"`
}