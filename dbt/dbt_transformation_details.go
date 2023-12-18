package dbt

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type DbtTransformationDetailsService struct {
	httputils.HttpService
	transformationId *string
}

func (s *DbtTransformationDetailsService) TransformationId(value string) *DbtTransformationDetailsService {
	s.transformationId = &value
	return s
}

func (s *DbtTransformationDetailsService) Do(ctx context.Context) (DbtTransformationResponse, error) {
	var response DbtTransformationResponse

	if s.transformationId == nil {
		return response, fmt.Errorf("missing required transformationId")
	}

	url := fmt.Sprintf("/dbt/transformations/%v", *s.transformationId)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}
