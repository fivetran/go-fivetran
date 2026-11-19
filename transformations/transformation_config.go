package transformations

import "github.com/fivetran/go-fivetran/utils"

func (elc *TransformationConfig) Request() *transformationConfigRequest {
	return &transformationConfigRequest{
		ProjectId:             elc.projectId,
		Name:                  elc.name,
		Steps:                 elc.steps,
		PackageName:           elc.packageName,
		SchemaName:            elc.schemaName,
		ConnectionIds:         elc.connectionIds,
		ExcludedModels:        elc.excludedModels,
		ConfigurableVariables: elc.configurableVariables,
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

func (elc *TransformationConfig) PackageName(value string) *TransformationConfig {
	elc.packageName = &value
	return elc
}

// SchemaName sets the base schema name of a QUICKSTART transformation. Models are written to
// <schema_name>_staging and <schema_name>_reports. Changing it on update flags all models for a
// full refresh.
func (elc *TransformationConfig) SchemaName(value string) *TransformationConfig {
	elc.schemaName = &value
	return elc
}

func (elc *TransformationConfig) ConnectionIds(value []string) *TransformationConfig {
	elc.connectionIds = &value
	return elc
}

func (elc *TransformationConfig) ExcludedModels(value []string) *TransformationConfig {
	elc.excludedModels = &value
	return elc
}

func (elc *TransformationConfig) ConfigurableVariables(value map[string]interface{}) *TransformationConfig {
	elc.configurableVariables = &value
	return elc
}