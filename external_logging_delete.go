package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ExternalLoggingDeleteService implements the Log Management, Delete a Log Service API.
// Ref. https://fivetran.com/docs/rest-api/log-service-management#deletealogservice
type ExternalLoggingDeleteService struct {
	c                 *Client
	externalLoggingId *string
}

func (c *Client) NewExternalLoggingDelete() *ExternalLoggingDeleteService {
	return &ExternalLoggingDeleteService{c: c}
}

func (s *ExternalLoggingDeleteService) ExternalLoggingId(value string) *ExternalLoggingDeleteService {
	s.externalLoggingId = &value
	return s
}

func (s *ExternalLoggingDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse

	if s.externalLoggingId == nil {
		return response, fmt.Errorf("missing required ExternalLoggingId")
	}

	url := fmt.Sprintf("%v/external-logging/%v", s.c.baseURL, *s.externalLoggingId)
	expectedStatus := 200

	headers := s.c.commonHeaders()

	r := httputils.Request{
		Method:           "DELETE",
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
