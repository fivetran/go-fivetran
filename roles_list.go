package fivetran

import (
    "context"
    "encoding/json"
    "fmt"
)

// RolesListService implements the Group Management, List All Roles API.
// Ref. https://fivetran.com/docs/rest-api/roles
type RolesListService struct {
    c      *Client
    limit  *int
    cursor *string
}

type RolesListResponse struct {
    Code    string `json:"code"`
    Data    struct {
        Items []struct {
            Name            string      `json:"name"`
            Description     string      `json:"description"`
            IsCustom        bool        `json:"is_custom"`
            Scope           []string    `json:"scope"`
        } `json:"items"`
        NextCursor string `json:"next_cursor"`
    } `json:"data"`
}

func (c *Client) NewRolesList() *RolesListService {
    return &RolesListService{c: c}
}

func (s *RolesListService) Limit(value int) *RolesListService {
    s.limit = &value
    return s
}

func (s *RolesListService) Cursor(value string) *RolesListService {
    s.cursor = &value
    return s
}

func (s *RolesListService) Do(ctx context.Context) (RolesListResponse, error) {
    var response RolesListResponse
    url := fmt.Sprintf("%v/Roles", s.c.baseURL)
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
