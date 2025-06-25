package proxy

import (
	"github.com/fivetran/go-fivetran/common"
)

type ProxyCreateData struct {
	AgentId          string `json:"agent_id"`
	AuthToken        string `json:"auth_token"`
	ClientCert 		 string `json:"client_cert"`
	ClientPrivateKey string `json:"client_private_key"`
}

type ProxyData struct {
	Id           string `json:"id"`
	AccountId    string `json:"account_id"`
	RegisteredAt string `json:"registered_at"`
	Region       string `json:"region"`
	CreatedBy    string `json:"created_by"`
	DisplayName  string `json:"display_name"`
	Token        string
	RegistredAt  string
	Salt         string
}

type proxyCreateRequest struct {
	DisplayName *string `json:"display_name,omitempty"`
	GroupRegion *string `json:"group_region,omitempty"`
}

type ProxyCreateResponse struct {
	common.CommonResponse
	Data ProxyCreateData `json:"data"`
}

type ProxyListResponse struct {
	Code string `json:"code"`
	Data struct {
		Items      []ProxyData `json:"items"`
		NextCursor string      `json:"next_cursor"`
	} `json:"data"`
}

type ProxyDetailsResponse struct {
	Code string    `json:"code"`
	Data ProxyData `json:"data"`
}

type ProxyConnectionMembershipsListResponse struct {
	Code string `json:"code"`
	Data struct {
		Items      []ProxyConnectionMembership `json:"items"`
		NextCursor string                      `json:"next_cursor"`
	} `json:"data"`
}

type ProxyConnectionMembership struct {
	ConnectionId string `json:"connection_id"`
}

type proxyConnectionMembershipCreateRequest struct {
	ConnectionId *string `json:"connection_id,omitempty"`
}
