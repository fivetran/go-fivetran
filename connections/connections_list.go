package connections

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type ConnectionsListService struct {
	httputils.HttpService
	groupID *string
	schema  *string
	limit   *int
	cursor  *string
}

func (s *ConnectionsListService) GroupID(value string) *ConnectionsListService {
	s.groupID = &value
	return s
}

func (s *ConnectionsListService) Schema(value string) *ConnectionsListService {
	s.schema = &value
	return s
}

func (s *ConnectionsListService) Limit(value int) *ConnectionsListService {
	s.limit = &value
	return s
}

func (s *ConnectionsListService) Cursor(value string) *ConnectionsListService {
	s.cursor = &value
	return s
}

func (s *ConnectionsListService) Do(ctx context.Context) (ConnectionsListResponse, error) {
	var response ConnectionsListResponse
	var queries map[string]string = nil
	if s.groupID != nil || s.schema != nil || s.cursor != nil || s.limit != nil {
		queries = make(map[string]string)
		if s.groupID != nil {
			queries["group_id"] = *s.groupID
		}
		if s.schema != nil {
			queries["schema"] = *s.schema
		}
		if s.cursor != nil {
			queries["cursor"] = *s.cursor
		}
		if s.limit != nil {
			queries["limit"] = fmt.Sprintf("%v", *s.limit)
		}
	}
	err := s.HttpService.Do(ctx, "GET", "/connections", nil, queries, 200, &response)
	return response, err
}
