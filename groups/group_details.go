package groups

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type GroupDetailsService struct {
	httputils.HttpService
	groupID *string
}

func (s *GroupDetailsService) GroupID(value string) *GroupDetailsService {
	s.groupID = &value
	return s
}

func (s *GroupDetailsService) Do(ctx context.Context) (GroupDetailsResponse, error) {
	var response GroupDetailsResponse
    if s.groupID == nil {
        return response, fmt.Errorf("missing required groupID")
    }

	url := fmt.Sprintf("/groups/%v", *s.groupID)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}
