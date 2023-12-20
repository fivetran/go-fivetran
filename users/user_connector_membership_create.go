package users

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// UserConnectorsCreateService implements the User Management, Add connector membership
// Ref. https://fivetran.com/docs/rest-api/users#addconnectormembership
type UserConnectorMembershipCreateService struct {
	httputils.HttpService
	userId      *string
	connectorId *string
	role        *string
}

func (s *UserConnectorMembershipCreateService) request() *userConnectorMembershipCreateRequest {
	return &userConnectorMembershipCreateRequest{
		ConnectorId: s.connectorId,
		Role:        s.role,
	}
}

func (s *UserConnectorMembershipCreateService) UserId(value string) *UserConnectorMembershipCreateService {
	s.userId = &value
	return s
}

func (s *UserConnectorMembershipCreateService) ConnectorId(value string) *UserConnectorMembershipCreateService {
	s.connectorId = &value
	return s
}

func (s *UserConnectorMembershipCreateService) Role(value string) *UserConnectorMembershipCreateService {
	s.role = &value
	return s
}

func (s *UserConnectorMembershipCreateService) Do(ctx context.Context) (UserConnectorMembershipCreateResponse, error) {
	var response UserConnectorMembershipCreateResponse
	url := fmt.Sprintf("/users/%v/connectors", *s.userId)
	err := s.HttpService.Do(ctx, "POST", url, s.request(), nil, 201, &response)
	return response, err
}