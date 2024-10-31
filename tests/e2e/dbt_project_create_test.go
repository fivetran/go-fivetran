package fivetran_test

import (
	"context"
	"testing"

	"github.com/fivetran/go-fivetran"
	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewDbtProjectCreateE2E(t *testing.T) {
	t.Skip("Unstable test; will be delete with refactoring DBT")

	dbtVersion := "1.3.1"
	gitRemoteUrl := "https://github.com/fivetran/dbt_demo"
	gitBranch := "main"
	folderPath := "/folder/path"
	defaultSchema := "default_schema"
	targetName := "target_name"
	threads := 1
	variable := "DBT_VARIABLE=VALUE"
	projectType := "GIT"

	testutils.CreateDbtDestination(t)

	created, err := testutils.Client.NewDbtProjectCreate().
		GroupID(testutils.PredefinedGroupId).
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

	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertNotEmpty(t, created.Message)
	testutils.AssertNotEmpty(t, created.Data.ID)

	testutils.AssertEqual(t, created.Data.GroupId, testutils.PredefinedGroupId)
	testutils.AssertNotEmpty(t, created.Data.CreatedAt)

	testutils.AssertEqual(t, created.Data.TargetName, targetName)
	testutils.AssertEqual(t, created.Data.DefaultSchema, defaultSchema)
	testutils.AssertNotEmpty(t, created.Data.PublicKey)
	testutils.AssertEqual(t, created.Data.CreatedById, testutils.PredefinedUserId)

	testutils.AssertEqual(t, created.Data.ProjectConfig.GitRemoteUrl, gitRemoteUrl)
	testutils.AssertEqual(t, created.Data.ProjectConfig.GitBranch, gitBranch)
	testutils.AssertEqual(t, created.Data.ProjectConfig.FolderPath, folderPath)

	testutils.AssertEqual(t, len(created.Data.EnvironmentVars), 1)
	testutils.AssertEqual(t, created.Data.EnvironmentVars[0], variable)

	t.Cleanup(func() {
		testutils.CleanupDbtProjects()
		testutils.DeleteDbtDestination()
	})
}
