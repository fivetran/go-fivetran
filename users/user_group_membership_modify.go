package users

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// UserGroupMembershipModifyService implements the User Management, Update group membership
// Ref. https://fivetran.com/docs/rest-api/users#updategroupmembership
type UserGroupMembershipModifyService struct {
	httputils.HttpService
	userId  *string
	groupId *string
	role    *string
}

func (s *UserGroupMembershipModifyService) request() *userGroupMembershipModifyRequest {
	return &userGroupMembershipModifyRequest{
		Role: s.role,
	}
}

func (s *UserGroupMembershipModifyService) UserId(value string) *UserGroupMembershipModifyService {
	s.userId = &value
	return s
}

func (s *UserGroupMembershipModifyService) GroupId(value string) *UserGroupMembershipModifyService {
	s.groupId = &value
	return s
}

func (s *UserGroupMembershipModifyService) Role(value string) *UserGroupMembershipModifyService {
	s.role = &value
	return s
}

func (s *UserGroupMembershipModifyService) Do(ctx context.Context) (common.CommonResponse, error) {
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
