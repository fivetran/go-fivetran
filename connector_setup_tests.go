package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type ConnectorSetupTestsService struct {
	c                 *Client
	connectorID       *string
	trustCertificates *bool
	trustFingerprints *bool
}

type connectorSetupTestsRequest struct {
	TrustCertificates *bool `json:"trust_certificates,omitempty"`
	TrustFingerprints *bool `json:"trust_fingerprints,omitempty"`
}

type ConnectorSetupTestsResponse struct {
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
		SetupTests []struct {
			Title   string `json:"title"`
			Status  string `json:"status"`
			Message string `json:"message"`
		} `json:"setup_tests"`
		Config ConnectorConfigResponse `json:"config"`
	} `json:"data"`
}

func (c *Client) NewConnectorSetupTests() *ConnectorSetupTestsService {
	return &ConnectorSetupTestsService{c: c}
}

func (s *ConnectorSetupTestsService) request() *connectorSetupTestsRequest {
	return &connectorSetupTestsRequest{
		TrustCertificates: s.trustCertificates,
		TrustFingerprints: s.trustFingerprints,
	}
}

func (s *ConnectorSetupTestsService) ConnectorID(value string) *ConnectorSetupTestsService {
	s.connectorID = &value
	return s
}

func (s *ConnectorSetupTestsService) TrustCertificates(value bool) *ConnectorSetupTestsService {
	s.trustCertificates = &value
	return s
}

func (s *ConnectorSetupTestsService) TrustFingerprints(value bool) *ConnectorSetupTestsService {
	s.trustFingerprints = &value
	return s
}

func (s *ConnectorSetupTestsService) Do(ctx context.Context) (ConnectorSetupTestsResponse, error) {
	var response ConnectorSetupTestsResponse

	if s.connectorID == nil {
		return response, fmt.Errorf("missing required ConnectorID")
	}

	url := fmt.Sprintf("%v/connectors/%v/test", s.c.baseURL, *s.connectorID)
	expectedStatus := 200

	headers := make(map[string]string)
	headers["Authorization"] = s.c.authorization
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := request{
		method:  "POST",
		url:     url,
		body:    reqBody,
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

	// // converts destinationCreate.Data.Config.Fport to int. Should be removed
	// // when https://fivetran.height.app/T-97508 fixed.
	// switch destinationSetupTests.Data.Config.Fport.(type) {
	// case string:
	// 	destinationSetupTests.Data.Config.Fport, err = strconv.Atoi(destinationSetupTests.Data.Config.Fport.(string))
	// 	if err != nil {
	// 		return DestinationSetupTests{}, err
	// 	}

	// default:
	// }

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return response, err
	}

	return response, nil
}
