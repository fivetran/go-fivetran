package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type DbtProjectDetailsService struct {
	c            *Client
	dbtProjectID *string
}

type DbtProjectDetailsdataBase struct {
	ID            string `json:"id"`
	GroupID       string `json:"group_id"`
	CreatedAt     string `json:"created_at"`
	CreatedById   string `json:"created_by_id"`
	PublicKey     string `json:"public_key"`
	GitRemoteUrl  string `json:"git_remote_url"`
	GitBranch     string `json:"git_branch"`
	DefaultSchema string `json:"default_schema"`
	FolderPath    string `json:"folder_path"`
	TargetName    string `json:"target_name"`
}

type DbtProjectDetailsResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		DbtProjectDetailsdataBase
	} `json:"data"`
}

func (c *Client) NewDbtDetails() *DbtProjectDetailsService {
	return &DbtProjectDetailsService{c: c}
}

func (s *DbtProjectDetailsService) DbtProjectID(value string) *DbtProjectDetailsService {
	s.dbtProjectID = &value
	return s
}

func (s *DbtProjectDetailsService) do(ctx context.Context, response any) error {
	if s.dbtProjectID == nil {
		return fmt.Errorf("missing required DbtProjectId")
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

func (s *DbtProjectDetailsService) Do(ctx context.Context) (DbtProjectDetailsResponse, error) {
	var response DbtProjectDetailsResponse

	err := s.do(ctx, &response)

	return response, err
}
