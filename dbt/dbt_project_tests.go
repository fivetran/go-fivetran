package dbt

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// DbtProjectTestService implements the Transformations for dbt Core Management, Test dbt Project.
// Ref. https://fivetran.com/docs/rest-api/transformations-for-dbt-core-management#request_4
type DbtProjectTestService struct {
	httputils.HttpService
	projectID     *string
}

func (s *DbtProjectTestService) ProjectID(value string) *DbtProjectTestService {
	s.projectID = &value
	return s
}

func (s *DbtProjectTestService) Do(ctx context.Context) (DbtProjectTestResponse, error) {
	var response DbtProjectTestResponse

	if s.projectID == nil {
		return response, fmt.Errorf("missing required projectID")
	}

	url := fmt.Sprintf("/dbt/projects/%v/test", *s.projectID)
	err := s.HttpService.Do(ctx, "POST", url, nil, nil, 200, &response)
	return response, err
}