package fivetran

import (
    "context"
    "encoding/json"
    "fmt"
)

// TeamUserMembershipDeleteService implements the Team Management, Delete a user from a team
// Ref. https://fivetran.com/docs/rest-api/teams#deleteauserfromateam
type TeamUserMembershipDeleteService struct {
    c                *Client
    teamId           *string
    userId           *string
}

type TeamUserMembershipDeleteResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
}

func (c *Client) NewTeamUserMembershipDelete() *TeamUserMembershipDeleteService {
    return &TeamUserMembershipDeleteService{c: c}
}

func (s *TeamUserMembershipDeleteService) TeamId(value string) *TeamUserMembershipDeleteService {
    s.teamId = &value
    return s
}

func (s *TeamUserMembershipDeleteService) UserId(value string) *TeamUserMembershipDeleteService {
    s.userId = &value
    return s
}

func (s *TeamUserMembershipDeleteService) Do(ctx context.Context) (TeamUserMembershipDeleteResponse, error) {
    var response TeamUserMembershipDeleteResponse

    if s.teamId == nil {
        return response, fmt.Errorf("missing required teamId")
    }

    if s.userId == nil {
        return response, fmt.Errorf("missing required userId")
    }

    url := fmt.Sprintf("%v/teams/%v/users/%v", s.c.baseURL, *s.teamId, *s.userId)
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
