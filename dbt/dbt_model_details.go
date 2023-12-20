package dbt

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type DbtModelDetailsService struct {
	httputils.HttpService
	modelId *string
}

func (s *DbtModelDetailsService) ModelId(value string) *DbtModelDetailsService {
	s.modelId = &value
	return s
}

func (s *DbtModelDetailsService) Do(ctx context.Context) (DbtModelDetailsResponse, error) {
	var response DbtModelDetailsResponse

	if s.modelId == nil {
		return response, fmt.Errorf("missing required modelId")
	}

	url := fmt.Sprintf("/dbt/models/%v", *s.modelId)
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}