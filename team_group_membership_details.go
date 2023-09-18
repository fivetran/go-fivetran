package fivetran

import (
    "context"
    "encoding/json"
    "fmt"
)

// TeamGroupMembershipDetailsService implements the Team Management, Retrieve group membership
// Ref. https://fivetran.com/docs/rest-api/teams#retrievegroupmembership
type TeamGroupMembershipDetailsService struct {
    c                 *Client
    teamId            *string
    groupId           *string
}

type TeamGroupMembershipDetailsResponse struct {
    Code    string `json:"code"`
    Data    struct {
        GroupId    string       `json:"id"`
        Role       string       `json:"role"`
        CreatedAt  string       `json:"created_at"`
    } `json:"data"`
}

func (c *Client) NewTeamGroupMembershipDetails() *TeamGroupMembershipDetailsService {
    return &TeamGroupMembershipDetailsService{c: c}
}

func (s *TeamGroupMembershipDetailsService) TeamId(value string) *TeamGroupMembershipDetailsService {
    s.teamId = &value
    return s
}

func (s *TeamGroupMembershipDetailsService) GroupId(value string) *TeamGroupMembershipDetailsService {
    s.groupId = &value
    return s
}

func (s *TeamGroupMembershipDetailsService) Do(ctx context.Context) (TeamGroupMembershipDetailsResponse, error) {
    var response TeamGroupMembershipDetailsResponse

    if s.teamId == nil {
        return response, fmt.Errorf("missing required teamId")
    }

    if s.groupId == nil {
        return response, fmt.Errorf("missing required groupId")
    }

    url := fmt.Sprintf("%v/teams/%v/groups/%v", s.c.baseURL, *s.teamId, *s.groupId)
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
