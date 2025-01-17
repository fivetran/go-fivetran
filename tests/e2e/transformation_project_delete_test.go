package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewTransformationProjectDeleteE2E(t *testing.T) {
    groupId := testutils.CreateGroup(t)
    projectId := testutils.CreateTransformationProject(t)
	
	deleted, err := testutils.Client.NewTransformationProjectDelete().ProjectId(projectId).Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", deleted)
		t.Error(err)
	}

	testutils.AssertEqual(t, deleted.Code, "Success")
	testutils.AssertNotEmpty(t, deleted.Message)

	t.Cleanup(func() { 
        testutils.DeleteGroup(t, groupId)
    })
}
