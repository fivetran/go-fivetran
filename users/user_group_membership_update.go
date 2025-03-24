package users

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type UserGroupMembershipUpdateService struct {
	httputils.HttpService
	userId  *string
	groupId *string
	role    *string
}

func (s *UserGroupMembershipUpdateService) request() *userGroupMembershipUpdateRequest {
	return &userGroupMembershipUpdateRequest{
		Role: s.role,
	}
}

func (s *UserGroupMembershipUpdateService) UserId(value string) *UserGroupMembershipUpdateService {
	s.userId = &value
	return s
}

func (s *UserGroupMembershipUpdateService) GroupId(value string) *UserGroupMembershipUpdateService {
	s.groupId = &value
	return s
}

func (s *UserGroupMembershipUpdateService) Role(value string) *UserGroupMembershipUpdateService {
	s.role = &value
	return s
}

func (s *UserGroupMembershipUpdateService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse

	if s.userId == nil {
		return response, fmt.Errorf("missing required userId")
	}

	if s.groupId == nil {
		return response, fmt.Errorf("missing required groupId")
	}

	url := fmt.Sprintf("/users/%v/groups/%v", *s.userId, *s.groupId)
	err := s.HttpService.Do(ctx, "PATCH", url, s.request(), nil, 200, &response)
	return response, err
}
