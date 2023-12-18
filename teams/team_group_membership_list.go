package teams

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// TeamListGroupsService implements the Team Management, List all group memberships
// Ref. https://fivetran.com/docs/rest-api/teams#listallgroupmemberships
type TeamGroupMembershipsListService struct {
	httputils.HttpService
	teamId *string
	limit  *int
	cursor *string
}

func (s *TeamGroupMembershipsListService) TeamId(value string) *TeamGroupMembershipsListService {
	s.teamId = &value
	return s
}

func (s *TeamGroupMembershipsListService) Limit(value int) *TeamGroupMembershipsListService {
	s.limit = &value
	return s
}

func (s *TeamGroupMembershipsListService) Cursor(value string) *TeamGroupMembershipsListService {
	s.cursor = &value
	return s
}

func (s *TeamGroupMembershipsListService) Do(ctx context.Context) (TeamGroupMembershipsListResponse, error) {
	var response TeamGroupMembershipsListResponse

	if s.teamId == nil {
		return response, fmt.Errorf("missing required teamId")
	}

	url := fmt.Sprintf("/teams/%v/groups", *s.teamId)
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