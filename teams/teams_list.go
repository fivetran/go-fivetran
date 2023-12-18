package teams

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// TeamsListService implements the Team Management, retrieve List all Teams.
// Ref. https://fivetran.com/docs/rest-api/teams#listallteams
type TeamsListService struct {
	httputils.HttpService
	limit  *int
	cursor *string
}

func (s *TeamsListService) Limit(value int) *TeamsListService {
	s.limit = &value
	return s
}

func (s *TeamsListService) Cursor(value string) *TeamsListService {
	s.cursor = &value
	return s
}

func (s *TeamsListService) Do(ctx context.Context) (TeamsListResponse, error) {
	var response TeamsListResponse

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
	err := s.HttpService.Do(ctx, "GET", "/teams", nil, queries, 200, &response)
	return response, err
}