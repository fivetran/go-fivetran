package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// UserInviteService implements the User Management, Invite a User API.
// Ref. https://fivetran.com/docs/rest-api/users#inviteauser
type UserInviteService struct {
	c          *Client
	email      *string
	givenName  *string
	familyName *string
	phone      *string
	picture    *string
	role       *string
}

type userInviteRequest struct {
	Email      *string `json:"email,omitempty"`
	GivenName  *string `json:"given_name,omitempty"`
	FamilyName *string `json:"family_name,omitempty"`
	Phone      *string `json:"phone,omitempty"`
	Picture    *string `json:"picture,omitempty"`
	Role       *string `json:"role,omitempty"`
}

type UserInviteResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID         string    `json:"id"`
		Email      string    `json:"email"`
		GivenName  string    `json:"given_name"`
		FamilyName string    `json:"family_name"`
		Verified   bool      `json:"verified"`
		Invited    bool      `json:"invited"`
		Picture    string    `json:"picture"`
		Phone      string    `json:"phone"`
		LoggedInAt time.Time `json:"logged_in_at"`
		CreatedAt  time.Time `json:"created_at"`
	} `json:"data"`
}

func (c *Client) NewUserInvite() *UserInviteService {
	return &UserInviteService{c: c}
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

func (s *UserInviteService) Do(ctx context.Context) (UserInviteResponse, error) {
	var response UserInviteResponse
	url := fmt.Sprintf("%v/users", s.c.baseURL)
	expectedStatus := 201

	headers := make(map[string]string)
	headers["Authorization"] = s.c.authorization
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := request{
		method:  "POST",
		url:     url,
		body:    reqBody,
		queries: nil,
		headers: headers,
	}

	respBody, respStatus, err := httpRequest(r, ctx)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return response, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return response, err
	}

	return response, nil
}
