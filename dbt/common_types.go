package dbt

import "github.com/fivetran/go-fivetran/common"

type DbtModelItem struct {
	ID        string `json:"id"`
	ModelName string `json:"model_name"`
	Scheduled bool   `json:"scheduled"`
}

type DbtModelDetailsResponse struct {
	common.CommonResponse
	Data DbtModelItem `json:"data"`
}

type DbtModelsListResponse struct {
	common.CommonResponse
	Data struct {
		Items      []DbtModelItem `json:"items"`
		NextCursor string         `json:"next_cursor"`
	} `json:"data"`
}

type DbtProjectItem struct {
	ID          string `json:"id"`
	GroupId     string `json:"group_id"`
	CreatedAt   string `json:"created_at"`
	CreatedById string `json:"created_by_id"`
}

type DbtTransformationResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID              string                            `json:"id"`
		Status          string                            `json:"status"`
		Schedule        DbtTransformationScheduleResponse `json:"schedule"`
		LastRun         string                            `json:"last_run"`
		OutputModelName string                            `json:"output_model_name"`
		DbtProjectId    string                            `json:"dbt_project_id"`
		DbtModelId      string                            `json:"dbt_model_id"`
		NextRun         string                            `json:"next_run"`
		CreatedAt       string                            `json:"created_at"`
		ModelIds        []string                          `json:"model_ids"`
		ConnectorIds    []string                          `json:"connector_ids"`
		RunTests        bool                              `json:"run_tests"`
		Paused          bool                              `json:"paused"`
	} `json:"data"`
}

type DbtProjectDetailsData struct {
	DbtProjectItem
	DbtVersion      string                   `json:"dbt_version"`
	PublicKey       string                   `json:"public_key"`
	DefaultSchema   string                   `json:"default_schema"`
	TargetName      string                   `json:"target_name"`
	Threads         int                      `json:"threads"`
	EnvironmentVars []string                 `json:"environment_vars"`
	Type            string                   `json:"type"`
	ProjectConfig   DbtProjectConfigResponse `json:"project_config"`
	Status          string                   `json:"status"`
	Errors          []string                 `json:"errors"`
}

type DbtProjectDetailsResponse struct {
	common.CommonResponse
	Data struct {
		DbtProjectDetailsData
	} `json:"data"`
}

type DbtProjectsListResponse struct {
	common.CommonResponse
	Data struct {
		Items      []DbtProjectItem `json:"items"`
		NextCursor string           `json:"next_cursor"`
	} `json:"data"`
}

type DbtTransformationSchedule struct {
	scheduleType *string
	daysOfWeek   []string
	interval     *int
	timeOfDay    *string
}

type dbtTransformationScheduleRequest struct {
	ScheduleType *string  `json:"schedule_type,omitempty"`
	DaysOfWeek   []string `json:"days_of_week,omitempty"`
	Interval     *int     `json:"interval,omitempty"`
	TimeOfDay    *string  `json:"time_of_day,omitempty"`
}

type DbtTransformationScheduleResponse struct {
	ScheduleType string   `json:"schedule_type"`
	DaysOfWeek   []string `json:"days_of_week"`
	Interval     int      `json:"interval"`
	TimeOfDay    string   `json:"time_of_day"`
}

type dbtTransformationModifyRequest struct {
	Schedule any   `json:"schedule,omitempty"`
	RunTests *bool `json:"run_tests,omitempty"`
	Paused   *bool `json:"paused,omitempty"`
}

type dbtTransformationCreateRequest struct {
	DbtModelId *string `json:"dbt_model_id,omitempty"`
	Schedule   any     `json:"schedule,omitempty"`
	RunTests   *bool   `json:"run_tests,omitempty"`
	Paused     *bool   `json:"paused,omitempty"`
}
type DbtProjectConfig struct {
	gitRemoteUrl *string
	gitBranch    *string
	folderPath   *string
}

type dbtProjectConfigRequest struct {
	GitRemoteUrl *string `json:"git_remote_url,omitempty"`
	GitBranch    *string `json:"git_branch,omitempty"`
	FolderPath   *string `json:"folder_path,omitempty"`
}

type DbtProjectConfigResponse struct {
	GitRemoteUrl string `json:"git_remote_url"`
	GitBranch    string `json:"git_branch"`
	FolderPath   string `json:"folder_path"`
}

type dbtProjectModifyRequest struct {
	DbtVersion      *string   `json:"dbt_version,omitempty"`
	TargetName      *string   `json:"target_name,omitempty"`
	Threads         *int      `json:"threads,omitempty"`
	EnvironmentVars *[]string `json:"environment_vars,omitempty"`
	ProjectConfig   any       `json:"project_config,omitempty"`
}


type dbtProjectCreateRequest struct {
	GroupID         *string   `json:"group_id,omitempty"`
	DbtVersion      *string   `json:"dbt_version,omitempty"`
	DefaultSchema   *string   `json:"default_schema,omitempty"`
	TargetName      *string   `json:"target_name,omitempty"`
	Threads         *int      `json:"threads,omitempty"`
	EnvironmentVars *[]string `json:"environment_vars,omitempty"`
	Type            *string   `json:"type,omitempty"`
	ProjectConfig   any       `json:"project_config,omitempty"`
}

type DbtProjectTestResponse struct {
	Code string `json:"code"`
	Data struct {
		DbtProjectId string `json:"dbt_project_id"`
		SetupTests []struct {
			Title   string `json:"title"`
			Status  string `json:"status"`
			Message string `json:"message"`
		} `json:"setup_tests"`
	} `json:"data"`
}