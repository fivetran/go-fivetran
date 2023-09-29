package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
)

// TeamConnectorMembershipModifyService implements the Team Management, Update connector membership
// Ref. https://fivetran.com/docs/rest-api/teams#updateconnectormembership
type TeamConnectorMembershipModifyService struct {
	c           *Client
	teamId      *string
	connectorId *string
	role        *string
}

type teamConnectorMembershipModifyRequest struct {
	Role *string `json:"role,omitempty"`
}

func (c *Client) NewTeamConnectorMembershipModify() *TeamConnectorMembershipModifyService {
	return &TeamConnectorMembershipModifyService{c: c}
}

func (s *TeamConnectorMembershipModifyService) request() *teamConnectorMembershipModifyRequest {
	return &teamConnectorMembershipModifyRequest{
		Role: s.role,
	}
}

func (s *TeamConnectorMembershipModifyService) TeamId(value string) *TeamConnectorMembershipModifyService {
	s.teamId = &value
	return s
}

func (s *TeamConnectorMembershipModifyService) ConnectorId(value string) *TeamConnectorMembershipModifyService {
	s.connectorId = &value
	return s
}

func (s *TeamConnectorMembershipModifyService) Role(value string) *TeamConnectorMembershipModifyService {
	s.role = &value
	return s
}

func (s *TeamConnectorMembershipModifyService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse

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
