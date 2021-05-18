package fivetran

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

// F stands for Field
// needs to be exported because of json.Marshal()
type DestinationCreateService struct {
	c                  *Client
	FgroupID           string             `json:"group_id,omitempty"`
	Fservice           string             `json:"service,omitempty"`
	Fregion            string             `json:"region,omitempty"`
	FtimeZoneOffset    string             `json:"time_zone_offset,omitempty"`
	Fconfig            *DestinationConfig `json:"config,omitempty"`
	FtrustCertificates *bool              `json:"trust_certificates,omitempty"`
	FtrustFingerprints *bool              `json:"trust_fingerprints,omitempty"`
	FrunSetupTests     *bool              `json:"run_setup_tests,omitempty"`
}

type DestinationCreate struct {
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
		Config DestinationConfigTemp `json:"config"` // When https://fivetran.height.app/T-97508 is fixed, type should change to DestinationConfig
	} `json:"data"`
}

func (c *Client) NewDestinationCreateService() *DestinationCreateService {
	return &DestinationCreateService{c: c}
}

func (s *DestinationCreateService) GroupID(groupID string) *DestinationCreateService {
	s.FgroupID = groupID
	return s
}

func (s *DestinationCreateService) Service(service string) *DestinationCreateService {
	s.Fservice = service
	return s
}

func (s *DestinationCreateService) Region(region string) *DestinationCreateService {
	s.Fregion = region
	return s
}

func (s *DestinationCreateService) TimeZoneOffset(timeZoneOffset string) *DestinationCreateService {
	s.FtimeZoneOffset = timeZoneOffset
	return s
}

func (s *DestinationCreateService) Config(config *DestinationConfig) *DestinationCreateService {
	s.Fconfig = config
	return s
}

func (s *DestinationCreateService) TrustCertificates(trustCertificates bool) *DestinationCreateService {
	s.FtrustCertificates = &trustCertificates
	return s
}

func (s *DestinationCreateService) TrustFingerprints(trustFingerprints bool) *DestinationCreateService {
	s.FtrustFingerprints = &trustFingerprints
	return s
}

func (s *DestinationCreateService) RunSetupTests(runSetupTests bool) *DestinationCreateService {
	s.FrunSetupTests = &runSetupTests
	return s
}

func (s *DestinationCreateService) Do(ctx context.Context) (DestinationCreate, error) {
	url := fmt.Sprintf("%v/destinations", s.c.baseURL)
	expectedStatus := 201
	headers := make(map[string]string)

	headers["Authorization"] = s.c.authorization
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s)
	if err != nil {
		return DestinationCreate{}, err
	}

	r := Request{
		method:  "POST",
		url:     url,
		body:    bytes.NewReader(reqBody),
		queries: nil,
		headers: headers,
	}

	respBody, respStatus, err := httpRequest(r, ctx)
	if err != nil {
		return DestinationCreate{}, err
	}

	var destinationCreate DestinationCreate
	if err := json.Unmarshal(respBody, &destinationCreate); err != nil {
		return DestinationCreate{}, err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected %v", respStatus, expectedStatus)
		return destinationCreate, err
	}

	return destinationCreate, nil
}
