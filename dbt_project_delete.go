package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
)

// ProjectDeleteService implements the Project Management, Delete a project API.
// Ref. https://fivetran.com/docs/rest-api/projects#deleteaproject
type DbtProjectDeleteService struct {
	c            *Client
	dbtProjectID *string
}

func (c *Client) NewDbtProjectDelete() *DbtProjectDeleteService {
	return &DbtProjectDeleteService{c: c}
}

func (s *DbtProjectDeleteService) DbtProjectID(value string) *DbtProjectDeleteService {
	s.dbtProjectID = &value
	return s
}

func (s *DbtProjectDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse

	if s.dbtProjectID == nil {
		return response, fmt.Errorf("missing required dbtProjectID")
	}

	url := fmt.Sprintf("%v/dbt/projects/%v", s.c.baseURL, *s.dbtProjectID)
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
