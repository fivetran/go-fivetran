package groups

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type GroupRemoveUserService struct {
	httputils.HttpService
	groupID *string
	userID  *string
}

func (s *GroupRemoveUserService) GroupID(value string) *GroupRemoveUserService {
	s.groupID = &value
	return s
}

func (s *GroupRemoveUserService) UserID(value string) *GroupRemoveUserService {
	s.userID = &value
	return s
}

func (s *GroupRemoveUserService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	url := fmt.Sprintf("/groups/%v/users/%v", *s.groupID, *s.userID)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}
