package fivetran

import (
    "context"
    "encoding/json"
    "fmt"
)

// WebhookListService implements the Webhook Management, retrieve List Webhook.
// Ref. https://fivetran.com/docs/rest-api/webhooks#retrievethelistofwebhooks
type WebhookListService struct {
    c                 *Client
    limit  *int
    cursor *string
}

type WebhookListResponse struct {
    Code    string `json:"code"`
    Data    struct {
        Items []struct {
            Id             string       `json:"id"`
            Type           string       `json:"type"`
            Url            string       `json:"url"`
            GroupId        string       `json:"group_id"`
            Events         []string     `json:"events"`
            Active         bool         `json:"active"`
            Secret         string       `json:"secret"`
            CreatedAt      string       `json:"created_at"`
            CreatedBy      string       `json:"created_by"`
        } `json:"items"`
        NextCursor string `json:"next_cursor"`
    } `json:"data"`
}

func (c *Client) NewWebhookList() *WebhookListService {
    return &WebhookListService{c: c}
}

func (s *WebhookListService) Limit(value int) *WebhookListService {
    s.limit = &value
    return s
}

func (s *WebhookListService) Cursor(value string) *WebhookListService {
    s.cursor = &value
    return s
}

func (s *WebhookListService) Do(ctx context.Context) (WebhookListResponse, error) {
    var response WebhookListResponse

    url := fmt.Sprintf("%v/webhooks", s.c.baseURL)
    expectedStatus := 200

    headers := s.c.commonHeaders()
    
    queries := make(map[string]string)
    if s.cursor != nil {
        queries["cursor"] = *s.cursor
    }
    if s.limit != nil {
        queries["limit"] = fmt.Sprint(*s.limit)
    }

    r := request{
        method:           "GET",
        url:              url,
        body:             nil,
        queries:          queries,
        headers:          headers,
        client:           s.c.httpClient,
        handleRateLimits: s.c.handleRateLimits,
        maxRetryAttempts: s.c.maxRetryAttempts,
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
