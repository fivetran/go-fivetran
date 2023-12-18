package users

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// UserGroupMembershipDeleteService implements the User Management, Delete group membership
// Ref. https://fivetran.com/docs/rest-api/users#deletegroupmembership
type UserGroupMembershipDeleteService struct {
	httputils.HttpService
	userId  *string
	groupId *string
}

func (s *UserGroupMembershipDeleteService) UserId(value string) *UserGroupMembershipDeleteService {
	s.userId = &value
	return s
}

func (s *UserGroupMembershipDeleteService) GroupId(value string) *UserGroupMembershipDeleteService {
	s.groupId = &value
	return s
}

func (s *UserGroupMembershipDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	
	if s.userId == nil {
		return response, fmt.Errorf("missing required userId")
	}

	if s.groupId == nil {
		return response, fmt.Errorf("missing required groupId")
	}

	url := fmt.Sprintf("/users/%v/groups/%v", *s.userId, *s.groupId)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}
