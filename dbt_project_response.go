package fivetran

type DbtProjectDetailsData struct {
	ID              string                   `json:"id"`
	DbtVersion      string                   `json:"dbt_version"`
	GroupID         string                   `json:"group_id"`
	CreatedAt       string                   `json:"created_at"`
	CreatedById     string                   `json:"created_by_id"`
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
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		DbtProjectDetailsData
	} `json:"data"`
}
