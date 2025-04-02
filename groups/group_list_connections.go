package groups

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type GroupListConnectionsService struct {
	httputils.HttpService
	groupID *string
	limit   *int
	cursor  *string
	schema  *string
}

func (s *GroupListConnectionsService) GroupID(value string) *GroupListConnectionsService {
	s.groupID = &value
	return s
}

func (s *GroupListConnectionsService) Limit(value int) *GroupListConnectionsService {
	s.limit = &value
	return s
}

func (s *GroupListConnectionsService) Cursor(value string) *GroupListConnectionsService {
	s.cursor = &value
	return s
}

func (s *GroupListConnectionsService) Schema(value string) *GroupListConnectionsService {
	s.schema = &value
	return s
}

func (s *GroupListConnectionsService) Do(ctx context.Context) (GroupListConnectionsResponse, error) {
	var response GroupListConnectionsResponse
	url := fmt.Sprintf("/groups/%v/connections", *s.groupID)
	var queries map[string]string = nil
	if s.cursor != nil || s.limit != nil || s.schema != nil {
		queries = make(map[string]string)
		if s.cursor != nil {
			queries["cursor"] = *s.cursor
		}
		if s.limit != nil {
			queries["limit"] = fmt.Sprintf("%v", *s.limit)
		}
		if s.schema != nil {
			queries["schema"] = *s.schema
		}
	}
	err := s.HttpService.Do(ctx, "GET", url, nil, queries, 200, &response)
	return response, err
}
