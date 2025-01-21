package transformations

import "github.com/fivetran/go-fivetran/utils"

func (elc *TransformationProjectConfig) CreateRequest() *transformationProjectConfigCreateRequest {
	return &transformationProjectConfigCreateRequest{
		DbtVersion:  		elc.dbtVersion,
		DefaultSchema:  	elc.defaultSchema,
		GitRemoteUrl: 		elc.gitRemoteUrl,
		FolderPath:     	elc.folderPath,
		GitBranch:   		elc.gitBranch,
		TargetName:     	elc.targetName,
		EnvironmentVars:	elc.environmentVars,
		Threads:    		elc.threads,
	}
}

func (elc *TransformationProjectConfig) UpdateRequest() *transformationProjectConfigUpdateRequest {
	return &transformationProjectConfigUpdateRequest{
		FolderPath:     	elc.folderPath,
		GitBranch:   		elc.gitBranch,
		TargetName:     	elc.targetName,
		EnvironmentVars:	elc.environmentVars,
		Threads:    		elc.threads,
	}
}

func (elc *TransformationProjectConfig) MergeForCreate(customConfig *map[string]interface{}) (*map[string]interface{}, error) {
	err := utils.MergeIntoMap(elc.CreateRequest(), customConfig)
	if err != nil {
		return nil, err
	}
	return customConfig, nil
}

func (elc *TransformationProjectConfig) MergeForUpdate(customConfig *map[string]interface{}) (*map[string]interface{}, error) {
	err := utils.MergeIntoMap(elc.UpdateRequest(), customConfig)
	if err != nil {
		return nil, err
	}
	return customConfig, nil
}

func (elc *TransformationProjectConfig) DbtVersion(value string) *TransformationProjectConfig {
	elc.dbtVersion = &value
	return elc
}

func (elc *TransformationProjectConfig) DefaultSchema(value string) *TransformationProjectConfig {
	elc.defaultSchema = &value
	return elc
}

func (elc *TransformationProjectConfig) GitRemoteUrl(value string) *TransformationProjectConfig {
	elc.gitRemoteUrl = &value
	return elc
}

func (elc *TransformationProjectConfig) FolderPath(value string) *TransformationProjectConfig {
	elc.folderPath = &value
	return elc
}

func (elc *TransformationProjectConfig) GitBranch(value string) *TransformationProjectConfig {
	elc.gitBranch = &value
	return elc
}

func (elc *TransformationProjectConfig) TargetName(value string) *TransformationProjectConfig {
	elc.targetName = &value
	return elc
}

func (elc *TransformationProjectConfig) EnvironmentVars(value []string) *TransformationProjectConfig {
	elc.environmentVars = &value
	return elc
}

func (elc *TransformationProjectConfig) Threads(value int) *TransformationProjectConfig {
	elc.threads = &value
	return elc
}
