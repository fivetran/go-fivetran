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
		UseLatestPatchVersion:	elc.useLatestPatchVersion,
		ShouldRunDbtDeps:	elc.shouldRunDbtDeps,
	}
}

func (elc *TransformationProjectConfig) UpdateRequest() *transformationProjectConfigUpdateRequest {
	return &transformationProjectConfigUpdateRequest{
		FolderPath:     	elc.folderPath,
		GitBranch:   		elc.gitBranch,
		TargetName:     	elc.targetName,
		EnvironmentVars:	elc.environmentVars,
		Threads:    		elc.threads,
		UseLatestPatchVersion:	elc.useLatestPatchVersion,
		ShouldRunDbtDeps:	elc.shouldRunDbtDeps,
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

// UseLatestPatchVersion sets whether the project runs on the latest patch version of its dbt minor version.
// Only the accounts that Fivetran doesn't keep on the latest patch version can set it to false,
// and setting it to false requires a full dbt version, for example 1.11.7.
func (elc *TransformationProjectConfig) UseLatestPatchVersion(value bool) *TransformationProjectConfig {
	elc.useLatestPatchVersion = &value
	return elc
}

// ShouldRunDbtDeps sets whether `dbt deps` installs the dbt packages before each transformation run.
func (elc *TransformationProjectConfig) ShouldRunDbtDeps(value bool) *TransformationProjectConfig {
	elc.shouldRunDbtDeps = &value
	return elc
}
