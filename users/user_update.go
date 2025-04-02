package users

import (
	"context"
	"errors"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type UserUpdateService struct {
	httputils.HttpService
	userID       *string
	givenName    *string
	familyName   *string
	phone        *string
	picture      *string
	role         *string
	clearPicture bool
	clearPhone   bool
}

func (s *UserUpdateService) request() *userUpdateRequest {
	result := userUpdateRequest{
		GivenName:  s.givenName,
		FamilyName: s.familyName,
		Phone:      common.NewNullableString(s.phone, s.clearPhone),
		Picture:    common.NewNullableString(s.picture, s.clearPicture),
		Role:       s.role,
	}

	return &result
}

func (s *UserUpdateService) UserID(value string) *UserUpdateService {
	s.userID = &value
	return s
}

func (s *UserUpdateService) GivenName(value string) *UserUpdateService {
	s.givenName = &value
	return s
}

func (s *UserUpdateService) FamilyName(value string) *UserUpdateService {
	s.familyName = &value
	return s
}

func (s *UserUpdateService) Phone(value string) *UserUpdateService {
	s.phone = &value
	return s
}

func (s *UserUpdateService) ClearPhone() *UserUpdateService {
	s.clearPhone = true
	return s
}

func (s *UserUpdateService) Picture(value string) *UserUpdateService {
	s.picture = &value
	return s
}

func (s *UserUpdateService) ClearPicture() *UserUpdateService {
	s.clearPicture = true
	return s
}

func (s *UserUpdateService) Role(value string) *UserUpdateService {
	s.role = &value
	return s
}

func (s *UserUpdateService) Do(ctx context.Context) (UserDetailsResponse, error) {
	var response UserDetailsResponse
	if s.userID == nil {
		return response, fmt.Errorf("missing required UserID")
	}

	if s.clearPhone && s.phone != nil {
		return response, errors.New("can't 'set phone' and 'clear phone' in one request")
	}

	if s.clearPicture && s.picture != nil {
		return response, errors.New("can't 'set picture' and 'clear picture' in one request")
	}

	url := fmt.Sprintf("/users/%v", *s.userID)
	err := s.HttpService.Do(ctx, "PATCH", url, s.request(), nil, 200, &response)
	return response, err
}
