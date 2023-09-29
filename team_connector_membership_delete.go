package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// TeamConnectorMembershipDeleteService implements the Team Management, Delete connector membership
// Ref. https://fivetran.com/docs/rest-api/teams#deleteconnectormembership
type TeamConnectorMembershipDeleteService struct {
	c           *Client
	teamId      *string
	connectorId *string
}

func (c *Client) NewTeamConnectorMembershipDelete() *TeamConnectorMembershipDeleteService {
	return &TeamConnectorMembershipDeleteService{c: c}
}

func (s *TeamConnectorMembershipDeleteService) TeamId(value string) *TeamConnectorMembershipDeleteService {
	s.teamId = &value
	return s
}

func (s *TeamConnectorMembershipDeleteService) ConnectorId(value string) *TeamConnectorMembershipDeleteService {
	s.connectorId = &value
	return s
}

func (s *TeamConnectorMembershipDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
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

	r := httputils.Request{
		Method:           "DELETE",
		Url:              url,
		Body:             nil,
		Queries:          nil,
		Headers:          headers,
		Client:           s.c.httpClient,
		HandleRateLimits: s.c.handleRateLimits,
		MaxRetryAttempts: s.c.maxRetryAttempts,
	}

	respBody, respStatus, err := r.Do(ctx)
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
