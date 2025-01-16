package transformations

import "github.com/fivetran/go-fivetran/utils"

func (elc *TransformationConfig) Request() *transformationConfigRequest {
	return &transformationConfigRequest{
		ProjectId:  	elc.projectId,
		Name:  			elc.name,
		Steps: 			elc.steps,
	}
}

func (elc *TransformationConfig) Merge(customConfig *map[string]interface{}) (*map[string]interface{}, error) {
	err := utils.MergeIntoMap(elc.Request(), customConfig)
	if err != nil {
		return nil, err
	}
	return customConfig, nil
}

func (elc *TransformationConfig) ProjectId(value string) *TransformationConfig {
	elc.projectId = &value
	return elc
}

func (elc *TransformationConfig) Name(value string) *TransformationConfig {
	elc.name = &value
	return elc
}

func (elc *TransformationConfig) Steps(value []TransformationStep) *TransformationConfig {
	elc.steps = &value
	return elc
}
