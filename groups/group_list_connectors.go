package groups

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type GroupListConnectorsService struct {
	httputils.HttpService
	groupID *string
	limit   *int
	cursor  *string
	schema  *string
}

func (s *GroupListConnectorsService) GroupID(value string) *GroupListConnectorsService {
	s.groupID = &value
	return s
}

func (s *GroupListConnectorsService) Limit(value int) *GroupListConnectorsService {
	s.limit = &value
	return s
}

func (s *GroupListConnectorsService) Cursor(value string) *GroupListConnectorsService {
	s.cursor = &value
	return s
}

func (s *GroupListConnectorsService) Schema(value string) *GroupListConnectorsService {
	s.schema = &value
	return s
}

func (s *GroupListConnectorsService) Do(ctx context.Context) (GroupListConnectorsResponse, error) {
	var response GroupListConnectorsResponse
	url := fmt.Sprintf("/groups/%v/connectors", *s.groupID)
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
