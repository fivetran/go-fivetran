package users

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// UserConnectorMembershipDeleteService implements the User Management, Delete connector membership
// Ref. https://fivetran.com/docs/rest-api/users#deleteconnectormembership
type UserConnectorMembershipDeleteService struct {
	httputils.HttpService
	userId      *string
	connectorId *string
}

func (s *UserConnectorMembershipDeleteService) UserId(value string) *UserConnectorMembershipDeleteService {
	s.userId = &value
	return s
}

func (s *UserConnectorMembershipDeleteService) ConnectorId(value string) *UserConnectorMembershipDeleteService {
	s.connectorId = &value
	return s
}

func (s *UserConnectorMembershipDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	
	if s.userId == nil {
		return response, fmt.Errorf("missing required userId")
	}

	if s.connectorId == nil {
		return response, fmt.Errorf("missing required connectorId")
	}

	url := fmt.Sprintf("/users/%v/connectors/%v", *s.userId, *s.connectorId)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}
