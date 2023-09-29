package fivetran

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	"github.com/fivetran/go-fivetran/connectors"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// GroupListConnectorsService implements the Group Management, List All Connectors within a Group API.
// Ref. https://fivetran.com/docs/rest-api/groups#listallconnectorswithinagroup
type GroupListConnectorsService struct {
	c       *Client
	groupID *string
	limit   *int
	cursor  *string
	schema  *string
}

type ConnectorsStatus struct {
	SetupState       string                  `json:"setup_state"`
	SyncState        string                  `json:"sync_state"`
	UpdateState      string                  `json:"update_state"`
	IsHistoricalSync *bool                   `json:"is_historical_sync"`
	Tasks            []common.CommonResponse `json:"tasks"`
	Warnings         []common.CommonResponse `json:"warnings"`
}

type GroupListConnectorsResponse struct {
	common.CommonResponse
	Data struct {
		Items      []connectors.DetailsResponseDataCommon `json:"items"`
		NextCursor string                                 `json:"next_cursor"`
	} `json:"data"`
}

func (c *Client) NewGroupListConnectors() *GroupListConnectorsService {
	return &GroupListConnectorsService{c: c}
}

func (s *GroupListConnectorsService) GroupID(value string) *GroupListConnectorsService {
	s.groupID = &value
	return s
}

func (s *GroupListConnectorsService) Limit(value int) *GroupListConnectorsService {
	s.limit = &value
	return s
}

func (s *GroupListConnectorsService) Cursor(value string) *GroupListConnectorsService {
	s.cursor = &value
	return s
}

func (s *GroupListConnectorsService) Schema(value string) *GroupListConnectorsService {
	s.schema = &value
	return s
}

func (s *GroupListConnectorsService) Do(ctx context.Context) (GroupListConnectorsResponse, error) {
	var response GroupListConnectorsResponse

	if s.groupID == nil {
		return response, fmt.Errorf("missing required GroupID")
	}

	url := fmt.Sprintf("%v/groups/%v/connectors", s.c.baseURL, *s.groupID)
	expectedStatus := 200

	headers := s.c.commonHeaders()

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

	r := httputils.Request{
		Method:           "GET",
		Url:              url,
		Body:             nil,
		Queries:          queries,
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
