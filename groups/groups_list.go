package groups

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type GroupsListService struct {
	httputils.HttpService
	limit  *int
	cursor *string
}

func (s *GroupsListService) Limit(value int) *GroupsListService {
	s.limit = &value
	return s
}

func (s *GroupsListService) Cursor(value string) *GroupsListService {
	s.cursor = &value
	return s
}

func (s *GroupsListService) Do(ctx context.Context) (GroupsListResponse, error) {
	var response GroupsListResponse
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
	err := s.HttpService.Do(ctx, "GET", "/groups", nil, queries, 200, &response)
	return response, err
}
