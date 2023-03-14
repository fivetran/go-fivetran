package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

// DbtProjectDeleteService implements the dbt management, delete a dbt project api.
// Ref. POST https://api.fivetran.com/v1/dbt/projects/{id}
type DbtProjectDeleteService struct {
	c     *Client
	dbtID *string
}

type DbtProjectDeleteResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (c *Client) NewDbtProjectDeleteService() *DbtProjectDeleteService {
	return &DbtProjectDeleteService{c: c}
}

func (s *DbtProjectDeleteService) DbtID(value string) *DbtProjectDeleteService {
	s.dbtID = &value
	return s
}

func (s *DbtProjectDeleteService) Do(ctx context.Context) (DbtProjectDeleteResponse, error) {
	var response DbtProjectDeleteResponse

	if s.dbtID == nil {
		return response, fmt.Errorf("missing required DbtID")
	}

	url := fmt.Sprintf("%v/dbt/projects/%v", s.c.baseURL, *s.dbtID)
	expectedStatus := 200

	headers := s.c.commonHeaders()

	r := request{
		method:  "DELETE",
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
