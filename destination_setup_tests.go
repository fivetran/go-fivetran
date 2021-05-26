package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

// F stands for Field
// needs to be exported because of json.Marshal()
type DestinationSetupTestsService struct {
	c                  *Client
	destinationID      string
	FtrustCertificates *bool `json:"trust_certificates,omitempty"`
	FtrustFingerprints *bool `json:"trust_fingerprints,omitempty"`
}

type DestinationSetupTests struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID             string `json:"id"`
		GroupID        string `json:"group_id"`
		Service        string `json:"service"`
		Region         string `json:"region"`
		TimeZoneOffset string `json:"time_zone_offset"`
		SetupStatus    string `json:"setup_status"`
		SetupTests     []struct {
			Title   string `json:"title"`
			Status  string `json:"status"`
			Message string `json:"message"`
		} `json:"setup_tests"`
		Config DestinationConfig `json:"config"`
	} `json:"data"`
}

func (c *Client) NewDestinationSetupTests() *DestinationSetupTestsService {
	return &DestinationSetupTestsService{c: c}
}

func (s *DestinationSetupTestsService) DestinationID(destinationID string) *DestinationSetupTestsService {
	s.destinationID = destinationID
	return s
}

func (s *DestinationSetupTestsService) TrustCertificates(trustCertificates bool) *DestinationSetupTestsService {
	s.FtrustCertificates = &trustCertificates
	return s
}

func (s *DestinationSetupTestsService) TrustFingerprints(trustFingerprints bool) *DestinationSetupTestsService {
	s.FtrustFingerprints = &trustFingerprints
	return s
}

func (s *DestinationSetupTestsService) Do(ctx context.Context) (DestinationSetupTests, error) {
	if s.destinationID == "" { // we don't validate business rules (unless it is strictly necessary)
		err := fmt.Errorf("missing required DestinationID")
		return DestinationSetupTests{}, err
	}

	url := fmt.Sprintf("%v/destinations/%v/test", s.c.baseURL, s.destinationID)
	expectedStatus := 200
	headers := make(map[string]string)

	headers["Authorization"] = s.c.authorization
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s)
	if err != nil {
		return DestinationSetupTests{}, err
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
		return DestinationSetupTests{}, err
	}

	var destinationSetupTests DestinationSetupTests
	if err := json.Unmarshal(respBody, &destinationSetupTests); err != nil {
		return DestinationSetupTests{}, err
	}

	// converts destinationCreate.Data.Config.Fport to int. Should be removed
	// when https://fivetran.height.app/T-97508 fixed.
	switch destinationSetupTests.Data.Config.Fport.(type) {
	case string:
		destinationSetupTests.Data.Config.Fport, err = strconv.Atoi(destinationSetupTests.Data.Config.Fport.(string))
		if err != nil {
			return DestinationSetupTests{}, err
		}

	default:
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return destinationSetupTests, err
	}

	return destinationSetupTests, nil
}
