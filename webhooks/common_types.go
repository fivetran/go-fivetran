package webhooks

import "github.com/fivetran/go-fivetran/common"

type webhookAccountCreateRequest struct {
	Url    *string   `json:"url,omitempty"`
	Events *[]string `json:"events,omitempty"`
	Active *bool     `json:"active,omitempty"`
	Secret *string   `json:"secret,omitempty"`
}

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

type webhookTestRequest struct {
	Event *string `json:"event,omitempty"`
}

type WebhookTestResponse struct {
	Code string `json:"code"`
	Data struct {
		Succeed bool   `json:"succeed"`
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"data"`
}

type webhookUpdateRequest struct {
	Url      *string   `json:"url,omitempty"`
	Events   *[]string `json:"events,omitempty"`
	Active   *bool     `json:"active,omitempty"`
	Secret   *string   `json:"secret,omitempty"`
	RunTests *bool     `json:"run_tests,omitempty"`
}

type webhookGroupCreateRequest struct {
	Url    *string   `json:"url,omitempty"`
	Events *[]string `json:"events,omitempty"`
	Active *bool     `json:"active,omitempty"`
	Secret *string   `json:"secret,omitempty"`
}