package dbt

import (
	"context"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type DbtTransformationCreateService struct {
	httputils.HttpService
	dbtModelId *string
	schedule   *DbtTransformationSchedule
	runTests   *bool
	paused     *bool
}

func (s *DbtTransformationCreateService) request() *dbtTransformationCreateRequest {
	var schedule interface{}

	if s.schedule != nil {
		schedule = s.schedule.Request()
	}

	return &dbtTransformationCreateRequest{
		DbtModelId: s.dbtModelId,
		Schedule:   schedule,
		RunTests:   s.runTests,
		Paused:     s.paused,
	}
}

func (s *DbtTransformationCreateService) DbtModelId(value string) *DbtTransformationCreateService {
	s.dbtModelId = &value
	return s
}

func (s *DbtTransformationCreateService) Schedule(value *DbtTransformationSchedule) *DbtTransformationCreateService {
	s.schedule = value
	return s
}

func (s *DbtTransformationCreateService) RunTests(value bool) *DbtTransformationCreateService {
	s.runTests = &value
	return s
}

func (s *DbtTransformationCreateService) Paused(value bool) *DbtTransformationCreateService {
	s.paused = &value
	return s
}

func (s *DbtTransformationCreateService) Do(ctx context.Context) (DbtTransformationResponse, error) {
	var response DbtTransformationResponse
	err := s.HttpService.Do(ctx, "POST", "/dbt/transformations", s.request(), nil, 201, &response)
	return response, err
}
