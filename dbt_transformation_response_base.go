package fivetran

import "time"

type DbtTransformationResponseBase struct {
	ID              string    `json:"id"`
	DbtModelID      string    `json:"dbt_model_id"`
	OutputModelName string    `json:"output_model_name"`
	DbtProjectID    string    `json:"dbt_project_id"`
	LastRun         time.Time `json:"last_run"`
	NextRun         time.Time `json:"next_run"`
	Status          string    `json:"status"`
	RunTests        bool      `json:"run_tests"`
	ConnectorIDs    []string  `json:"connector_ids"`
	ModelIDs        []string  `json:"model_ids"`
}
