package dbt

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type DbtTransformationDeleteService struct {
	httputils.HttpService
	transformationId *string
}

func (s *DbtTransformationDeleteService) TransformationId(value string) *DbtTransformationDeleteService {
	s.transformationId = &value
	return s
}

func (s *DbtTransformationDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	
	if s.transformationId == nil {
		return response, fmt.Errorf("missing required transformationId")
	}

	url := fmt.Sprintf("/dbt/transformations/%v", *s.transformationId)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}
