package transformations

import "github.com/fivetran/go-fivetran/common"

type transformationProjectResponseBase struct {
    Id    		string `json:"id,omitempty"`
    ProjectType string `json:"type,omitempty"`
    CreatedAt 	string `json:"created_at,omitempty"`
    GroupId 	string `json:"group_id,omitempty"`
    CreatedById string `json:"created_by_id,omitempty"`
}

type TransformationProjectConfig struct {
	dbtVersion  	*string
	defaultSchema   *string
	gitRemoteUrl 	*string
	folderPath      *string
	gitBranch   	*string
	targetName      *string
	environmentVars	*string
	threads         *int
}

type transformationProjectConfigResponse struct {
	DbtVersion  	string `json:"dbt_version,omitempty"`
	DefaultSchema   string `json:"default_schema,omitempty"`
	GitRemoteUrl 	string `json:"git_remote_url,omitempty"`
	FolderPath      string `json:"folder_path,omitempty"`
	GitBranch   	string `json:"git_branch,omitempty"`
	TargetName      string `json:"target_name,omitempty"`
	EnvironmentVars	string `json:"environment_vars,omitempty"`
	Threads         int    `json:"threads,omitempty"`
}

type transformationProjectConfigCreateRequest struct {
	DbtVersion  	*string `json:"dbt_version,omitempty"`
	DefaultSchema   *string `json:"default_schema,omitempty"`
	GitRemoteUrl 	*string `json:"git_remote_url,omitempty"`
	FolderPath      *string `json:"folder_path,omitempty"`
	GitBranch   	*string `json:"git_branch,omitempty"`
	TargetName      *string `json:"target_name,omitempty"`
	EnvironmentVars	*string `json:"environment_vars,omitempty"`
	Threads         *int    `json:"threads,omitempty"`
}

type transformationProjectConfigUpdateRequest struct {
	FolderPath      *string `json:"folder_path,omitempty"`
	GitBranch   	*string `json:"git_branch,omitempty"`
	TargetName      *string `json:"target_name,omitempty"`
	EnvironmentVars	*string `json:"environment_vars,omitempty"`
	Threads         *int    `json:"threads,omitempty"`
}

type transformationProjectResponse struct {
	transformationProjectResponseBase
    Status 			string 								`json:"status,omitempty"`
    SetupTests 		[]common.SetupTestResponse  		`json:"setup_tests,omitempty"`
    Errors 			[]string							`json:"errors,omitempty"`
}

type transformationProjectCreateRequestBase struct {
	Id      		*string `json:"id,omitempty"`
	GroupId 		*string `json:"group_id,omitempty"`
	ProjectType 	*string `json:"type,omitempty"`
	RunTests 		*bool   `json:"run_tests,omitempty"`
}

type transformationProjectCreateRequest struct {
	transformationProjectCreateRequestBase
	ProjectConfig any `json:"project_config,omitempty"`
}

type transformationProjectCustomCreateRequest struct {
	transformationProjectCreateRequestBase
	ProjectConfig *map[string]interface{} `json:"project_config,omitempty"`
}

type transformationProjectUpdateRequestBase struct {
	RunTests *bool `json:"run_tests,omitempty"`
}

type transformationProjectUpdateRequest struct {
	transformationProjectUpdateRequestBase
	ProjectConfig any `json:"project_config,omitempty"`
}

type transformationProjectCustomUpdateRequest struct {
	transformationProjectUpdateRequestBase
	ProjectConfig *map[string]interface{} `json:"project_config,omitempty"`
}

type TransformationProjectsListResponse struct {
	common.CommonResponse
	Data struct {
		Items      []transformationProjectResponseBase `json:"items"`
		NextCursor string                              `json:"next_cursor"`
	} `json:"data"`
}

type TransformationProjectResponse struct {
	common.CommonResponse
	Data 	struct {
		transformationProjectResponse 			
		ProjectConfig 	transformationProjectConfigResponse `json:"project_config,omitempty"`
	}`json:"data"`
}

type TransformationProjectCustomResponse struct {
	common.CommonResponse
	Data 	struct {
		transformationProjectResponse 			
		ProjectConfig 	map[string]interface{} `json:"project_config"`
	}`json:"data"`
}

type TransformationProjectCustomMergedResponse struct {
	common.CommonResponse
	Data 	struct {
		transformationProjectResponse 			
		ProjectConfig 			transformationProjectConfigResponse // no mapping here
		ProjectConfigCustom 	map[string]interface{} `json:"project_config"`
	}`json:"data"`
}