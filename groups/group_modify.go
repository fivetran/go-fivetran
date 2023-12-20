package groups

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type GroupModifyService struct {
	httputils.HttpService
	groupID *string
	name    *string
}

func (s *GroupModifyService) request() *groupModifyRequest {
	return &groupModifyRequest{
		Name: s.name,
	}
}

func (s *GroupModifyService) GroupID(value string) *GroupModifyService {
	s.groupID = &value
	return s
}

func (s *GroupModifyService) Name(value string) *GroupModifyService {
	s.name = &value
	return s
}

func (s *GroupModifyService) Do(ctx context.Context) (GroupDetailsResponse, error) {
	var response GroupDetailsResponse
	if s.groupID == nil {
		return response, fmt.Errorf("missing required groupID")
	}
	url := fmt.Sprintf("/groups/%v", *s.groupID)
	err := s.HttpService.Do(ctx, "PATCH", url, s.request(), nil, 200, &response)
	return response, err
}
