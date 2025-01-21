package transformations

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
	"github.com/fivetran/go-fivetran/utils"
)

type TransformationUpdateService struct {
	httputils.HttpService
	transformationId  				*string
	paused     						*bool
	transformationConfig       		*TransformationConfig
	transformationConfigCustom 		*map[string]interface{}
	transformationSchedule       	*TransformationSchedule
	transformationScheduleCustom 	*map[string]interface{}
}

func (s *TransformationUpdateService) requestBase() transformationUpdateRequestBase {
	return transformationUpdateRequestBase{
		Paused: 			s.paused,
	}
}

func (s *TransformationUpdateService) request() *transformationUpdateRequest {
	var config interface{}
	if s.transformationConfig != nil {
		config = s.transformationConfig.Request()
	}

	var schedule interface{}
	if s.transformationSchedule != nil {
		schedule = s.transformationSchedule.Request()
	}

	r := &transformationUpdateRequest{
		transformationUpdateRequestBase: 		s.requestBase(),
		TransformationConfig:                   config,
		TransformationSchedule:                 schedule,
	}

	return r
}

func (s *TransformationUpdateService) requestCustom() *transformationCustomUpdateRequest {
	return &transformationCustomUpdateRequest{
		transformationUpdateRequestBase: 		s.requestBase(),
		TransformationConfig:                 	s.transformationConfigCustom,
		TransformationSchedule:                 s.transformationScheduleCustom,
	}
}

func (s *TransformationUpdateService) requestCustomMerged() (*transformationCustomUpdateRequest, error) {
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

	return &transformationCustomUpdateRequest{
		transformationUpdateRequestBase: 		s.requestBase(),
		TransformationConfig:                   currentConfig,
		TransformationSchedule:                 currentSchedule,
	}, nil
}

func (s *TransformationUpdateService) TransformationId(value string) *TransformationUpdateService {
	s.transformationId = &value
	return s
}

func (s *TransformationUpdateService) Paused(value bool) *TransformationUpdateService {
	s.paused = &value
	return s
}

func (s *TransformationUpdateService) TransformationConfig(value *TransformationConfig) *TransformationUpdateService {
	s.transformationConfig = value
	return s
}

func (s *TransformationUpdateService) TransformationConfigCustom(value *map[string]interface{}) *TransformationUpdateService {
	s.transformationConfigCustom = value
	return s
}

func (s *TransformationUpdateService) TransformationSchedule(value *TransformationSchedule) *TransformationUpdateService {
	s.transformationSchedule = value
	return s
}

func (s *TransformationUpdateService) TransformationScheduleCustom(value *map[string]interface{}) *TransformationUpdateService {
	s.transformationScheduleCustom = value
	return s
}

func (s *TransformationUpdateService) do(ctx context.Context, req, response any) error {
    if s.transformationId == nil {
        return fmt.Errorf("missing required transformationId")
    }

    url := fmt.Sprintf("/transformations/%v", *s.transformationId)
    err := s.HttpService.Do(ctx, "PATCH", url, req, nil, 200, &response)
    return err
}

func (s *TransformationUpdateService) Do(ctx context.Context) (TransformationResponse, error) {
	var response TransformationResponse

	err := s.do(ctx, s.request(), &response)

	return response, err
}

func (s *TransformationUpdateService) DoCustom(ctx context.Context) (TransformationCustomResponse, error) {
	var response TransformationCustomResponse

	err := s.do(ctx, s.requestCustom(), &response)

	return response, err
}

func (s *TransformationUpdateService) DoCustomMerged(ctx context.Context) (TransformationCustomMergedResponse, error) {
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
