package teams

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// TeamConnectorMembershipDetailsService implements the Team Management, Retrieve connector membership
// Ref. https://fivetran.com/docs/rest-api/teams#retrieveconnectormembership
type TeamConnectorMembershipDetailsService struct {
	httputils.HttpService
	teamId      *string
	connectorId *string
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

	url := fmt.Sprintf("/teams/%v/connectors/%v", *s.teamId, *s.connectorId)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}