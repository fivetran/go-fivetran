package fivetran

import (
    "context"
    "encoding/json"
    "fmt"
)

// TeamsDetailsService implements the Team Management, retrieve Team Details.
// Ref. https://fivetran.com/docs/rest-api/teams#retrieveteamdetails
type TeamsDetailsService struct {
    c                 *Client
    teamId            *string
}

type TeamsDetailsResponse struct {
    Code    string `json:"code"`
    Data    struct {
        Id              string       `json:"id"`
        Name            string       `json:"name"`
        Description     string       `json:"description"`
        Role            string       `json:"role"`
    } `json:"data"`
}

func (c *Client) NewTeamsDetails() *TeamsDetailsService {
    return &TeamsDetailsService{c: c}
}

func (s *TeamsDetailsService) TeamId(value string) *TeamsDetailsService {
    s.teamId = &value
    return s
}

func (s *TeamsDetailsService) Do(ctx context.Context) (TeamsDetailsResponse, error) {
    var response TeamsDetailsResponse

    if s.teamId == nil {
        return response, fmt.Errorf("missing required teamId")
    }

    url := fmt.Sprintf("%v/teams/%v", s.c.baseURL, *s.teamId)
    expectedStatus := 200

    headers := s.c.commonHeaders()
    headers["Content-Type"] = "application/json"
    headers["Accept"] = restAPIv2

    r := request{
        method:  "GET",
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
