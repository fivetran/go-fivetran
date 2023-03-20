package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

// DbtTransformationService implements the dbt transformation management,
// retrive dbt transformation details api
// Ref. GET https://api.fivetran.com/v1/dbt/transformations/{transformation_id}

type DbtTransformationDetailsService struct {
	c                   *Client
	dbtTransformationID *string
}

type DbtTransformationDetailsResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		DbtTransformationResponseBase
		Schedule DbtTransformationScheduleResponse `json:"schedule"`
	} `json:"data"`
}

func (c *Client) NewDbtTransformationDetailsService() *DbtTransformationDetailsService {
	return &DbtTransformationDetailsService{c: c}
}

func (s *DbtTransformationDetailsService) DbtTransformationID(value string) *DbtTransformationDetailsService {
	s.dbtTransformationID = &value
	return s
}

func (s *DbtTransformationDetailsService) do(ctx context.Context, response any) error {
	if s.dbtTransformationID == nil {
		return fmt.Errorf("missing required DbtTransformationID")
	}

	url := fmt.Sprintf("%v/dbt/transformations/%v", s.c.baseURL, *s.dbtTransformationID)
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
		return err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return err
	}

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected: %v", respStatus, expectedStatus)
		return err
	}

	return nil
}

func (s *DbtTransformationDetailsService) Do(ctx context.Context) (DbtTransformationDetailsResponse, error) {
	var response DbtTransformationDetailsResponse

	err := s.do(ctx, &response)

	return response, err
}
