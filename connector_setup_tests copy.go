package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// F stands for Field
// needs to be exported because of json.Marshal()
type ConnectorSetupTestsService struct {
	c                  *Client
	connectorID        *string
	FtrustCertificates *bool `json:"trust_certificates,omitempty"`
	FtrustFingerprints *bool `json:"trust_fingerprints,omitempty"`
}

type ConnectorSetupTests struct {
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
		Config ConnectorConfig `json:"config"`
	} `json:"data"`
}

func (c *Client) NewConnectorSetupTests() *ConnectorSetupTestsService {
	return &ConnectorSetupTestsService{c: c}
}

func (s *ConnectorSetupTestsService) ConnectorID(connectorID string) *ConnectorSetupTestsService {
	s.connectorID = &connectorID
	return s
}

func (s *ConnectorSetupTestsService) TrustCertificates(trustCertificates bool) *ConnectorSetupTestsService {
	s.FtrustCertificates = &trustCertificates
	return s
}

func (s *ConnectorSetupTestsService) TrustFingerprints(trustFingerprints bool) *ConnectorSetupTestsService {
	s.FtrustFingerprints = &trustFingerprints
	return s
}

func (s *ConnectorSetupTestsService) Do(ctx context.Context) (ConnectorSetupTests, error) {
	if s.connectorID == nil { // we don't validate business rules (unless it is strictly necessary)
		err := fmt.Errorf("missing required ConnectorID")
		return ConnectorSetupTests{}, err
	}

	url := fmt.Sprintf("%v/connectors/%v/test", s.c.baseURL, *s.connectorID)
	expectedStatus := 200
	headers := make(map[string]string)

	headers["Authorization"] = s.c.authorization
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s)
	if err != nil {
		return ConnectorSetupTests{}, err
	}

	r := Request{
		method:  "POST",
		url:     url,
		body:    reqBody,
		queries: nil,
		headers: headers,
	}

	respBody, respStatus, err := httpRequest(r, ctx)
	if err != nil {
		return ConnectorSetupTests{}, err
	}

	var connectorSetupTests ConnectorSetupTests
	if err := json.Unmarshal(respBody, &connectorSetupTests); err != nil {
		return ConnectorSetupTests{}, err
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
		return connectorSetupTests, err
	}

	return connectorSetupTests, nil
}
