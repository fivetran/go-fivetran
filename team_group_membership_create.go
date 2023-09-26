package fivetran

import (
    "context"
    "encoding/json"
    "fmt"
)

// TeamGroupMembershipCreateService implements the Team Management, Add group membership
// Ref. https://fivetran.com/docs/rest-api/teams#addgroupmembership
type TeamGroupMembershipCreateService struct {
    c                 *Client
    teamId            *string
    groupId           *string
    role              *string
}

type teamGroupMembershipCreateRequest struct {
    GroupId           *string `json:"id,omitempty"`
    Role              *string `json:"role,omitempty"`
}

type TeamGroupMembershipCreateResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Data    struct {
        GroupId    string       `json:"id"`
        Role       string       `json:"role"`
        CreatedAt  string       `json:"created_at"`
    } `json:"data"`
}

func (c *Client) NewTeamGroupMembershipCreate() *TeamGroupMembershipCreateService {
    return &TeamGroupMembershipCreateService{c: c}
}

func (s *TeamGroupMembershipCreateService) request() *teamGroupMembershipCreateRequest {
    return &teamGroupMembershipCreateRequest{
        GroupId:        s.groupId,
        Role:           s.role,
    }
}

func (s *TeamGroupMembershipCreateService) TeamId(value string) *TeamGroupMembershipCreateService {
    s.teamId = &value
    return s
}

func (s *TeamGroupMembershipCreateService) GroupId(value string) *TeamGroupMembershipCreateService {
    s.groupId = &value
    return s
}

func (s *TeamGroupMembershipCreateService) Role(value string) *TeamGroupMembershipCreateService {
    s.role = &value
    return s
}

func (s *TeamGroupMembershipCreateService) Do(ctx context.Context) (TeamGroupMembershipCreateResponse, error) {
    var response TeamGroupMembershipCreateResponse
    
    if s.teamId == nil {
        return response, fmt.Errorf("missing required teamId")
    }

    url := fmt.Sprintf("%v/teams/%v/groups", s.c.baseURL, *s.teamId)
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
