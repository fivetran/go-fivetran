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
