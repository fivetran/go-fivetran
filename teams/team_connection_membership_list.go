package teams

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type TeamConnectionMembershipsListService struct {
	httputils.HttpService
	teamId *string
	limit  *int
	cursor *string
}

func (s *TeamConnectionMembershipsListService) TeamId(value string) *TeamConnectionMembershipsListService {
	s.teamId = &value
	return s
}

func (s *TeamConnectionMembershipsListService) Limit(value int) *TeamConnectionMembershipsListService {
	s.limit = &value
	return s
}

func (s *TeamConnectionMembershipsListService) Cursor(value string) *TeamConnectionMembershipsListService {
	s.cursor = &value
	return s
}

func (s *TeamConnectionMembershipsListService) Do(ctx context.Context) (TeamConnectionMembershipsListResponse, error) {
	var response TeamConnectionMembershipsListResponse

	if s.teamId == nil {
		return response, fmt.Errorf("missing required teamId")
	}

	url := fmt.Sprintf("/teams/%v/connections", *s.teamId)
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