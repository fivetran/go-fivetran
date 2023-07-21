package fivetran_test

import (
	"context"
	"testing"
)

func TestNewDbtProjectCreateE2E(t *testing.T) {
	created, err := Client.NewDbtProjectCreate().
		GroupID(PredefinedGroupId).
		DbtVersion("").
		GitRemoteUrl("").
		GitBranch("").
		DefaultSchema("").
		FolderPath("").
		TargetName("").
		Threads(0).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	AssertEqual(t, created.Code, "Success")
	AssertNotEmpty(t, created.Message)
	AssertEqual(t, created.Data.ID, PredefinedGroupId)
	AssertEqual(t, created.Data.GroupID, PredefinedGroupId)
	AssertEqual(t, created.Data.FolderPath, "")
	//AssertEqual(t, created.Data.CreatedAt, "")
	AssertEqual(t, created.Data.TargetName, "")
	AssertEqual(t, created.Data.GitRemoteUrl, "")
	AssertEqual(t, created.Data.DefaultSchema, "")
	AssertEqual(t, created.Data.PublicKey, "")
	AssertEqual(t, created.Data.CreatedById, "")
	AssertEqual(t, created.Data.GitBranch, "")
}
