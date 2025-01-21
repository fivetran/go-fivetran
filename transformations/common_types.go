package transformations

import "github.com/fivetran/go-fivetran/common"

/* Transformation Projects*/
type transformationProjectResponseBase struct {
    Id          string `json:"id,omitempty"`
    ProjectType string `json:"type,omitempty"`
    CreatedAt   string `json:"created_at,omitempty"`
    GroupId     string `json:"group_id,omitempty"`
    CreatedById string `json:"created_by_id,omitempty"`
}

type TransformationProjectConfig struct {
    dbtVersion      *string
    defaultSchema   *string
    gitRemoteUrl    *string
    folderPath      *string
    gitBranch       *string
    targetName      *string
    environmentVars *[]string
    threads         *int
}

type transformationProjectConfigResponse struct {
    DbtVersion      string `json:"dbt_version,omitempty"`
    DefaultSchema   string `json:"default_schema,omitempty"`
    GitRemoteUrl    string `json:"git_remote_url,omitempty"`
    FolderPath      string `json:"folder_path,omitempty"`
    GitBranch       string `json:"git_branch,omitempty"`
    TargetName      string `json:"target_name,omitempty"`
    EnvironmentVars []string `json:"environment_vars,omitempty"`
    PublicKey       string `json:"public_key,omitempty"`
    Threads         int    `json:"threads,omitempty"`
}

type transformationProjectConfigCreateRequest struct {
    DbtVersion      *string `json:"dbt_version,omitempty"`
    DefaultSchema   *string `json:"default_schema,omitempty"`
    GitRemoteUrl    *string `json:"git_remote_url,omitempty"`
    FolderPath      *string `json:"folder_path,omitempty"`
    GitBranch       *string `json:"git_branch,omitempty"`
    TargetName      *string `json:"target_name,omitempty"`
    EnvironmentVars *[]string `json:"environment_vars,omitempty"`
    Threads         *int    `json:"threads,omitempty"`
}

type transformationProjectConfigUpdateRequest struct {
    FolderPath      *string `json:"folder_path,omitempty"`
    GitBranch       *string `json:"git_branch,omitempty"`
    TargetName      *string `json:"target_name,omitempty"`
    EnvironmentVars *[]string `json:"environment_vars,omitempty"`
    Threads         *int    `json:"threads,omitempty"`
}

type transformationProjectResponse struct {
    transformationProjectResponseBase
    Status          string                              `json:"status,omitempty"`
    SetupTests      []common.SetupTestResponse          `json:"setup_tests,omitempty"`
    Errors          []string                            `json:"errors,omitempty"`
}

type transformationProjectCreateRequestBase struct {
    Id              *string `json:"id,omitempty"`
    GroupId         *string `json:"group_id,omitempty"`
    ProjectType     *string `json:"type,omitempty"`
    RunTests        *bool   `json:"run_tests,omitempty"`
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
    Data    struct {
        transformationProjectResponse           
        ProjectConfig   transformationProjectConfigResponse `json:"project_config,omitempty"`
    }`json:"data"`
}

type TransformationProjectCustomResponse struct {
    common.CommonResponse
    Data    struct {
        transformationProjectResponse           
        ProjectConfig   map[string]interface{} `json:"project_config"`
    }`json:"data"`
}

type TransformationProjectCustomMergedResponse struct {
    common.CommonResponse
    Data    struct {
        transformationProjectResponse           
        ProjectConfig           transformationProjectConfigResponse // no mapping here
        ProjectConfigCustom     map[string]interface{} `json:"project_config"`
    }`json:"data"`
}

/* Transformations */

type transformationCreateRequestBase struct {
    ProjectType     *string `json:"type,omitempty"`
    Paused          *bool   `json:"paused,omitempty"`
}

type TransformationSchedule struct {
    cron            *[]string
    connectionIds   *[]string
    daysOfWeek      *[]string
    timeOfDay       *string
    scheduleType    *string
    interval        *int
    smartSyncing    *bool
}

type transformationScheduleRequest struct {
    Cron            *[]string `json:"cron,omitempty"`
    ConnectionIds   *[]string `json:"connection_ids,omitempty"`
    DaysOfWeek      *[]string `json:"days_of_week,omitempty"`
    TimeOfDay       *string   `json:"time_of_day,omitempty"`
    ScheduleType    *string   `json:"schedule_type,omitempty"`
    Interval        *int      `json:"interval,omitempty"`
    SmartSyncing    *bool     `json:"smart_syncing,omitempty"`
}

type transformationScheduleResponse struct {
    Cron            []string `json:"cron,omitempty"`
    ConnectionIds   []string `json:"connection_ids,omitempty"`
    DaysOfWeek      []string `json:"days_of_week,omitempty"`
    TimeOfDay       string   `json:"time_of_day,omitempty"`
    ScheduleType    string   `json:"schedule_type,omitempty"`
    Interval        int      `json:"interval,omitempty"`
    SmartSyncing    bool     `json:"smart_syncing,omitempty"`
}

type TransformationConfig struct {
    /* DBT_CORE */
    projectId       *string
    name            *string
    steps           *[]TransformationStep
    /* QUICKSTART */
    packageName     *string
    connectionIds   *[]string
    excludedModels  *[]string 
}

type transformationConfigRequest struct {
    /* DBT_CORE */
    ProjectId       *string `json:"project_id,omitempty"`
    Name            *string `json:"name,omitempty"`
    Steps           *[]TransformationStep `json:"steps,omitempty"`
    /* QUICKSTART */
    PackageName     *string `json:"package_name,omitempty"`
    ConnectionIds   *[]string `json:"connection_ids,omitempty"`
    ExcludedModels  *[]string `json:"excluded_models,omitempty"`
}

type TransformationStep struct {
    Name        string `json:"name,omitempty"`
    Command     string `json:"command,omitempty"` 
}

type transformationConfigResponse struct {
    /* DBT_CORE */
    ProjectId           string               `json:"project_id,omitempty"`
    Name                string               `json:"name,omitempty"`
    Steps               []TransformationStep `json:"steps,omitempty"`
    /* QUICKSTART */
    PackageName         string               `json:"package_name,omitempty"`
    ConnectionIds       []string             `json:"connection_ids,omitempty"`
    ExcludedModels      []string             `json:"excluded_models,omitempty"`
    UpgradeAvailable    bool               `json:"upgrade_available,omitempty"`
}

type transformationCreateRequest struct {
    transformationCreateRequestBase
    TransformationSchedule  any     `json:"schedule,omitempty"`
    TransformationConfig    any     `json:"transformation_config,omitempty"`
}

type transformationCustomCreateRequest struct {
    transformationCreateRequestBase
    TransformationSchedule  *map[string]interface{}     `json:"schedule,omitempty"`
    TransformationConfig    *map[string]interface{}     `json:"transformation_config,omitempty"`
}

type transformationResponse struct {
    transformationResponseBase
    Status          string                              `json:"status,omitempty"`
    SetupTests      []common.SetupTestResponse          `json:"setup_tests,omitempty"`
    Errors          []string                            `json:"errors,omitempty"`
}

type transformationResponseBase struct {
    Id                  string      `json:"id,omitempty"`
    ProjectType         string      `json:"type,omitempty"`
    Status              string      `json:"status,omitempty"`
    Paused              bool        `json:"paused,omitempty"`
    CreatedAt           string      `json:"created_at,omitempty"`
    CreatedById         string      `json:"created_by_id,omitempty"`
    OutputModelNames    []string    `json:"output_model_names,omitempty"`
}

type TransformationResponse struct {
    common.CommonResponse
    Data    struct {
        transformationResponseBase          
        TransformationConfig    transformationConfigResponse `json:"transformation_config,omitempty"`
        TransformationSchedule  transformationScheduleResponse `json:"schedule,omitempty"`
    }`json:"data"`
}

type TransformationsListResponse struct {
    common.CommonResponse
    Data struct {
        Items      []struct {
            transformationResponseBase          
            TransformationConfig    transformationConfigResponse `json:"transformation_config,omitempty"`
            TransformationSchedule  transformationScheduleResponse `json:"schedule,omitempty"`
        } `json:"items"`
        NextCursor string                              `json:"next_cursor"`
    } `json:"data"`
}

type TransformationCustomResponse struct {
    common.CommonResponse
    Data    struct {
        transformationResponseBase          
        TransformationConfig    map[string]interface{} `json:"transformation_config"`
        TransformationSchedule  map[string]interface{} `json:"schedule"`
    }`json:"data"`
}

type TransformationCustomMergedResponse struct {
    common.CommonResponse
    Data    struct {
        transformationResponseBase          
        TransformationConfig            transformationConfigResponse // no mapping here
        TransformationConfigCustom      map[string]interface{} `json:"transformation_config"`
        TransformationSchedule          transformationScheduleResponse // no mapping here
        TransformationScheduleCustom    map[string]interface{} `json:"schedule"`
    }`json:"data"`
}

type transformationUpdateRequestBase struct {
    Paused          *bool   `json:"paused,omitempty"`
}

type transformationUpdateRequest struct {
    transformationUpdateRequestBase
    TransformationSchedule  any     `json:"schedule,omitempty"`
    TransformationConfig    any     `json:"transformation_config,omitempty"`
}

type transformationCustomUpdateRequest struct {
    transformationUpdateRequestBase
    TransformationSchedule  *map[string]interface{}     `json:"schedule,omitempty"`
    TransformationConfig    *map[string]interface{}     `json:"transformation_config,omitempty"`
}

/* Quickstart metadata details*/
type quickstartPackageResponseBase struct {
    Id                  string      `json:"id,omitempty"`
    Name                string      `json:"name,omitempty"`
    Version             string      `json:"version,omitempty"`
    ConnectorTypes      []string      `json:"connector_types,omitempty"`
    OutputModelNames    []string    `json:"output_model_names,omitempty"`
}

type QuickstartPackageResponse struct {
    common.CommonResponse
    Data    struct {
        quickstartPackageResponseBase          
    }`json:"data"`
}

type QuickstartPackagesListResponse struct {
    common.CommonResponse
    Data struct {
        Items      []quickstartPackageResponseBase `json:"items"`
        NextCursor string                          `json:"next_cursor"`
    } `json:"data"`
}
