package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/dbt"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type DbtProjectDetailsService struct {
	c            *Client
	dbtProjectID *string
}

func (c *Client) NewDbtProjectDetails() *DbtProjectDetailsService {
	return &DbtProjectDetailsService{c: c}
}

func (s *DbtProjectDetailsService) DbtProjectID(value string) *DbtProjectDetailsService {
	s.dbtProjectID = &value
	return s
}

func (s *DbtProjectDetailsService) Do(ctx context.Context) (dbt.DbtProjectDetailsResponse, error) {
	var response dbt.DbtProjectDetailsResponse

	if s.dbtProjectID == nil {
		return response, fmt.Errorf("missing required DbtProjectId")
	}

	url := fmt.Sprintf("%v/dbt/projects/%v", s.c.baseURL, *s.dbtProjectID)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Accept"] = restAPIv2

	r := httputils.Request{
		Method:           "GET",
		Url:              url,
		Body:             nil,
		Queries:          nil,
		Headers:          headers,
		Client:           s.c.httpClient,
		HandleRateLimits: s.c.handleRateLimits,
		MaxRetryAttempts: s.c.maxRetryAttempts,
	}

	respBody, respStatus, err := r.Do(ctx)

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
