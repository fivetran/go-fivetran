package privatelink

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// PrivateLinkListService Returns a list of all private links within your Fivetran account.
// Ref. https://fivetran.com/docs/rest-api/private-links-management#listallprivatelinkswithinaccount
type PrivateLinkListService struct {
	httputils.HttpService
	limit  *int
	cursor *string
}

func (s *PrivateLinkListService) Limit(value int) *PrivateLinkListService {
	s.limit = &value
	return s
}

func (s *PrivateLinkListService) Cursor(value string) *PrivateLinkListService {
	s.cursor = &value
	return s
}

func (s *PrivateLinkListService) Do(ctx context.Context) (PrivateLinkListResponse, error) {
	var response PrivateLinkListResponse

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
	err := s.HttpService.Do(ctx, "GET", "/private-links", nil, queries, 200, &response)
	return response, err
}