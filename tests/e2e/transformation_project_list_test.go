package fivetran_test

import (
    "context"
    "testing"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewTransformationProjectsListE2E(t *testing.T) {
    t.Skip("Destination must be CONNECTED")
    groupId := testutils.CreateGroup(t)
    destinationId := testutils.CreateDestination(t)
    projectId := testutils.CreateTransformationProject(t)

    result, err := testutils.Client.NewTransformationProjectsList().Do(context.Background())
    if err != nil {
        t.Logf("%+v\n", result)
        t.Error(err)
    }

    testutils.AssertEqual(t, result.Code, "Success")
    testutils.AssertEqual(t, result.Data.Items[0].Id, projectId)
    testutils.AssertNotEmpty(t, result.Data.Items[0].Id)
    testutils.AssertEqual(t, result.Data.Items[0].ProjectType, "DBT_GIT")
    testutils.AssertNotEmpty(t, result.Data.Items[0].CreatedAt)
    testutils.AssertEqual(t, result.Data.Items[0].GroupId, groupId)
    testutils.AssertNotEmpty(t, result.Data.Items[0].CreatedById)

    t.Cleanup(func() { 
        testutils.DeleteTransformationProject(t, projectId)
        testutils.DeleteDestination(t, destinationId)
        testutils.DeleteGroup(t, groupId)
    })
}
