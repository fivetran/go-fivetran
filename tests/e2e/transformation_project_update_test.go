package fivetran_test

import (
    "context"
    "testing"

    "github.com/fivetran/go-fivetran"
    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewTransformationProjectUpdateE2E(t *testing.T) {
    groupId := testutils.CreateGroup(t)
    projectId := testutils.CreateTransformationProject(t)

    updated, err := testutils.Client.NewTransformationProjectUpdate().
        ProjectId(projectId).
        ProjectConfig(fivetran.NewTransformationProjectConfig().
                            DbtVersion("1.0.0").
                            GitRemoteUrl("git@some-host.com/project.git").
                            Threads(1)).
        Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", updated)
        t.Error(err)
    }

    testutils.AssertEqual(t, updated.Code, "Success")
    testutils.AssertNotEmpty(t, updated.Message)
    testutils.AssertNotEmpty(t, updated.Data.Id)
    testutils.AssertEqual(t, updated.Data.ProjectType, "DBT_GIT")
    testutils.AssertNotEmpty(t, updated.Data.CreatedAt)
    testutils.AssertEqual(t, updated.Data.GroupId, groupId)
    testutils.AssertNotEmpty(t, updated.Data.CreatedById)
    testutils.AssertNotEmpty(t, updated.Data.Status)
    testutils.AssertEqual(t, updated.Data.ProjectConfig.DbtVersion, "1.0.0")
    testutils.AssertEqual(t, updated.Data.ProjectConfig.GitRemoteUrl, "git@some-host.com/project.git")
    testutils.AssertEqual(t, updated.Data.ProjectConfig.Threads, 1)

    t.Cleanup(func() { 
        testutils.DeleteTransformationProject(t, projectId) 
        testutils.DeleteGroup(t, groupId)
    })
}
