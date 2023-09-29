package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ExternalLoggingSetupTestsService implements the Log Management, Run Log service setup tests API.
// Ref. https://fivetran.com/docs/rest-api/log-service-management#runlogservicesetuptests
type ExternalLoggingSetupTestsService struct {
	c                 *Client
	externalLoggingId *string
}

type ExternalLoggingSetupTestsResponse struct {
	common.CommonResponse
	Data struct {
		SetupTests []common.SetupTestResponse `json:"setup_tests"`
	} `json:"data"`
}

func (c *Client) NewExternalLoggingSetupTests() *ExternalLoggingSetupTestsService {
	return &ExternalLoggingSetupTestsService{c: c}
}

func (s *ExternalLoggingSetupTestsService) ExternalLoggingId(value string) *ExternalLoggingSetupTestsService {
	s.externalLoggingId = &value
	return s
}

func (s *ExternalLoggingSetupTestsService) Do(ctx context.Context) (ExternalLoggingSetupTestsResponse, error) {
	var response ExternalLoggingSetupTestsResponse

	if s.externalLoggingId == nil {
		return response, fmt.Errorf("missing required externalLoggingId")
	}

	url := fmt.Sprintf("%v/external-logging/%v/test", s.c.baseURL, *s.externalLoggingId)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"
	headers["Accept"] = restAPIv2

	r := httputils.Request{
		Method:           "POST",
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
