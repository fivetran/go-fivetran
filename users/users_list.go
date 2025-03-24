package users

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type UsersListService struct {
	httputils.HttpService
	limit  *int
	cursor *string
}

func (s *UsersListService) Limit(value int) *UsersListService {
	s.limit = &value
	return s
}

func (s *UsersListService) Cursor(value string) *UsersListService {
	s.cursor = &value
	return s
}

func (s *UsersListService) Do(ctx context.Context) (UsersListResponse, error) {
	var response UsersListResponse
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
	err := s.HttpService.Do(ctx, "GET", "/users", nil, queries, 200, &response)
	return response, err
}
