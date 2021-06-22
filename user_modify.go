package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type userModifyService struct {
	c          *Client
	userID     *string
	givenName  *string
	familyName *string
	phone      *string
	picture    *string
	role       *string
}

type userModifyRequest struct {
	GivenName  *string `json:"given_name,omitempty"`
	FamilyName *string `json:"family_name,omitempty"`
	Phone      *string `json:"phone,omitempty"`
	Picture    *string `json:"picture,omitempty"`
	Role       *string `json:"role,omitempty"`
}

type UserModifyResponse struct {
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

func (c *Client) NewUserModify() *userModifyService {
	return &userModifyService{c: c}
}

func (s *userModifyService) request() userModifyRequest {
	return userModifyRequest{
		GivenName:  s.givenName,
		FamilyName: s.familyName,
		Phone:      s.phone,
		Picture:    s.picture,
		Role:       s.role,
	}
}

func (s *userModifyService) UserID(value string) *userModifyService {
	s.userID = &value
	return s
}

func (s *userModifyService) GivenName(value string) *userModifyService {
	s.givenName = &value
	return s
}

func (s *userModifyService) FamilyName(value string) *userModifyService {
	s.familyName = &value
	return s
}

func (s *userModifyService) Phone(value string) *userModifyService {
	s.phone = &value
	return s
}

func (s *userModifyService) Picture(value string) *userModifyService {
	s.picture = &value
	return s
}

func (s *userModifyService) Role(value string) *userModifyService {
	s.role = &value
	return s
}

func (s *userModifyService) Do(ctx context.Context) (UserModifyResponse, error) {
	var response UserModifyResponse

	if s.userID == nil {
		return response, fmt.Errorf("missing required UserID")
	}

	url := fmt.Sprintf("%v/users/%v", s.c.baseURL, *s.userID)
	expectedStatus := 200

	headers := make(map[string]string)
	headers["Authorization"] = s.c.authorization
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := Request{
		method:  "PATCH",
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
