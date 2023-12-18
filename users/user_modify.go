package users

import (
	"context"
	"errors"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// UserModifyService implements the User Management, Modify a User API.
// Ref. https://fivetran.com/docs/rest-api/users#modifyauser
type UserModifyService struct {
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

func (s *UserModifyService) request() *userModifyRequest {
	result := userModifyRequest{
		GivenName:  s.givenName,
		FamilyName: s.familyName,
		Phone:      common.NewNullableString(s.phone, s.clearPhone),
		Picture:    common.NewNullableString(s.picture, s.clearPicture),
		Role:       s.role,
	}

	return &result
}

func (s *UserModifyService) UserID(value string) *UserModifyService {
	s.userID = &value
	return s
}

func (s *UserModifyService) GivenName(value string) *UserModifyService {
	s.givenName = &value
	return s
}

func (s *UserModifyService) FamilyName(value string) *UserModifyService {
	s.familyName = &value
	return s
}

func (s *UserModifyService) Phone(value string) *UserModifyService {
	s.phone = &value
	return s
}

func (s *UserModifyService) ClearPhone() *UserModifyService {
	s.clearPhone = true
	return s
}

func (s *UserModifyService) Picture(value string) *UserModifyService {
	s.picture = &value
	return s
}

func (s *UserModifyService) ClearPicture() *UserModifyService {
	s.clearPicture = true
	return s
}

func (s *UserModifyService) Role(value string) *UserModifyService {
	s.role = &value
	return s
}

func (s *UserModifyService) Do(ctx context.Context) (UserDetailsResponse, error) {
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
