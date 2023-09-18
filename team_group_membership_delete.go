package fivetran

import (
    "context"
    "encoding/json"
    "fmt"
)

// TeamGroupMembershipDeleteService implements the Team Management, Delete group membership
// Ref. https://fivetran.com/docs/rest-api/teams#deletegroupmembership
type TeamGroupMembershipDeleteService struct {
    c                *Client
    teamId           *string
    groupId          *string
}

type TeamGroupMembershipDeleteResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
}

func (c *Client) NewTeamGroupMembershipDelete() *TeamGroupMembershipDeleteService {
    return &TeamGroupMembershipDeleteService{c: c}
}

func (s *TeamGroupMembershipDeleteService) TeamId(value string) *TeamGroupMembershipDeleteService {
    s.teamId = &value
    return s
}

func (s *TeamGroupMembershipDeleteService) GroupId(value string) *TeamGroupMembershipDeleteService {
    s.groupId = &value
    return s
}

func (s *TeamGroupMembershipDeleteService) Do(ctx context.Context) (TeamGroupMembershipDeleteResponse, error) {
    var response TeamGroupMembershipDeleteResponse

    if s.teamId == nil {
        return response, fmt.Errorf("missing required teamId")
    }

    if s.groupId == nil {
        return response, fmt.Errorf("missing required groupId")
    }

    url := fmt.Sprintf("%v/teams/%v/groups/%v", s.c.baseURL, *s.teamId, *s.groupId)
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
