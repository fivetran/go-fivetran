package users

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// UserConnectorMembershipDetailsService implements the User Management, Retrieve connector membership
// Ref. https://fivetran.com/docs/rest-api/users#retrieveconnectormembership
type UserConnectorMembershipDetailsService struct {
	httputils.HttpService
	userId      *string
	connectorId *string
}

func (s *UserConnectorMembershipDetailsService) UserId(value string) *UserConnectorMembershipDetailsService {
	s.userId = &value
	return s
}

func (s *UserConnectorMembershipDetailsService) ConnectorId(value string) *UserConnectorMembershipDetailsService {
	s.connectorId = &value
	return s
}

func (s *UserConnectorMembershipDetailsService) Do(ctx context.Context) (UserConnectorMembershipDetailsResponse, error) {
	var response UserConnectorMembershipDetailsResponse

	if s.userId == nil {
		return response, fmt.Errorf("missing required userId")
	}

	if s.connectorId == nil {
		return response, fmt.Errorf("missing required connectorId")
	}

	url := fmt.Sprintf("/users/%v/connectors/%v", *s.userId, *s.connectorId)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}
