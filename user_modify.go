package fivetran

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/fivetran/go-fivetran/users"
)

// UserModifyService implements the User Management, Modify a User API.
// Ref. https://fivetran.com/docs/rest-api/users#modifyauser
type UserModifyService struct {
	c            *Client
	userID       *string
	givenName    *string
	familyName   *string
	phone        *string
	picture      *string
	role         *string
	clearPicture bool
	clearPhone   bool
}

type userModifyRequest struct {
	GivenName  *string         `json:"given_name,omitempty"`
	FamilyName *string         `json:"family_name,omitempty"`
	Phone      *nullableString `json:"phone,omitempty"`
	Picture    *nullableString `json:"picture,omitempty"`
	Role       *string         `json:"role,omitempty"`
}

func (c *Client) NewUserModify() *UserModifyService {
	return &UserModifyService{c: c}
}

func (s *UserModifyService) request() *userModifyRequest {
	return &userModifyRequest{
		GivenName:  s.givenName,
		FamilyName: s.familyName,
		Phone:      newNullableString(s.phone, s.clearPhone),
		Picture:    newNullableString(s.picture, s.clearPicture),
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

func (s *UserModifyService) Do(ctx context.Context) (users.UserDetailsResponse, error) {
	var response users.UserDetailsResponse

	if s.userID == nil {
		return response, fmt.Errorf("missing required UserID")
	}

	if s.clearPhone && s.phone != nil {
		return response, errors.New("can't 'set phone' and 'clear phone' in one request")
	}

	if s.clearPicture && s.picture != nil {
		return response, errors.New("can't 'set picture' and 'clear picture' in one request")
	}

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	url := fmt.Sprintf("%v/users/%v", s.c.baseURL, *s.userID)
	expectedStatus := 200
	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"

	r := request{
		method:           "PATCH",
		url:              url,
		body:             reqBody,
		queries:          nil,
		headers:          headers,
		client:           s.c.httpClient,
		handleRateLimits: s.c.handleRateLimits,
		maxRetryAttempts: s.c.maxRetryAttempts,
	}

	respBody, respStatus, err := r.httpRequest(ctx)
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
