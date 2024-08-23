package dbt

func (dc *DbtProjectConfig) Request() *dbtProjectConfigRequest {
	return &dbtProjectConfigRequest{
		GitRemoteUrl: dc.gitRemoteUrl,
		GitBranch:    dc.gitBranch,
		FolderPath:   dc.folderPath,
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
