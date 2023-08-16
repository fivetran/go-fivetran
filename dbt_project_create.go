package fivetran

import (
	"context"
	"encoding/json"
	"fmt"
)

type DbtProjectCreateService struct {
	c             *Client
	groupID       *string
	dbtVersion    *string
	gitRemoteUrl  *string
	gitBranch     *string
	defaultSchema *string
	folderPath    *string
	targetName    *string
	threads       *int
}

type dbtProjectCreateRequest struct {
	GroupID       *string `json:"group_id,omitempty"`
	DbtVersion    *string `json:"dbt_version,omitempty"`
	GitRemoteUrl  *string `json:"git_remote_url,omitempty"`
	GitBranch     *string `json:"git_branch,omitempty"`
	DefaultSchema *string `json:"default_schema,omitempty"`
	FolderPath    *string `json:"folder_path,omitempty"`
	TargetName    *string `json:"target_name,omitempty"`
	Threads       *int    `json:"threads,omitempty"`
}

type DbtProjectCreateResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID            string `json:"id"`
		GroupID       string `json:"group_id"`
		FolderPath    string `json:"folder_path"`
		CreatedAt     string `json:"created_at"`
		TargetName    string `json:"target_name"`
		GitRemoteUrl  string `json:"git_remote_url"`
		DefaultSchema string `json:"default_schema"`
		PublicKey     string `json:"public_key"`
		CreatedById   string `json:"created_by_id"`
		GitBranch     string `json:"git_branch"`
	} `json:"data"`
}

func (c *Client) NewDbtProjectCreate() *DbtProjectCreateService {
	return &DbtProjectCreateService{c: c}
}

func (s *DbtProjectCreateService) request() *dbtProjectCreateRequest {
	return &dbtProjectCreateRequest{
		GroupID:       s.groupID,
		DbtVersion:    s.dbtVersion,
		GitRemoteUrl:  s.gitRemoteUrl,
		GitBranch:     s.gitBranch,
		DefaultSchema: s.defaultSchema,
		FolderPath:    s.folderPath,
		TargetName:    s.targetName,
		Threads:       s.threads,
	}
}

func (s *DbtProjectCreateService) GroupID(value string) *DbtProjectCreateService {
	s.groupID = &value
	return s
}

func (s *DbtProjectCreateService) DbtVersion(value string) *DbtProjectCreateService {
	s.dbtVersion = &value
	return s
}

func (s *DbtProjectCreateService) GitRemoteUrl(value string) *DbtProjectCreateService {
	s.gitRemoteUrl = &value
	return s
}

func (s *DbtProjectCreateService) GitBranch(value string) *DbtProjectCreateService {
	s.gitBranch = &value
	return s
}

func (s *DbtProjectCreateService) DefaultSchema(value string) *DbtProjectCreateService {
	s.defaultSchema = &value
	return s
}

func (s *DbtProjectCreateService) FolderPath(value string) *DbtProjectCreateService {
	s.folderPath = &value
	return s
}

func (s *DbtProjectCreateService) TargetName(value string) *DbtProjectCreateService {
	s.targetName = &value
	return s
}

func (s *DbtProjectCreateService) Threads(value int) *DbtProjectCreateService {
	s.threads = &value
	return s
}

func (s *DbtProjectCreateService) Do(ctx context.Context) (DbtProjectCreateResponse, error) {
	var response DbtProjectCreateResponse
	url := fmt.Sprintf("%v/dbt/projects", s.c.baseURL)
	expectedStatus := 201

	headers := s.c.commonHeaders()
	headers["Content-Type"] = "application/json"

	reqBody, err := json.Marshal(s.request())
	if err != nil {
		return response, err
	}

	r := request{
		method:           "POST",
		url:              url,
		body:             reqBody,
		queries:          nil,
		headers:          headers,
		client:           s.c.httpClient,
		handleRateLimits: s.c.handleRateLimits,
		maxRetryAttempts: s.c.maxRetryAttempts,
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
