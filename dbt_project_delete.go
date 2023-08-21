package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

// ProjectDeleteService implements the Project Management, Delete a project API.
// Ref. https://fivetran.com/docs/rest-api/projects#deleteaproject
type ProjectDeleteService struct {
	c         *Client
	projectID *string
}

type projectDeleteResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) NewprojectDelete() *ProjectDeleteService {
	return &ProjectDeleteService{c: c}
}

func (s *ProjectDeleteService) ProjectID(value string) *ProjectDeleteService {
	s.projectID = &value
	return s
}

func (s *ProjectDeleteService) Do(ctx context.Context) (projectDeleteResponse, error) {
	var response projectDeleteResponse

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
