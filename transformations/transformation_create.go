package transformations

import (
	"context"

	httputils "github.com/fivetran/go-fivetran/http_utils"
	"github.com/fivetran/go-fivetran/utils"
)

type TransformationCreateService struct {
	httputils.HttpService
	projectType    					*string
	paused     						*bool
	transformationConfig       		*TransformationConfig
	transformationConfigCustom 		*map[string]interface{}
	transformationSchedule       	*TransformationSchedule
	transformationScheduleCustom 	*map[string]interface{}
}

func (s *TransformationCreateService) requestBase() transformationCreateRequestBase {
	return transformationCreateRequestBase{
		ProjectType: 		s.projectType,
		Paused: 			s.paused,
	}
}

func (s *TransformationCreateService) request() *transformationCreateRequest {
	var config interface{}
	if s.transformationConfig != nil {
		config = s.transformationConfig.Request()
	}

	var schedule interface{}
	if s.transformationSchedule != nil {
		schedule = s.transformationSchedule.Request()
	}

	r := &transformationCreateRequest{
		transformationCreateRequestBase: 		s.requestBase(),
		TransformationConfig:                   config,
		TransformationSchedule:                 schedule,
	}

	return r
}

func (s *TransformationCreateService) requestCustom() *transformationCustomCreateRequest {
	return &transformationCustomCreateRequest{
		transformationCreateRequestBase: 		s.requestBase(),
		TransformationConfig:                 	s.transformationConfigCustom,
		TransformationSchedule:                 s.transformationScheduleCustom,
	}
}

func (s *TransformationCreateService) requestCustomMerged() (*transformationCustomCreateRequest, error) {
	currentConfig := s.transformationConfigCustom

	if s.transformationConfig != nil {
		var err error
		currentConfig, err = s.transformationConfig.Merge(currentConfig)
		if err != nil {
			return nil, err
		}
	}

	currentSchedule := s.transformationScheduleCustom

	if s.transformationSchedule != nil {
		var err error
		currentSchedule, err = s.transformationSchedule.Merge(currentSchedule)
		if err != nil {
			return nil, err
		}
	}

	return &transformationCustomCreateRequest{
		transformationCreateRequestBase: 		s.requestBase(),
		TransformationConfig:                   currentConfig,
		TransformationSchedule:                 currentSchedule,
	}, nil
}

func (s *TransformationCreateService) ProjectType(value string) *TransformationCreateService {
	s.projectType = &value
	return s
}

func (s *TransformationCreateService) Paused(value bool) *TransformationCreateService {
	s.paused = &value
	return s
}

func (s *TransformationCreateService) TransformationConfig(value *TransformationConfig) *TransformationCreateService {
	s.transformationConfig = value
	return s
}

func (s *TransformationCreateService) TransformationConfigCustom(value *map[string]interface{}) *TransformationCreateService {
	s.transformationConfigCustom = value
	return s
}

func (s *TransformationCreateService) TransformationSchedule(value *TransformationSchedule) *TransformationCreateService {
	s.transformationSchedule = value
	return s
}

func (s *TransformationCreateService) TransformationScheduleCustom(value *map[string]interface{}) *TransformationCreateService {
	s.transformationScheduleCustom = value
	return s
}

func (s *TransformationCreateService) do(ctx context.Context, req, response any) error {
	err := s.HttpService.Do(ctx, "POST", "/transformations", req, nil, 201, &response)
	return err
}

func (s *TransformationCreateService) Do(ctx context.Context) (TransformationResponse, error) {
	var response TransformationResponse

	err := s.do(ctx, s.request(), &response)

	return response, err
}

func (s *TransformationCreateService) DoCustom(ctx context.Context) (TransformationCustomResponse, error) {
	var response TransformationCustomResponse

	err := s.do(ctx, s.requestCustom(), &response)

	return response, err
}

func (s *TransformationCreateService) DoCustomMerged(ctx context.Context) (TransformationCustomMergedResponse, error) {
	var response TransformationCustomMergedResponse

	req, err := s.requestCustomMerged()

	if err != nil {
		return response, err
	}

	err = s.do(ctx, req, &response)

	if err == nil {
		err = utils.FetchFromMap(&response.Data.TransformationConfigCustom, &response.Data.TransformationConfig)
		if err == nil {
			err = utils.FetchFromMap(&response.Data.TransformationScheduleCustom, &response.Data.TransformationSchedule)
		}
	}

	return response, err
}
