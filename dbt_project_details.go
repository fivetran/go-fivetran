package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

// DbtProjectDetailsService implements the dbt project management, retrive dbt project details api
// Ref. GET https://api.fivetran.com/v1/dbt/projects/{id}
type DbtProjectDetailsService struct {
	c            *Client
	dbtProjectID *string
}

type DbtProjectDetailsResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		DbtProjectResponseBase
	} `json:"data"`
}

func (c *Client) NewDbtProjectDetailsService() *DbtProjectDetailsService {
	return &DbtProjectDetailsService{c: c}
}

func (s *DbtProjectDetailsService) DbtProjectID(value string) *DbtProjectDetailsService {
	s.dbtProjectID = &value
	return s
}

func (s *DbtProjectDetailsService) Do(ctx context.Context) (DbtProjectDetailsResponse, error) {
	var response DbtProjectDetailsResponse
	if s.dbtProjectID == nil {
		return response, fmt.Errorf("missing required DbtProjectID")
	}

	url := fmt.Sprintf("%v/dbt/projects/%v", s.c.baseURL, *s.dbtProjectID)
	expectedStatus := 200

	headers := s.c.commonHeaders()
	headers["Accept"] = restAPIv2

	r := request{
		method:  "GET",
		url:     url,
		body:    nil,
		queries: nil,
		headers: headers,
		client:  s.c.httpClient,
	}

	respBody, respStatus, err := r.httpRequest(ctx)

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
