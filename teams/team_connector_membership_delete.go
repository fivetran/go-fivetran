package teams

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// TeamConnectorMembershipDeleteService implements the Team Management, Delete connector membership
// Ref. https://fivetran.com/docs/rest-api/teams#deleteconnectormembership
type TeamConnectorMembershipDeleteService struct {
	httputils.HttpService
	teamId      *string
	connectorId *string
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

	url := fmt.Sprintf("/teams/%v/connectors/%v", *s.teamId, *s.connectorId)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}