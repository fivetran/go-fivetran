package fivetran_test

import (
	"context"
	"testing"

	"github.com/fivetran/go-fivetran"
)

func TestNewDbtProjectCreateE2E(t *testing.T) {
	dbtVersion := "1.3.1"
	gitRemoteUrl := "https://github.com/fivetran/dbt_demo"
	gitBranch := "main"
	folderPath := "/folder/path"
	defaultSchema := "default_schema"
	targetName := "target_name"
	threads := 1
	variable := "ENV_VAR1=VALUE"
	projectType := "GIT"

	created, err := Client.NewDbtProjectCreate().
		GroupID(PredefinedGroupId).
		DbtVersion(dbtVersion).
		ProjectConfig(fivetran.NewDbtProjectConfig().
			GitRemoteUrl(gitRemoteUrl).
			GitBranch(gitBranch).
			FolderPath(folderPath)).
		DefaultSchema(defaultSchema).
		TargetName(targetName).
		Threads(threads).
		EnvironmentVars([]string{variable}).
		Type(projectType).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	AssertEqual(t, created.Code, "Success")
	AssertNotEmpty(t, created.Message)
	AssertNotEmpty(t, created.Data.ID)

	AssertEqual(t, created.Data.GroupID, PredefinedGroupId)
	AssertNotEmpty(t, created.Data.CreatedAt)

	AssertEqual(t, created.Data.TargetName, targetName)
	AssertEqual(t, created.Data.DefaultSchema, defaultSchema)
	AssertNotEmpty(t, created.Data.PublicKey)
	AssertEqual(t, created.Data.CreatedById, PredefinedUserId)

	AssertEqual(t, created.Data.ProjectConfig.GitRemoteUrl, gitRemoteUrl)
	AssertEqual(t, created.Data.ProjectConfig.GitBranch, gitBranch)
	AssertEqual(t, created.Data.ProjectConfig.FolderPath, folderPath)

	AssertEqual(t, len(created.Data.EnvironmentVars), 1)
	AssertEqual(t, created.Data.EnvironmentVars[0], variable)

	t.Cleanup(func() { cleanupDbtProjects() })
}
