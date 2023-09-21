package fivetran

import (
    "context"
    "encoding/json"
    "fmt"
)

// TeamConnectorMembershipsListService implements the Team Management, List all connector memberships
// Ref. https://fivetran.com/docs/rest-api/teams#listallconnectormemberships
type TeamConnectorMembershipsListService struct {
    c           *Client
    teamId      *string
    limit       *int
    cursor      *string
}

type TeamConnectorMembershipsListResponse struct {
    Code    string `json:"code"`
    Data    struct {
        Items []struct {
            ConnectorId   string       `json:"id"`
            Role          string       `json:"role"`
            CreatedAt     string       `json:"created_at"`
        } `json:"items"`
        NextCursor string `json:"next_cursor"`
    } `json:"data"`
}

func (c *Client) NewTeamConnectorMembershipsList() *TeamConnectorMembershipsListService {
    return &TeamConnectorMembershipsListService{c: c}
}

func (s *TeamConnectorMembershipsListService) TeamId(value string) *TeamConnectorMembershipsListService {
    s.teamId = &value
    return s
}

func (s *TeamConnectorMembershipsListService) Limit(value int) *TeamConnectorMembershipsListService {
    s.limit = &value
    return s
}

func (s *TeamConnectorMembershipsListService) Cursor(value string) *TeamConnectorMembershipsListService {
    s.cursor = &value
    return s
}

func (s *TeamConnectorMembershipsListService) Do(ctx context.Context) (TeamConnectorMembershipsListResponse, error) {
    var response TeamConnectorMembershipsListResponse

    if s.teamId == nil {
        return response, fmt.Errorf("missing required teamId")
    }

    url := fmt.Sprintf("%v/teams/%v/connectors", s.c.baseURL, *s.teamId)
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
