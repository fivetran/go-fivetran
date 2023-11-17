package groups

import (
	"context"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type GroupCreateService struct {
	httputils.HttpService
	name *string
}

func (s *GroupCreateService) request() *groupCreateRequest {
	return &groupCreateRequest{
		Name: s.name,
	}
}

func (s *GroupCreateService) Name(value string) *GroupCreateService {
	s.name = &value
	return s
}

func (s *GroupCreateService) Do(ctx context.Context) (GroupDetailsResponse, error) {
	var response GroupDetailsResponse
	err := s.HttpService.Do(ctx, "POST", "/groups", s.request(), nil, 201, &response)
	return response, err
}
