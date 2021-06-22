package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type groupListConnectorsService struct {
	c       *Client
	groupID *string
	limit   *int
	cursor  *string
	schema  *string
}

type GroupListConnectorsResponse struct {
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

func (c *Client) NewGroupListConnectors() *groupListConnectorsService {
	return &groupListConnectorsService{c: c}
}

func (s *groupListConnectorsService) GroupID(value string) *groupListConnectorsService {
	s.groupID = &value
	return s
}

func (s *groupListConnectorsService) Limit(value int) *groupListConnectorsService {
	s.limit = &value
	return s
}

func (s *groupListConnectorsService) Cursor(value string) *groupListConnectorsService {
	s.cursor = &value
	return s
}

func (s *groupListConnectorsService) Schema(value string) *groupListConnectorsService {
	s.schema = &value
	return s
}

func (s *groupListConnectorsService) Do(ctx context.Context) (GroupListConnectorsResponse, error) {
	var response GroupListConnectorsResponse

	if s.groupID == nil {
		return response, fmt.Errorf("missing required GroupID")
	}

	url := fmt.Sprintf("%v/groups/%v/connectors", s.c.baseURL, *s.groupID)
	expectedStatus := 200

	headers := make(map[string]string)
	headers["Authorization"] = s.c.authorization

	queries := make(map[string]string)
	if s.cursor != nil {
		queries["cursor"] = *s.cursor
	}
	if s.limit != nil {
		queries["limit"] = fmt.Sprint(*s.limit)
	}
	if s.schema != nil {
		queries["schema"] = *s.schema
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
		return response, err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return response, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return response, err
	}

	return response, nil
}
