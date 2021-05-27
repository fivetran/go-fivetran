package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type ConnectorDetailsService struct {
	c           *Client
	connectorID *string
}

type ConnectorDetails struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID             string    `json:"id"`
		GroupID        string    `json:"group_id"`
		Service        string    `json:"service"`
		ServiceVersion int       `json:"service_version"`
		Schema         string    `json:"schema"`
		ConnectedBy    string    `json:"connected_by"`
		CreatedAt      time.Time `json:"created_at"`
		SucceededAt    time.Time `json:"succeeded_at"`
		FailedAt       time.Time `json:"failed_at"`
		SyncFrequency  int       `json:"sync_frequency"`
		ScheduleType   string    `json:"schedule_type"`
		Status         struct {
			SetupState       string `json:"setup_state"`
			SyncState        string `json:"sync_state"`
			UpdateState      string `json:"update_state"`
			IsHistoricalSync bool   `json:"is_historical_sync"`
			Tasks            []struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			} `json:"tasks"`
			Warnings []struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			} `json:"warnings"`
		} `json:"status"`
		Config ConnectorConfig `json:"config"`
	} `json:"data"`
}

func (c *Client) NewConnectorDetailsService() *ConnectorDetailsService {
	return &ConnectorDetailsService{c: c}
}

func (s *ConnectorDetailsService) ConnectorID(connectorID string) *ConnectorDetailsService {
	s.connectorID = &connectorID
	return s
}

func (s *ConnectorDetailsService) Do(ctx context.Context) (ConnectorDetails, error) {
	if s.connectorID == nil { // we don't validate business rules (unless it is strictly necessary)
		err := fmt.Errorf("missing required ConnectorID")
		return ConnectorDetails{}, err
	}

	url := fmt.Sprintf("%v/connectors/%v", s.c.baseURL, *s.connectorID)
	expectedStatus := 200
	headers := make(map[string]string)

	headers["Authorization"] = s.c.authorization

	r := Request{
		method:  "GET",
		url:     url,
		body:    nil,
		queries: nil,
		headers: headers,
	}

	respBody, respStatus, err := httpRequest(r, ctx)
	if err != nil {
		return ConnectorDetails{}, err
	}

	var connectorDetails ConnectorDetails
	if err := json.Unmarshal(respBody, &connectorDetails); err != nil {
		return ConnectorDetails{}, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return connectorDetails, err
	}

	return connectorDetails, nil
}
