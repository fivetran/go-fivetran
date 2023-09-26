package fivetran

import (
    "context"
    "encoding/json"
    "fmt"
)

// TeamsListService implements the Team Management, retrieve List all Teams.
// Ref. https://fivetran.com/docs/rest-api/teams#listallteams
type TeamsListService struct {
    c           *Client
    limit       *int
    cursor      *string
}

type TeamsListResponse struct {
    Code    string `json:"code"`
    Data    struct {
        Items []struct {
            Id              string       `json:"id"`
            Name            string       `json:"name"`
            Description     string       `json:"description"`
            Role            string       `json:"role"`
        } `json:"items"`
        NextCursor string `json:"next_cursor"`
    } `json:"data"`
}

func (c *Client) NewTeamsList() *TeamsListService {
    return &TeamsListService{c: c}
}

func (s *TeamsListService) Limit(value int) *TeamsListService {
    s.limit = &value
    return s
}

func (s *TeamsListService) Cursor(value string) *TeamsListService {
    s.cursor = &value
    return s
}

func (s *TeamsListService) Do(ctx context.Context) (TeamsListResponse, error) {
    var response TeamsListResponse

    url := fmt.Sprintf("%v/teams", s.c.baseURL)
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
