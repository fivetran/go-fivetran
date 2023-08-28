package fivetran

import (
    "context"
    "encoding/json"
    "fmt"
)

// WebhookDetailsService implements the Webhook Management, retrieve Webhook Details.
// Ref. https://fivetran.com/docs/rest-api/webhooks#retrievewebhookdetails
type WebhookDetailsService struct {
    c                 *Client
    webhookId         *string
}

type WebhookDetailsResponse struct {
    Code    string `json:"code"`
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

func (c *Client) NewWebhookDetails() *WebhookDetailsService {
    return &WebhookDetailsService{c: c}
}

func (s *WebhookDetailsService) WebhookId(value string) *WebhookDetailsService {
    s.webhookId = &value
    return s
}

func (s *WebhookDetailsService) Do(ctx context.Context) (WebhookDetailsResponse, error) {
    var response WebhookDetailsResponse

    if s.webhookId == nil {
        return response, fmt.Errorf("missing required webhookId")
    }

    url := fmt.Sprintf("%v/webhooks/%v", s.c.baseURL, *s.webhookId)
    expectedStatus := 201

    headers := s.c.commonHeaders()
    headers["Content-Type"] = "application/json"
    headers["Accept"] = restAPIv2

    r := request{
        method:  "POST",
        url:     url,
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
