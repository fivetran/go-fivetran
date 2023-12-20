package users

import (
	"context"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// UserInviteService implements the User Management, Invite a User API.
// Ref. https://fivetran.com/docs/rest-api/users#inviteauser
type UserInviteService struct {
	httputils.HttpService
	email      *string
	givenName  *string
	familyName *string
	phone      *string
	picture    *string
	role       *string
}

func (s *UserInviteService) request() *userInviteRequest {
	return &userInviteRequest{
		Email:      s.email,
		GivenName:  s.givenName,
		FamilyName: s.familyName,
		Phone:      s.phone,
		Picture:    s.picture,
		Role:       s.role,
	}
}

func (s *UserInviteService) Email(value string) *UserInviteService {
	s.email = &value
	return s
}

func (s *UserInviteService) GivenName(value string) *UserInviteService {
	s.givenName = &value
	return s
}

func (s *UserInviteService) FamilyName(value string) *UserInviteService {
	s.familyName = &value
	return s
}

func (s *UserInviteService) Phone(value string) *UserInviteService {
	s.phone = &value
	return s
}

func (s *UserInviteService) Picture(value string) *UserInviteService {
	s.picture = &value
	return s
}

func (s *UserInviteService) Role(value string) *UserInviteService {
	s.role = &value
	return s
}

func (s *UserInviteService) Do(ctx context.Context) (UserDetailsResponse, error) {
	var response UserDetailsResponse
	err := s.HttpService.Do(ctx, "POST", "/users", s.request(), nil, 201, &response)
	return response, err
}