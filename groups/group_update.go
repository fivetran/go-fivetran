package groups

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type GroupUpdateService struct {
	httputils.HttpService
	groupID *string
	name    *string
}

func (s *GroupUpdateService) request() *groupUpdateRequest {
	return &groupUpdateRequest{
		Name: s.name,
	}
}

func (s *GroupUpdateService) GroupID(value string) *GroupUpdateService {
	s.groupID = &value
	return s
}

func (s *GroupUpdateService) Name(value string) *GroupUpdateService {
	s.name = &value
	return s
}

func (s *GroupUpdateService) Do(ctx context.Context) (GroupDetailsResponse, error) {
	var response GroupDetailsResponse
	if s.groupID == nil {
		return response, fmt.Errorf("missing required groupID")
	}
	url := fmt.Sprintf("/groups/%v", *s.groupID)
	err := s.HttpService.Do(ctx, "PATCH", url, s.request(), nil, 200, &response)
	return response, err
}
