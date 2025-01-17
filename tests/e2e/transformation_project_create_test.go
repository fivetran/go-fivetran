package fivetran_test

import (
    "context"
    "testing"

    "github.com/fivetran/go-fivetran"
    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewTransformationProjectCreateE2E(t *testing.T) {
    groupId := testutils.CreateGroup(t)

    created, err := testutils.Client.NewTransformationProjectCreate().
        GroupId(groupId).
        ProjectType("DBT_GIT").
        RunTests(false).
        ProjectConfig(fivetran.NewTransformationProjectConfig().
                            DbtVersion("1.0.1").
                            GitRemoteUrl("git@some-host.com/project.git").
                            Threads(1)).
        Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", created)
        t.Error(err)
    }

    testutils.AssertEqual(t, created.Code, "Success")
    testutils.AssertNotEmpty(t, created.Message)
    testutils.AssertNotEmpty(t, created.Data.Id)
    testutils.AssertEqual(t, created.Data.ProjectType, "DBT_GIT")
    testutils.AssertNotEmpty(t, created.Data.CreatedAt)
    testutils.AssertEqual(t, created.Data.GroupId, groupId)
    testutils.AssertNotEmpty(t, created.Data.CreatedById)
    testutils.AssertNotEmpty(t, created.Data.Status)
    testutils.AssertEqual(t, created.Data.ProjectConfig.DbtVersion, "1.0.1")
    testutils.AssertEqual(t, created.Data.ProjectConfig.GitRemoteUrl, "git@some-host.com/project.git")
    testutils.AssertEqual(t, created.Data.ProjectConfig.Threads, 1)

    t.Cleanup(func() { 
        testutils.DeleteTransformationProject(t, created.Data.Id) 
        testutils.DeleteGroup(t, groupId)
    })
}
