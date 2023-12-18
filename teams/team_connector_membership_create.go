package teams

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// TeamConnectorsCreateService implements the Team Management, Add connector membership
// Ref. https://fivetran.com/docs/rest-api/teams#addconnectormembership
type TeamConnectorMembershipCreateService struct {
	httputils.HttpService
	teamId      *string
	connectorId *string
	role        *string
}

func (s *TeamConnectorMembershipCreateService) request() *teamConnectorMembershipCreateRequest {
	return &teamConnectorMembershipCreateRequest{
		ConnectorId: s.connectorId,
		Role:        s.role,
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
	url := fmt.Sprintf("/teams/%v/connectors", *s.teamId)
	err := s.HttpService.Do(ctx, "POST", url, s.request(), nil, 201, &response)
	return response, err
}
