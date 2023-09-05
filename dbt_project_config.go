package fivetran

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

type updateDbtProjectConfigRequest struct {
	GitBranch  *string `json:"git_branch,omitempty"`
	FolderPath *string `json:"folder_path,omitempty"`
}

type DbtProjectConfigResponse struct {
	GitRemoteUrl string `json:"git_remote_url"`
	GitBranch    string `json:"git_branch"`
	FolderPath   string `json:"folder_path"`
}

func NewDbtProjectConfig() *DbtProjectConfig {
	return &DbtProjectConfig{}
}

func (dc *DbtProjectConfig) request() *dbtProjectConfigRequest {
	return &dbtProjectConfigRequest{
		GitRemoteUrl: dc.gitRemoteUrl,
		GitBranch:    dc.gitBranch,
		FolderPath:   dc.folderPath,
	}
}

func (dc *DbtProjectConfig) updateRequest() *updateDbtProjectConfigRequest {
	return &updateDbtProjectConfigRequest{
		GitBranch:  dc.gitBranch,
		FolderPath: dc.folderPath,
	}
}

func (dc *DbtProjectConfig) GitRemoteUrl(value string) *DbtProjectConfig {
	dc.gitRemoteUrl = &value
	return dc
}

func (dc *DbtProjectConfig) GitBranch(value string) *DbtProjectConfig {
	dc.gitBranch = &value
	return dc
}

func (dc *DbtProjectConfig) FolderPath(value string) *DbtProjectConfig {
	dc.folderPath = &value
	return dc
}
