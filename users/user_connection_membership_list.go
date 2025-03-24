package users

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type UserConnectionMembershipsListService struct {
	httputils.HttpService
	userId *string
	limit  *int
	cursor *string
}

func (s *UserConnectionMembershipsListService) UserId(value string) *UserConnectionMembershipsListService {
	s.userId = &value
	return s
}

func (s *UserConnectionMembershipsListService) Limit(value int) *UserConnectionMembershipsListService {
	s.limit = &value
	return s
}

func (s *UserConnectionMembershipsListService) Cursor(value string) *UserConnectionMembershipsListService {
	s.cursor = &value
	return s
}

func (s *UserConnectionMembershipsListService) Do(ctx context.Context) (UserConnectionMembershipsListResponse, error) {
	var response UserConnectionMembershipsListResponse

	if s.userId == nil {
		return response, fmt.Errorf("missing required userId")
	}

	url := fmt.Sprintf("/users/%v/connections", *s.userId)
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