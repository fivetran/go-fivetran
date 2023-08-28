package fivetran

import (
    "context"
    "encoding/json"
    "fmt"
)

// WebhookModifyService implements the Webhook Management, Modify a Webhook.
// Ref. https://fivetran.com/docs/rest-api/webhooks#updatewebhook
type WebhookModifyService struct {
    c                 *Client
    webhookId         *string
    url               *string
    events            *[]string
    active            *bool
    secret            *string
    runTests          *bool
}

type webhookModifyRequest struct {
    Url               *string           `json:"url,omitempty"`
    Events            *[]string         `json:"events,omitempty"`
    Active            *bool             `json:"active,omitempty"`
    Secret            *string           `json:"secret,omitempty"`
    RunTests          *bool             `json:"run_tests,omitempty"`
}


type WebhookModifyResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Data    struct {
        Id             string       `json:"id"`
        Type           string       `json:"type"`
        Url            string       `json:"url"`
        GroupId        string       `json:"group_id"`
        Events         []string     `json:"events"`
        Active         bool         `json:"active"`
        Secret         string       `json:"secret"`
        CreatedAt      string       `json:"created_at"`
        CreatedBy      string       `json:"created_by"`
    } `json:"data"`
}

func (c *Client) NewWebhookModify() *WebhookModifyService {
    return &WebhookModifyService{c: c}
}

func (s *WebhookModifyService) request() *webhookModifyRequest {
    return &webhookModifyRequest{
        Url:              s.url,
        Events:           s.events,
        Active:           s.active,
        Secret:           s.secret,
    }
}

func (s *WebhookModifyService) Url(value string) *WebhookModifyService {
    s.url = &value
    return s
}

func (s *WebhookModifyService) Secret(value string) *WebhookModifyService {
    s.secret = &value
    return s
}

func (s *WebhookModifyService) Active(value bool) *WebhookModifyService {
    s.active = &value
    return s
}

func (s *WebhookModifyService) Events(value []string) *WebhookModifyService {
    s.events = &value
    return s
}

func (s *WebhookModifyService) WebhookId(value string) *WebhookModifyService {
    s.webhookId = &value
    return s
}

func (s *WebhookModifyService) RunTests(value bool) *WebhookModifyService {
    s.runTests = &value
    return s
}

func (s *WebhookModifyService) Do(ctx context.Context) (WebhookModifyResponse, error) {
    var response WebhookModifyResponse

    if s.webhookId == nil {
        return response, fmt.Errorf("missing required webhookId")
    }

    url := fmt.Sprintf("%v/webhooks/%v", s.c.baseURL, *s.webhookId)
    expectedStatus := 201

    headers := s.c.commonHeaders()
    headers["Content-Type"] = "application/json"
    headers["Accept"] = restAPIv2

    reqBody, err := json.Marshal(s.request())
    if err != nil {
        return response, err
    }

    r := request{
        method:  "PATCH",
        url:     url,
        body:    reqBody,
        queries: nil,
        headers: headers,
        client:  s.c.httpClient,
    }

    respBody, respStatus, err := r.httpRequest(ctx)
    if err != nil {
        return response, err
    }

    if err := json.Unmarshal(respBody, &response); err != nil {
        return response, err
    }

    if respStatus != expectedStatus {
        err := fmt.Errorf("status code: %v; expected: %v", respStatus, expectedStatus)
        return response, err
    }

    return response, nil
}
