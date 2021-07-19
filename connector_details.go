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

type ConnectorDetailsResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
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
		Config ConnectorConfigResponse `json:"config"`
	} `json:"data"`
	// SourceSyncDetails not implemented yet. T-114130
}

func (c *Client) NewConnectorDetails() *ConnectorDetailsService {
	return &ConnectorDetailsService{c: c}
}

func (s *ConnectorDetailsService) ConnectorID(value string) *ConnectorDetailsService {
	s.connectorID = &value
	return s
}

func (s *ConnectorDetailsService) Do(ctx context.Context) (ConnectorDetailsResponse, error) {
	var response ConnectorDetailsResponse

	if s.connectorID == nil {
		return response, fmt.Errorf("missing required ConnectorID")
	}

	url := fmt.Sprintf("%v/connectors/%v", s.c.baseURL, *s.connectorID)
	expectedStatus := 200

	headers := make(map[string]string)
	headers["Authorization"] = s.c.authorization
	headers["Accept"] = restAPIv2

	r := request{
		method:  "GET",
		url:     url,
		body:    nil,
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
