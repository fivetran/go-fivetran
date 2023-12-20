package teams

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// TeamConnectorMembershipsListService implements the Team Management, List all connector memberships
// Ref. https://fivetran.com/docs/rest-api/teams#listallconnectormemberships
type TeamConnectorMembershipsListService struct {
	httputils.HttpService
	teamId *string
	limit  *int
	cursor *string
}

func (s *TeamConnectorMembershipsListService) TeamId(value string) *TeamConnectorMembershipsListService {
	s.teamId = &value
	return s
}

func (s *TeamConnectorMembershipsListService) Limit(value int) *TeamConnectorMembershipsListService {
	s.limit = &value
	return s
}

func (s *TeamConnectorMembershipsListService) Cursor(value string) *TeamConnectorMembershipsListService {
	s.cursor = &value
	return s
}

func (s *TeamConnectorMembershipsListService) Do(ctx context.Context) (TeamConnectorMembershipsListResponse, error) {
	var response TeamConnectorMembershipsListResponse

	if s.teamId == nil {
		return response, fmt.Errorf("missing required teamId")
	}

	url := fmt.Sprintf("/teams/%v/connectors", *s.teamId)
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