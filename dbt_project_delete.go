package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

// ProjectDeleteService implements the Project Management, Delete a project API.
// Ref. https://fivetran.com/docs/rest-api/projects#deleteaproject
type DbtProjectDeleteService struct {
	c         *Client
	projectID *string
}

type DbtProjectDeleteResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) NewDbtProjectDelete() *ProjectDeleteService {
	return &ProjectDeleteService{c: c}
}

func (s *ProjectDeleteService) ProjectID(value string) *ProjectDeleteService {
	s.projectID = &value
	return s
}

func (s *ProjectDeleteService) Do(ctx context.Context) (ProjectDeleteResponse, error) {
	var response ProjectDeleteResponse

	if s.projectID == nil {
		return response, fmt.Errorf("missing required projectID")
	}

	url := fmt.Sprintf("%v/projects/%v", s.c.baseURL, *s.projectID)
	expectedStatus := 200

	headers := s.c.commonHeaders()

	r := request{
		method:           "DELETE",
		url:              url,
		body:             nil,
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
