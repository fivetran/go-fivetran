package dbt

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type DbtTransformationModifyService struct {
	httputils.HttpService
	dbtTransformationId *string
	schedule            *DbtTransformationSchedule
	runTests            *bool
	paused              *bool
}

func (s *DbtTransformationModifyService) request() *dbtTransformationModifyRequest {
	var schedule interface{}

	if s.schedule != nil {
		schedule = s.schedule.Request()
	}

	return &dbtTransformationModifyRequest{
		Schedule: schedule,
		RunTests: s.runTests,
		Paused:   s.paused,
	}
}

func (s *DbtTransformationModifyService) DbtTransformationId(value string) *DbtTransformationModifyService {
	s.dbtTransformationId = &value
	return s
}

func (s *DbtTransformationModifyService) Schedule(value *DbtTransformationSchedule) *DbtTransformationModifyService {
	s.schedule = value
	return s
}

func (s *DbtTransformationModifyService) RunTests(value bool) *DbtTransformationModifyService {
	s.runTests = &value
	return s
}

func (s *DbtTransformationModifyService) Paused(value bool) *DbtTransformationModifyService {
	s.paused = &value
	return s
}

func (s *DbtTransformationModifyService) Do(ctx context.Context) (DbtTransformationResponse, error) {
	var response DbtTransformationResponse
	if s.dbtTransformationId == nil {
		return response, fmt.Errorf("missing required dbtTransformationId")
	}
	url := fmt.Sprintf("/dbt/transformations/%v", *s.dbtTransformationId)
	err := s.HttpService.Do(ctx, "PATCH", url, s.request(), nil, 200, &response)
	return response, err
}