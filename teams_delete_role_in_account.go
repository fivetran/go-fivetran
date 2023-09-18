package fivetran

import (
    "context"
    "encoding/json"
    "fmt"
)

// TeamRoleDeleteService implements the Team Management, Delete team role in account
// Ref. https://fivetran.com/docs/rest-api/teams#deleteteamroleinaccount
type TeamsDeleteRoleInAccountService struct {
    c                *Client
    teamId           *string
}

type TeamsDeleteRoleInAccountResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
}

func (c *Client) NewTeamsDeleteRoleInAccount() *TeamsDeleteRoleInAccountService {
    return &TeamsDeleteRoleInAccountService{c: c}
}

func (s *TeamsDeleteRoleInAccountService) TeamId(value string) *TeamsDeleteRoleInAccountService {
    s.teamId = &value
    return s
}

func (s *TeamsDeleteRoleInAccountService) Do(ctx context.Context) (TeamsDeleteRoleInAccountResponse, error) {
    var response TeamsDeleteRoleInAccountResponse

    if s.teamId == nil {
        return response, fmt.Errorf("missing required teamId")
    }

    url := fmt.Sprintf("%v/teams/%v/role", s.c.baseURL, *s.teamId)
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
