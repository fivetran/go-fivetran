package users

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// UserConnectorMembershipModifyService implements the User Management, Update connector membership
// Ref. https://fivetran.com/docs/rest-api/users#updateconnectormembership
type UserConnectorMembershipModifyService struct {
	httputils.HttpService
	userId      *string
	connectorId *string
	role        *string
}

func (s *UserConnectorMembershipModifyService) request() *userConnectorMembershipModifyRequest {
	return &userConnectorMembershipModifyRequest{
		Role: s.role,
	}
}

func (s *UserConnectorMembershipModifyService) UserId(value string) *UserConnectorMembershipModifyService {
	s.userId = &value
	return s
}

func (s *UserConnectorMembershipModifyService) ConnectorId(value string) *UserConnectorMembershipModifyService {
	s.connectorId = &value
	return s
}

func (s *UserConnectorMembershipModifyService) Role(value string) *UserConnectorMembershipModifyService {
	s.role = &value
	return s
}

func (s *UserConnectorMembershipModifyService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse

	if s.userId == nil {
		return response, fmt.Errorf("missing required userId")
	}

	if s.connectorId == nil {
		return response, fmt.Errorf("missing required connectorId")
	}

	url := fmt.Sprintf("/users/%v/connectors/%v", *s.userId, *s.connectorId)
	err := s.HttpService.Do(ctx, "PATCH", url, s.request(), nil, 200, &response)
	return response, err
}
