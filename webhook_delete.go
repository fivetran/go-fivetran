package fivetran

import (
    "context"
    "encoding/json"
    "fmt"
)

// WebhookDeleteService implements the Webhook Management, Delete a Webhook.
// Ref. https://fivetran.com/docs/rest-api/webhooks#deletewebhook
type WebhookDeleteService struct {
    c                   *Client
    webhookId           *string
}

type WebhookDeleteResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
}

func (c *Client) NewWebhookDelete() *WebhookDeleteService {
    return &WebhookDeleteService{c: c}
}

func (s *WebhookDeleteService) WebhookId(value string) *WebhookDeleteService {
    s.webhookId = &value
    return s
}

func (s *WebhookDeleteService) Do(ctx context.Context) (WebhookDeleteResponse, error) {
    var response WebhookDeleteResponse

    if s.webhookId == nil {
        return response, fmt.Errorf("missing required WebhookId")
    }

    url := fmt.Sprintf("%v/webhooks/%v", s.c.baseURL, *s.webhookId)
    expectedStatus := 200

    headers := s.c.commonHeaders()

    r := request{
        method:  "DELETE",
        url:     url,
        body:    nil,
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
