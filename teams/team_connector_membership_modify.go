package teams

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// TeamConnectorMembershipModifyService implements the Team Management, Update connector membership
// Ref. https://fivetran.com/docs/rest-api/teams#updateconnectormembership
type TeamConnectorMembershipModifyService struct {
	httputils.HttpService
	teamId      *string
	connectorId *string
	role        *string
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

	url := fmt.Sprintf("/teams/%v/connectors/%v", *s.teamId, *s.connectorId)
	err := s.HttpService.Do(ctx, "PATCH", url, s.request(), nil, 200, &response)
	return response, err
}