package fivetran_test

import (
    "context"
    "testing"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewTransformationProjectDetailsE2E(t *testing.T) {
    t.Skip("Destination must be CONNECTED")
    groupId := testutils.CreateGroup(t)
    destinationId := testutils.CreateDestination(t)
    projectId := testutils.CreateTransformationProject(t)

    result, err := testutils.Client.NewTransformationProjectDetails().
        ProjectId(projectId).
        Do(context.Background())
    if err != nil {
        t.Logf("%+v\n", result)
        t.Error(err)
    }

    testutils.AssertEqual(t, result.Code, "Success")
    testutils.AssertEqual(t, result.Data.Id, projectId)
    testutils.AssertNotEmpty(t, result.Data.Id)
    testutils.AssertEqual(t, result.Data.ProjectType, "DBT_GIT")
    testutils.AssertNotEmpty(t, result.Data.CreatedAt)
    testutils.AssertEqual(t, result.Data.GroupId, groupId)
    testutils.AssertNotEmpty(t, result.Data.CreatedById)
    testutils.AssertNotEmpty(t, result.Data.Status)
    testutils.AssertEqual(t, result.Data.ProjectConfig.DbtVersion, "1.0.1")
    testutils.AssertEqual(t, result.Data.ProjectConfig.GitRemoteUrl, "git@some-host.com/project.git")
    testutils.AssertEqual(t, result.Data.ProjectConfig.Threads, 1)

    t.Cleanup(func() { 
        testutils.DeleteTransformationProject(t, projectId)
        testutils.DeleteDestination(t, destinationId)
        testutils.DeleteGroup(t, groupId)
    })
}
