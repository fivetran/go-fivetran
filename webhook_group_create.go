package fivetran

import (
    "context"
    "encoding/json"
    "fmt"
)

// WebhookGroupCreateService implements the Webhook Management, Create a Group Webhook.
// Ref. https://fivetran.com/docs/rest-api/webhooks#creategroupwebhook
type WebhookGroupCreateService struct {
    c                 *Client
    groupId           *string
    url               *string
    events            *[]string
    active            *bool
    secret            *string
}

type webhookGroupCreateRequest struct {
    Url               *string           `json:"url,omitempty"`
    Events            *[]string         `json:"events,omitempty"`
    Active            *bool             `json:"active,omitempty"`
    Secret            *string           `json:"secret,omitempty"`
}


type WebhookGroupCreateResponse struct {
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

func (c *Client) NewWebhookGroupCreate() *WebhookGroupCreateService {
    return &WebhookGroupCreateService{c: c}
}

func (s *WebhookGroupCreateService) request() *webhookGroupCreateRequest {
    return &webhookGroupCreateRequest{
        Url:              s.url,
        Events:           s.events,
        Active:           s.active,
        Secret:           s.secret,
    }
}

func (s *WebhookGroupCreateService) Url(value string) *WebhookGroupCreateService {
    s.url = &value
    return s
}

func (s *WebhookGroupCreateService) Secret(value string) *WebhookGroupCreateService {
    s.secret = &value
    return s
}

func (s *WebhookGroupCreateService) Active(value bool) *WebhookGroupCreateService {
    s.active = &value
    return s
}

func (s *WebhookGroupCreateService) Events(value []string) *WebhookGroupCreateService {
    s.events = &value
    return s
}

func (s *WebhookGroupCreateService) GroupId(value string) *WebhookGroupCreateService {
    s.groupId = &value
    return s
}

func (s *WebhookGroupCreateService) Do(ctx context.Context) (WebhookGroupCreateResponse, error) {
    var response WebhookGroupCreateResponse

    if s.groupId == nil {
        return response, fmt.Errorf("missing required groupId")
    }


    url := fmt.Sprintf("%v/webhooks/group/%v", s.c.baseURL, *s.groupId)
    expectedStatus := 201

    headers := s.c.commonHeaders()
    headers["Content-Type"] = "application/json"
    headers["Accept"] = restAPIv2

    reqBody, err := json.Marshal(s.request())
    if err != nil {
        return response, err
    }

    r := request{
        method:  "POST",
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
