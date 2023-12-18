package teams

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// TeamUserMembershipsListService implements the Team Management, List all user memberships
// Ref. https://fivetran.com/docs/rest-api/teams#listallusermemberships
type TeamUserMembershipsListService struct {
	httputils.HttpService
	teamId *string
	limit  *int
	cursor *string
}

func (s *TeamUserMembershipsListService) TeamId(value string) *TeamUserMembershipsListService {
	s.teamId = &value
	return s
}

func (s *TeamUserMembershipsListService) Limit(value int) *TeamUserMembershipsListService {
	s.limit = &value
	return s
}

func (s *TeamUserMembershipsListService) Cursor(value string) *TeamUserMembershipsListService {
	s.cursor = &value
	return s
}

func (s *TeamUserMembershipsListService) Do(ctx context.Context) (TeamUserMembershipsListResponse, error) {
	var response TeamUserMembershipsListResponse

	if s.teamId == nil {
		return response, fmt.Errorf("missing required teamId")
	}

	url := fmt.Sprintf("/teams/%v/users", *s.teamId)
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