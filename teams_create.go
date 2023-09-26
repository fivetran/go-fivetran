package fivetran

import (
    "context"
    "encoding/json"
    "fmt"
)

// TeamsCreateService implements the Team Management, Create a Team.
// Ref. https://fivetran.com/docs/rest-api/teams#createateam
type TeamsCreateService struct {
    c                 *Client
    name              *string
    description       *string
    role              *string
}

type teamsCreateRequest struct {
    Name              *string `json:"name,omitempty"`
    Description       *string `json:"description,omitempty"`
    Role              *string `json:"role,omitempty"`
}

type TeamsCreateResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Data    struct {
        Id              string       `json:"id"`
        Name            string       `json:"name"`
        Description     string       `json:"description"`
        Role            string       `json:"role"`
    } `json:"data"`
}

func (c *Client) NewTeamsCreate() *TeamsCreateService {
    return &TeamsCreateService{c: c}
}

func (s *TeamsCreateService) request() *teamsCreateRequest {
    return &teamsCreateRequest{
        Name:           s.name,
        Description:    s.description,
        Role:           s.role,
    }
}

func (s *TeamsCreateService) Name(value string) *TeamsCreateService {
    s.name = &value
    return s
}

func (s *TeamsCreateService) Role(value string) *TeamsCreateService {
    s.role = &value
    return s
}

func (s *TeamsCreateService) Description(value string) *TeamsCreateService {
    s.description = &value
    return s
}

func (s *TeamsCreateService) Do(ctx context.Context) (TeamsCreateResponse, error) {
    var response TeamsCreateResponse
    url := fmt.Sprintf("%v/teams", s.c.baseURL)
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
