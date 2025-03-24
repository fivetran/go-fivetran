package users

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type UserDetailsService struct {
	httputils.HttpService
	userID *string
}

func (s *UserDetailsService) UserID(value string) *UserDetailsService {
	s.userID = &value
	return s
}

func (s *UserDetailsService) Do(ctx context.Context) (UserDetailsResponse, error) {
	var response UserDetailsResponse
	if s.userID == nil {
		return response, fmt.Errorf("missing required userID")
	}
	url := fmt.Sprintf("/users/%v", *s.userID)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}