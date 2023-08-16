package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// ConnectorDetailsService implements the Connector Management, Retrieve Connector Details API.
// Ref. https://fivetran.com/docs/rest-api/connectors#retrieveconnectordetails
type ConnectorDetailsService struct {
	c           *Client
	connectorID *string
}

type ConnectorDetailsdataBase struct {
	ID              string    `json:"id"`
	GroupID         string    `json:"group_id"`
	Service         string    `json:"service"`
	ServiceVersion  *int      `json:"service_version"`
	Schema          string    `json:"schema"`
	ConnectedBy     string    `json:"connected_by"`
	CreatedAt       time.Time `json:"created_at"`
	SucceededAt     time.Time `json:"succeeded_at"`
	FailedAt        time.Time `json:"failed_at"`
	Paused          *bool     `json:"paused"`
	PauseAfterTrial *bool     `json:"pause_after_trial"`
	DailySyncTime   string    `json:"daily_sync_time"`
	SyncFrequency   *int      `json:"sync_frequency"`
	ScheduleType    string    `json:"schedule_type"`
	Status          struct {
		SetupState       string `json:"setup_state"`
		SyncState        string `json:"sync_state"`
		UpdateState      string `json:"update_state"`
		IsHistoricalSync *bool  `json:"is_historical_sync"`
		Tasks            []struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"tasks"`
		Warnings []struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"warnings"`
	} `json:"status"`
}

type ConnectorDetailsResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ConnectorDetailsdataBase
		Config ConnectorConfigResponse `json:"config"`
	} `json:"data"`
}

type ConnectorCustomDetailsResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ConnectorDetailsdataBase
		Config map[string]interface{} `json:"config"`
	} `json:"data"`
}

type ConnectorCustomMergedDetailsResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ConnectorDetailsdataBase
		CustomConfig map[string]interface{}  `json:"config"`
		Config       ConnectorConfigResponse // no mapping here
	} `json:"data"`
}

func (c *Client) NewConnectorDetails() *ConnectorDetailsService {
	return &ConnectorDetailsService{c: c}
}

func (s *ConnectorDetailsService) ConnectorID(value string) *ConnectorDetailsService {
	s.connectorID = &value
	return s
}

func (s *ConnectorDetailsService) do(ctx context.Context, response any) error {
	if s.connectorID == nil {
		return fmt.Errorf("missing required ConnectorID")
	}

	url := fmt.Sprintf("%v/connectors/%v", s.c.baseURL, *s.connectorID)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Accept"] = restAPIv2

	r := request{
		method:           "GET",
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

func (s *ConnectorDetailsService) Do(ctx context.Context) (ConnectorDetailsResponse, error) {
	var response ConnectorDetailsResponse

	err := s.do(ctx, &response)

	return response, err
}

func (s *ConnectorDetailsService) DoCustom(ctx context.Context) (ConnectorCustomDetailsResponse, error) {
	var response ConnectorCustomDetailsResponse

	err := s.do(ctx, &response)

	return response, err
}

func (s *ConnectorDetailsService) DoCustomMerged(ctx context.Context) (ConnectorCustomMergedDetailsResponse, error) {
	var response ConnectorCustomMergedDetailsResponse

	err := s.do(ctx, &response)

	if err == nil {
		err = FetchFromMap(&response.Data.CustomConfig, &response.Data.Config)
	}

	return response, err
}
