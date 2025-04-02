package users

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type UserDeleteService struct {
	httputils.HttpService
	userID *string
}

func (s *UserDeleteService) UserID(value string) *UserDeleteService {
	s.userID = &value
	return s
}

func (s *UserDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	
	if s.userID == nil {
		return response, fmt.Errorf("missing required userID")
	}

	url := fmt.Sprintf("/users/%v", *s.userID)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}