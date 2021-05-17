package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type GroupListConnectorsService struct {
	c      *Client
	id     string
	limit  int
	cursor string
	schema string
}

type GroupListConnectors struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Items []struct {
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
			DailySyncTime string `json:"daily_sync_time"`
		} `json:"items"`
		NextCursor string `json:"next_cursor"`
	} `json:"data"`
}

func (c *Client) NewGroupListConnectorsService() *GroupListConnectorsService {
	return &GroupListConnectorsService{c: c}
}

func (s *GroupListConnectorsService) ID(id string) *GroupListConnectorsService {
	s.id = id
	return s
}

func (s *GroupListConnectorsService) Limit(limit int) *GroupListConnectorsService {
	s.limit = limit
	return s
}

func (s *GroupListConnectorsService) Cursor(cursor string) *GroupListConnectorsService {
	s.cursor = cursor
	return s
}

func (s *GroupListConnectorsService) Schema(schema string) *GroupListConnectorsService {
	s.schema = schema
	return s
}

func (s *GroupListConnectorsService) Do(ctx context.Context) (GroupListConnectors, error) {
	if s.id == "" {
		err := fmt.Errorf("missing required ID")
		return GroupListConnectors{}, err
	}

	url := fmt.Sprintf("%v/groups/%v/connectors", s.c.baseURL, s.id)
	expectedStatus := 200
	headers := make(map[string]string)
	queries := make(map[string]string)

	headers["Authorization"] = s.c.authorization

	if s.cursor != "" {
		queries["cursor"] = s.cursor
	}

	if s.limit != 0 {
		queries["limit"] = fmt.Sprint(s.limit)
	}

	if s.schema != "" {
		queries["schema"] = s.schema
	}

	r := Request{
		method:  "GET",
		url:     url,
		body:    nil,
		queries: queries,
		headers: headers,
	}

	respBody, respStatus, err := httpRequest(r, ctx)
	if err != nil {
		return GroupListConnectors{}, err
	}

	var groupListConnectors GroupListConnectors
	if err := json.Unmarshal(respBody, &groupListConnectors); err != nil {
		return GroupListConnectors{}, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return groupListConnectors, err
	}

	return groupListConnectors, nil
}
