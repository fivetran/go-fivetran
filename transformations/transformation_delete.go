package transformations

import (
	"context"
	"fmt"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type TransformationDeleteService struct {
	httputils.HttpService
	transformationId *string
}

func (s *TransformationDeleteService) TransformationId(value string) *TransformationDeleteService {
	s.transformationId = &value
	return s
}

func (s *TransformationDeleteService) Do(ctx context.Context) (common.CommonResponse, error) {
	var response common.CommonResponse
	
	if s.transformationId == nil {
		return response, fmt.Errorf("missing required transformationId")
	}

	url := fmt.Sprintf("/transformations/%v", *s.transformationId)
	err := s.HttpService.Do(ctx, "DELETE", url, nil, nil, 200, &response)
	return response, err
}
