package groups

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type GroupListPrivateLinksService struct {
	httputils.HttpService
	groupID *string
	limit   *int
	cursor  *string
}

func (s *GroupListPrivateLinksService) GroupID(value string) *GroupListPrivateLinksService {
	s.groupID = &value
	return s
}

func (s *GroupListPrivateLinksService) Limit(value int) *GroupListPrivateLinksService {
	s.limit = &value
	return s
}

func (s *GroupListPrivateLinksService) Cursor(value string) *GroupListPrivateLinksService {
	s.cursor = &value
	return s
}

func (s *GroupListPrivateLinksService) Do(ctx context.Context) (GroupListPrivateLinksResponse, error) {
	var response GroupListPrivateLinksResponse
	url := fmt.Sprintf("/groups/%v/private-links", *s.groupID)
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
