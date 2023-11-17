package groups

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type GroupListUsersService struct {
	httputils.HttpService
	groupID *string
	limit   *int
	cursor  *string
}

func (s *GroupListUsersService) GroupID(value string) *GroupListUsersService {
	s.groupID = &value
	return s
}

func (s *GroupListUsersService) Limit(value int) *GroupListUsersService {
	s.limit = &value
	return s
}

func (s *GroupListUsersService) Cursor(value string) *GroupListUsersService {
	s.cursor = &value
	return s
}

func (s *GroupListUsersService) Do(ctx context.Context) (GroupListUsersResponse, error) {
	var response GroupListUsersResponse
	url := fmt.Sprintf("/groups/%v/users", *s.groupID)
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
