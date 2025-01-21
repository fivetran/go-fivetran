package transformations

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type TransformationDetailsService struct {
	httputils.HttpService
	transformationId *string
}

func (s *TransformationDetailsService) TransformationId(value string) *TransformationDetailsService {
	s.transformationId = &value
	return s
}

func (s *TransformationDetailsService) Do(ctx context.Context) (TransformationResponse, error) {
	var response TransformationResponse

	if s.transformationId == nil {
		return response, fmt.Errorf("missing required transformationId")
	}

	url := fmt.Sprintf("/transformations/%v", *s.transformationId)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}
