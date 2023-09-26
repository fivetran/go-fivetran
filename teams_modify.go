package fivetran

import (
    "context"
    "encoding/json"
    "fmt"
)

// TeamsModifyService implements the Team Management, Modify a Team.
// Ref. https://fivetran.com/docs/rest-api/teams#modifyateam
type TeamsModifyService struct {
    c                 *Client
    teamId            *string
    name              *string
    description       *string
    role              *string
}

type teamsModifyRequest struct {
    Name              *string `json:"name,omitempty"`
    Description       *string `json:"description,omitempty"`
    Role              *string `json:"role,omitempty"`
}

type TeamsModifyResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Data    struct {
        Id              string       `json:"id"`
        Name            string       `json:"name"`
        Description     string       `json:"description"`
        Role            string       `json:"role"`
    } `json:"data"`
}

func (c *Client) NewTeamsModify() *TeamsModifyService {
    return &TeamsModifyService{c: c}
}

func (s *TeamsModifyService) request() *teamsModifyRequest {
    return &teamsModifyRequest{
        Name:           s.name,
        Description:    s.description,
        Role:           s.role,
    }
}

func (s *TeamsModifyService) TeamId(value string) *TeamsModifyService {
    s.teamId = &value
    return s
}

func (s *TeamsModifyService) Name(value string) *TeamsModifyService {
    s.name = &value
    return s
}

func (s *TeamsModifyService) Role(value string) *TeamsModifyService {
    s.role = &value
    return s
}

func (s *TeamsModifyService) Description(value string) *TeamsModifyService {
    s.description = &value
    return s
}

func (s *TeamsModifyService) Do(ctx context.Context) (TeamsModifyResponse, error) {
    var response TeamsModifyResponse

    if s.teamId == nil {
        return response, fmt.Errorf("missing required teamId")
    }

    url := fmt.Sprintf("%v/teams/%v", s.c.baseURL, *s.teamId)
    expectedStatus := 200

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
