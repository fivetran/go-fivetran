package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	externallogging "github.com/fivetran/go-fivetran/external_logging"
	httputils "github.com/fivetran/go-fivetran/http_utils"
	"github.com/fivetran/go-fivetran/utils"
)

// ExternalLoggingDetailsService implements the Log Management, Retrieve Log Service details API.
// Ref. https://fivetran.com/docs/rest-api/log-service-management#retrievelogservicedetails
type ExternalLoggingDetailsService struct {
	c                 *Client
	externalLoggingId *string
}

func (c *Client) NewExternalLoggingDetails() *ExternalLoggingDetailsService {
	return &ExternalLoggingDetailsService{c: c}
}

func (s *ExternalLoggingDetailsService) ExternalLoggingId(value string) *ExternalLoggingDetailsService {
	s.externalLoggingId = &value
	return s
}

func (s *ExternalLoggingDetailsService) Do(ctx context.Context) (externallogging.ExternalLoggingResponse, error) {
	var response externallogging.ExternalLoggingResponse

	err := s.do(ctx, &response)

	return response, err
}

func (s *ExternalLoggingDetailsService) DoCustom(ctx context.Context) (externallogging.ExternalLoggingCustomResponse, error) {
	var response externallogging.ExternalLoggingCustomResponse

	err := s.do(ctx, &response)

	return response, err
}

func (s *ExternalLoggingDetailsService) DoCustomMerged(ctx context.Context) (externallogging.ExternalLoggingCustomMergedResponse, error) {
	var response externallogging.ExternalLoggingCustomMergedResponse

	err := s.do(ctx, &response)

	if err == nil {
		err = utils.FetchFromMap(&response.Data.CustomConfig, &response.Data.Config)
	}

	return response, err
}

func (s *ExternalLoggingDetailsService) do(ctx context.Context, response any) error {

	if s.externalLoggingId == nil {
		return fmt.Errorf("missing required ExternalLoggingId")
	}

	url := fmt.Sprintf("%v/external-logging/%v", s.c.baseURL, *s.externalLoggingId)
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
		return err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected: %v", respStatus, expectedStatus)
		return err
	}

	return nil
}
