package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// DbtTransformationCreateService implements the dbt transformation management,
// create a dbt transformation API
// Ref. POST https://api.fivetran.com/v1/dbt/projects/{project_id}/transformations

type DbtTransformationCreateService struct {
	c            *Client
	dbtProjectID *string
	dbtModelID   *string
	schedule     *DbtTransformationSchedule
	runTests     *bool
}

type dbtTransformationCreateRequestBase struct {
	DbtModelID *string `json:"dbt_model_id,omitempty"`
	RunTests   *bool   `json:"run_tests,omitempty`
}

type dbtTransformationCreateRequest struct {
	dbtTransformationCreateRequestBase
	Schedule *dbtTransformationScheduleRequest `json:"schedule,omitempty"`
}

type DbtTransformationCreateResponseBase struct {
	ID              string    `json:"id"`
	DbtModelID      string    `json:"dbt_model_id"`
	OutputModelName string    `json:"output_model_name"`
	DbtProjectID    string    `json:"dbt_project_id"`
	LastRun         time.Time `json:"last_run"`
	NextRun         time.Time `json:"next_run"`
	Status          string    `json:"status"`
	RunTests        bool      `json:"run_tests"`
	ConnectorIDs    []string  `json:"connector_ids"`
	ModelIDs        []string  `json:"model_ids"`
}

type DbtTransformationCreateResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		DbtTransformationCreateResponseBase
		Schedule DbtTransformationScheduleResponse
	} `json:"data"`
}

func (c *Client) NewDbtTransformationCreateService() *DbtTransformationCreateService {
	return &DbtTransformationCreateService{c: c}
}

func (s *DbtTransformationCreateService) requestBase() dbtTransformationCreateRequestBase {
	return dbtTransformationCreateRequestBase{
		DbtModelID: s.dbtModelID,
		RunTests:   s.runTests,
	}
}

func (s *DbtTransformationCreateService) request() *dbtTransformationCreateRequest {
	var schedule *dbtTransformationScheduleRequest
	if s.schedule != nil {
		schedule = s.schedule.request()
	}

	r := &dbtTransformationCreateRequest{
		dbtTransformationCreateRequestBase: s.requestBase(),
		Schedule:                           schedule,
	}

	return r
}

func (s *DbtTransformationCreateService) DbtProjectID(value string) *DbtTransformationCreateService {
	s.dbtProjectID = &value
	return s
}

func (s *DbtTransformationCreateService) DbtModelID(value string) *DbtTransformationCreateService {
	s.dbtModelID = &value
	return s
}

func (s *DbtTransformationCreateService) Schedule(value *DbtTransformationSchedule) *DbtTransformationCreateService {
	s.schedule = value
	return s
}

func (s *DbtTransformationCreateService) RunTests(value bool) *DbtTransformationCreateService {
	s.runTests = &value
	return s
}

func (s *DbtTransformationCreateService) do(ctx context.Context, req, response any) error {
	url := fmt.Sprintf("%v/dbt/projects/%v/transformations", s.c.baseURL, *s.dbtProjectID)
	expectedStatus := 201

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"
	headers["Accept"] = restAPIv2

	reqBody, err := json.Marshal(req)
	if err != nil {
		return err
	}

	r := request{
		method:  "POST",
		url:     url,
		body:    reqBody,
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

	if respStatus != expectedStatus {
		err := fmt.Errorf("status code: %v; expected: %v", respStatus, expectedStatus)
		return err
	}

	return nil
}

func (s *DbtTransformationCreateService) Do(ctx context.Context) (DbtTransformationCreateResponse, error) {
	var response DbtTransformationCreateResponse

	err := s.do(ctx, s.request(), &response)

	return response, err
}
