package dbt

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

// ProjectDeleteService implements the Project Management, Delete a project API.
// Ref. https://fivetran.com/docs/rest-api/projects#deleteaproject
type DbtProjectDeleteService struct {
	httputils.HttpService
	dbtProjectID *string
}

func (s *DbtProjectDeleteService) DbtProjectID(value string) *DbtProjectDeleteService {
	s.dbtProjectID = &value
	return s
}

func (s *DbtProjectDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	
	if s.dbtProjectID == nil {
		return response, fmt.Errorf("missing required dbtProjectID")
	}

	url := fmt.Sprintf("/dbt/projects/%v", *s.dbtProjectID)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}