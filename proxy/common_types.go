package proxy

import (
	"github.com/fivetran/go-fivetran/common"
)

type ProxyCreateData struct {
	AgentId       	string `json:"agent_id"`
	AuthToken       string `json:"auth_token"`
	ProxyServerUri 	string `json:"proxy_server_uri"`
}

type ProxyData struct {
	Id          	string `json:"id"`
	AccountId       string `json:"account_id"`
	RegistredAt 	string `json:"registred_at"`
	Region 			string `json:"region"`
	Token 			string `json:"token"`
	Salt 			string `json:"salt"`
	CreatedBy 		string `json:"created_by"`
	DisplayName 	string `json:"display_name"`
}

type proxyCreateRequest struct {
	DisplayName   *string `json:"display_name,omitempty"`
	GroupId 	  *string `json:"group_id,omitempty"`
}

type ProxyCreateResponse struct {
	common.CommonResponse
	Data   ProxyCreateData `json:"data"`
}

type ProxyListResponse struct {
	Code string `json:"code"`
	Data struct {
		Items      []ProxyData `json:"items"`
		NextCursor string      `json:"next_cursor"`
	} `json:"data"`
}

type ProxyDetailsResponse struct {
	Code string      `json:"code"`
	Data ProxyData 	 `json:"data"`
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
	ConnectionId 	*string `json:"connection_id,omitempty"`
}