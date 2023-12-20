package dbt

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type DbtProjectDetailsService struct {
	httputils.HttpService
	dbtProjectID *string
}

func (s *DbtProjectDetailsService) DbtProjectID(value string) *DbtProjectDetailsService {
	s.dbtProjectID = &value
	return s
}

func (s *DbtProjectDetailsService) Do(ctx context.Context) (DbtProjectDetailsResponse, error) {
	var response DbtProjectDetailsResponse
	if s.dbtProjectID == nil {
		return response, fmt.Errorf("missing required dbtProjectID")
	}
	url := fmt.Sprintf("/dbt/projects/%v", *s.dbtProjectID)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}