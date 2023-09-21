package fivetran

import (
    "context"
    "encoding/json"
    "fmt"
)

// TeamUsersModifyService implements the Team Management, Modify a user membership
// Ref. https://fivetran.com/docs/rest-api/teams#modifyausermembership
type TeamUserMembershipModifyService struct {
    c                 *Client
    teamId            *string
    userId            *string
    role              *string
}

type teamUserMembershipModifyRequest struct {
    Role              *string `json:"role,omitempty"`
}

type TeamUserMembershipModifyResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
}

func (c *Client) NewTeamUserMembershipModify() *TeamUserMembershipModifyService {
    return &TeamUserMembershipModifyService{c: c}
}

func (s *TeamUserMembershipModifyService) request() *teamUserMembershipModifyRequest {
    return &teamUserMembershipModifyRequest{
        Role:           s.role,
    }
}

func (s *TeamUserMembershipModifyService) TeamId(value string) *TeamUserMembershipModifyService {
    s.teamId = &value
    return s
}

func (s *TeamUserMembershipModifyService) UserId(value string) *TeamUserMembershipModifyService {
    s.userId = &value
    return s
}

func (s *TeamUserMembershipModifyService) Role(value string) *TeamUserMembershipModifyService {
    s.role = &value
    return s
}

func (s *TeamUserMembershipModifyService) Do(ctx context.Context) (TeamUserMembershipModifyResponse, error) {
    var response TeamUserMembershipModifyResponse

    if s.teamId == nil {
        return response, fmt.Errorf("missing required teamId")
    }

    if s.userId == nil {
        return response, fmt.Errorf("missing required userId")
    }

    url := fmt.Sprintf("%v/teams/%v/users/%v", s.c.baseURL, *s.teamId, *s.userId)
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
