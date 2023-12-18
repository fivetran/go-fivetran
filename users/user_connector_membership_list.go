package users

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// UserConnectorMembershipsListService implements the User Management, List all connector memberships
// Ref. https://fivetran.com/docs/rest-api/users#listallconnectormemberships
type UserConnectorMembershipsListService struct {
	httputils.HttpService
	userId *string
	limit  *int
	cursor *string
}

func (s *UserConnectorMembershipsListService) UserId(value string) *UserConnectorMembershipsListService {
	s.userId = &value
	return s
}

func (s *UserConnectorMembershipsListService) Limit(value int) *UserConnectorMembershipsListService {
	s.limit = &value
	return s
}

func (s *UserConnectorMembershipsListService) Cursor(value string) *UserConnectorMembershipsListService {
	s.cursor = &value
	return s
}

func (s *UserConnectorMembershipsListService) Do(ctx context.Context) (UserConnectorMembershipsListResponse, error) {
	var response UserConnectorMembershipsListResponse

	if s.userId == nil {
		return response, fmt.Errorf("missing required userId")
	}

	url := fmt.Sprintf("/users/%v/connectors", *s.userId)
	var queries map[string]string = nil
	if s.cursor != nil || s.limit != nil {
		queries = make(map[string]string)
		if s.cursor != nil {
			queries["cursor"] = *s.cursor
		}
		if s.limit != nil {
			queries["limit"] = fmt.Sprintf("%v", *s.limit)
		}
	}
	err := s.HttpService.Do(ctx, "GET", url, nil, queries, 200, &response)
	return response, err
}