package fivetran

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

// F stands for Field
// needs to be exported because of json.Marshal()
type DestinationModifyService struct {
	c                  *Client
	destinationID      string
	Fregion            string             `json:"region,omitempty"`
	FtimeZoneOffset    string             `json:"time_zone_offset,omitempty"`
	Fconfig            *DestinationConfig `json:"config,omitempty"`
	FtrustCertificates *bool              `json:"trust_certificates,omitempty"`
	FtrustFingerprints *bool              `json:"trust_fingerprints,omitempty"`
	FrunSetupTests     *bool              `json:"run_setup_tests,omitempty"`
}

type DestinationModify struct {
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

func (c *Client) NewDestinationModifyService() *DestinationModifyService {
	return &DestinationModifyService{c: c}
}

func (s *DestinationModifyService) DestinationID(destinationID string) *DestinationModifyService {
	s.destinationID = destinationID
	return s
}

func (s *DestinationModifyService) Region(region string) *DestinationModifyService {
	s.Fregion = region
	return s
}

func (s *DestinationModifyService) TimeZoneOffset(timeZoneOffset string) *DestinationModifyService {
	s.FtimeZoneOffset = timeZoneOffset
	return s
}

func (s *DestinationModifyService) Config(config *DestinationConfig) *DestinationModifyService {
	s.Fconfig = config
	return s
}

func (s *DestinationModifyService) TrustCertificates(trustCertificates bool) *DestinationModifyService {
	s.FtrustCertificates = &trustCertificates
	return s
}

func (s *DestinationModifyService) TrustFingerprints(trustFingerprints bool) *DestinationModifyService {
	s.FtrustFingerprints = &trustFingerprints
	return s
}

func (s *DestinationModifyService) RunSetupTests(runSetupTests bool) *DestinationModifyService {
	s.FrunSetupTests = &runSetupTests
	return s
}

func (s *DestinationModifyService) Do(ctx context.Context) (DestinationModify, error) {
	if s.destinationID == "" { // we don't validate business rules (unless it is strictly necessary)
		err := fmt.Errorf("missing required DestinationID")
		return DestinationModify{}, err
	}

	url := fmt.Sprintf("%v/destinations/%v", s.c.baseURL, s.destinationID)
	expectedStatus := 200
	headers := make(map[string]string)

	headers["Authorization"] = s.c.authorization
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s)
	if err != nil {
		return DestinationModify{}, err
	}

	r := Request{
		method:  "PATCH",
		url:     url,
		body:    bytes.NewReader(reqBody),
		queries: nil,
		headers: headers,
	}

	respBody, respStatus, err := httpRequest(r, ctx)
	if err != nil {
		return DestinationModify{}, err
	}

	var destinationModify DestinationModify
	if err := json.Unmarshal(respBody, &destinationModify); err != nil {
		return DestinationModify{}, err
	}

	// converts destinationModify.Data.Config.Fport to int. Should be removed
	// when https://fivetran.height.app/T-97508 fixed.
	switch destinationModify.Data.Config.Fport.(type) {
	case string:
		destinationModify.Data.Config.Fport, err = strconv.Atoi(destinationModify.Data.Config.Fport.(string))
		if err != nil {
			return DestinationModify{}, err
		}

	default:
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return destinationModify, err
	}

	return destinationModify, nil
}
