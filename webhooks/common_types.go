package webhooks

import "github.com/fivetran/go-fivetran/common"

type WebhookCommonData struct {
	Id        string   `json:"id"`
	Type      string   `json:"type"`
	Url       string   `json:"url"`
	Events    []string `json:"events"`
	Active    bool     `json:"active"`
	Secret    string   `json:"secret"`
	GroupId   string   `json:"group_id"`
	CreatedAt string   `json:"created_at"`
	CreatedBy string   `json:"created_by"`
}

type WebhookResponse struct {
	common.CommonResponse
	Data struct {
		WebhookCommonData
	} `json:"data"`
}

type WebhookListResponse struct {
	Code string `json:"code"`
	Data struct {
		Items      []WebhookCommonData `json:"items"`
		NextCursor string              `json:"next_cursor"`
	} `json:"data"`
}
