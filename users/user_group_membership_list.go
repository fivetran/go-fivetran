package users

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// UserListGroupsService implements the User Management, List all group memberships
// Ref. https://fivetran.com/docs/rest-api/users#listallgroupmemberships
type UserGroupMembershipsListService struct {
	httputils.HttpService
	userId *string
	limit  *int
	cursor *string
}

func (s *UserGroupMembershipsListService) UserId(value string) *UserGroupMembershipsListService {
	s.userId = &value
	return s
}

func (s *UserGroupMembershipsListService) Limit(value int) *UserGroupMembershipsListService {
	s.limit = &value
	return s
}

func (s *UserGroupMembershipsListService) Cursor(value string) *UserGroupMembershipsListService {
	s.cursor = &value
	return s
}

func (s *UserGroupMembershipsListService) Do(ctx context.Context) (UserGroupMembershipsListResponse, error) {
	var response UserGroupMembershipsListResponse

	if s.userId == nil {
		return response, fmt.Errorf("missing required userId")
	}

	url := fmt.Sprintf("/users/%v/groups", *s.userId)
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