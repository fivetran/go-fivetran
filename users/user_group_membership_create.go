package users

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type UserGroupMembershipCreateService struct {
	httputils.HttpService
	userId  *string
	groupId *string
	role    *string
}

func (s *UserGroupMembershipCreateService) request() *userGroupMembershipCreateRequest {
	return &userGroupMembershipCreateRequest{
		GroupId: s.groupId,
		Role:    s.role,
	}
}

func (s *UserGroupMembershipCreateService) UserId(value string) *UserGroupMembershipCreateService {
	s.userId = &value
	return s
}

func (s *UserGroupMembershipCreateService) GroupId(value string) *UserGroupMembershipCreateService {
	s.groupId = &value
	return s
}

func (s *UserGroupMembershipCreateService) Role(value string) *UserGroupMembershipCreateService {
	s.role = &value
	return s
}

func (s *UserGroupMembershipCreateService) Do(ctx context.Context) (UserGroupMembershipCreateResponse, error) {
	var response UserGroupMembershipCreateResponse
	url := fmt.Sprintf("/users/%v/groups", *s.userId)
	err := s.HttpService.Do(ctx, "POST", url, s.request(), nil, 201, &response)
	return response, err
}
