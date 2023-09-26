package fivetran

import (
    "context"
    "encoding/json"
    "fmt"
)

// TeamConnectorsCreateService implements the Team Management, Add connector membership
// Ref. https://fivetran.com/docs/rest-api/teams#addconnectormembership
type TeamConnectorMembershipCreateService struct {
    c                 *Client
    teamId            *string
    connectorId       *string
    role              *string
}

type teamConnectorMembershipCreateRequest struct {
    ConnectorId       *string `json:"id,omitempty"`
    Role              *string `json:"role,omitempty"`
}

type TeamConnectorMembershipCreateResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Data    struct {
        ConnectorId     string       `json:"id"`
        Role            string       `json:"role"`
        CreatedAt       string       `json:"created_at"`
    } `json:"data"`
}

func (c *Client) NewTeamConnectorMembershipCreate() *TeamConnectorMembershipCreateService {
    return &TeamConnectorMembershipCreateService{c: c}
}

func (s *TeamConnectorMembershipCreateService) request() *teamConnectorMembershipCreateRequest {
    return &teamConnectorMembershipCreateRequest{
        ConnectorId:  s.connectorId,
        Role:         s.role,
    }
}

func (s *TeamConnectorMembershipCreateService) TeamId(value string) *TeamConnectorMembershipCreateService {
    s.teamId = &value
    return s
}

func (s *TeamConnectorMembershipCreateService) ConnectorId(value string) *TeamConnectorMembershipCreateService {
    s.connectorId = &value
    return s
}

func (s *TeamConnectorMembershipCreateService) Role(value string) *TeamConnectorMembershipCreateService {
    s.role = &value
    return s
}

func (s *TeamConnectorMembershipCreateService) Do(ctx context.Context) (TeamConnectorMembershipCreateResponse, error) {
    var response TeamConnectorMembershipCreateResponse

    if s.teamId == nil {
        return response, fmt.Errorf("missing required teamId")
    }

    url := fmt.Sprintf("%v/teams/%v/connectors", s.c.baseURL, *s.teamId)
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
