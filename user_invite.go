package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type userInviteService struct {
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

func (c *Client) NewUserInvite() *userInviteService {
	return &userInviteService{c: c}
}

func (s *userInviteService) request() userInviteRequest {
	return userInviteRequest{
		Email:      s.email,
		GivenName:  s.givenName,
		FamilyName: s.familyName,
		Phone:      s.phone,
		Picture:    s.picture,
		Role:       s.role,
	}
}

func (s *userInviteService) Email(value string) *userInviteService {
	s.email = &value
	return s
}

func (s *userInviteService) GivenName(value string) *userInviteService {
	s.givenName = &value
	return s
}

func (s *userInviteService) FamilyName(value string) *userInviteService {
	s.familyName = &value
	return s
}

func (s *userInviteService) Phone(value string) *userInviteService {
	s.phone = &value
	return s
}

func (s *userInviteService) Picture(value string) *userInviteService {
	s.picture = &value
	return s
}

func (s *userInviteService) Role(value string) *userInviteService {
	s.role = &value
	return s
}

func (s *userInviteService) Do(ctx context.Context) (UserInviteResponse, error) {
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

	r := Request{
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
