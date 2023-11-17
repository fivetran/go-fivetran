package groups

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type GroupAddUserService struct {
	httputils.HttpService
	groupID *string
	email   *string
	role    *string
}

func (s *GroupAddUserService) request() *groupAddUserRequest {
	return &groupAddUserRequest{
		Email: s.email,
		Role:  s.role,
	}
}

func (s *GroupAddUserService) GroupID(value string) *GroupAddUserService {
	s.groupID = &value
	return s
}

func (s *GroupAddUserService) Email(value string) *GroupAddUserService {
	s.email = &value
	return s
}

func (s *GroupAddUserService) Role(value string) *GroupAddUserService {
	s.role = &value
	return s
}

func (s *GroupAddUserService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	url := fmt.Sprintf("/groups/%v/users", *s.groupID)
	err := s.HttpService.Do(ctx, "POST", url, s.request(), nil, 200, &response)
	return response, err
}
