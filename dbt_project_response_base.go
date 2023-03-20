package fivetran

import "time"

type DbtProjectResponseBase struct {
	ID            string    `json:"id"`
	GroupID       string    `json:"group_id"`
	CreatedAt     time.Time `json:"created_at"`
	CreatedById   string    `json:"created_by_id"`
	PublicKey     string    `json:"public_key"`
	GitRemoteUrl  string    `json:"git_remote_url"`
	GitBranch     string    `json:"git_branch"`
	DefaultSchema string    `json:"default_schema"`
	FolderPath    string    `json:"folder_path"`
	TargetName    string    `json:"target_name"`
}
