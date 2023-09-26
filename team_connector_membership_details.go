package fivetran

import (
    "context"
    "encoding/json"
    "fmt"
)

// TeamConnectorMembershipDetailsService implements the Team Management, Retrieve connector membership
// Ref. https://fivetran.com/docs/rest-api/teams#retrieveconnectormembership
type TeamConnectorMembershipDetailsService struct {
    c                 *Client
    teamId            *string
    connectorId       *string
}

type TeamConnectorMembershipDetailsResponse struct {
    Code    string `json:"code"`
    Data    struct {
        ConnectorId    string       `json:"id"`
        Role           string       `json:"role"`
        CreatedAt      string       `json:"created_at"`
    } `json:"data"`
}

func (c *Client) NewTeamConnectorMembershipDetails() *TeamConnectorMembershipDetailsService {
    return &TeamConnectorMembershipDetailsService{c: c}
}

func (s *TeamConnectorMembershipDetailsService) TeamId(value string) *TeamConnectorMembershipDetailsService {
    s.teamId = &value
    return s
}

func (s *TeamConnectorMembershipDetailsService) ConnectorId(value string) *TeamConnectorMembershipDetailsService {
    s.connectorId = &value
    return s
}

func (s *TeamConnectorMembershipDetailsService) Do(ctx context.Context) (TeamConnectorMembershipDetailsResponse, error) {
    var response TeamConnectorMembershipDetailsResponse

    if s.teamId == nil {
        return response, fmt.Errorf("missing required teamId")
    }

    if s.connectorId == nil {
        return response, fmt.Errorf("missing required connectorId")
    }

    url := fmt.Sprintf("%v/teams/%v/connectors/%v", s.c.baseURL, *s.teamId, *s.connectorId)
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
