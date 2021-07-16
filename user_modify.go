package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// UserModifyService implements the User Management, Modify a User API.
// Ref. https://fivetran.com/docs/rest-api/users#modifyauser
type UserModifyService struct {
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
		Verified   *bool     `json:"verified"`
		Invited    *bool     `json:"invited"`
		Picture    string    `json:"picture"`
		Phone      string    `json:"phone"`
		LoggedInAt time.Time `json:"logged_in_at"`
		CreatedAt  time.Time `json:"created_at"`
	} `json:"data"`
}

func (c *Client) NewUserModify() *UserModifyService {
	return &UserModifyService{c: c}
}

func (s *UserModifyService) request() *userModifyRequest {
	return &userModifyRequest{
		GivenName:  s.givenName,
		FamilyName: s.familyName,
		Phone:      s.phone,
		Picture:    s.picture,
		Role:       s.role,
	}
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

func (s *UserModifyService) Picture(value string) *UserModifyService {
	s.picture = &value
	return s
}

func (s *UserModifyService) Role(value string) *UserModifyService {
	s.role = &value
	return s
}

func (s *UserModifyService) Do(ctx context.Context) (UserModifyResponse, error) {
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

	r := request{
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
		err := fmt.Errorf("status code: %v; expected: %v", respStatus, expectedStatus)
		return response, err
	}

	return response, nil
}
