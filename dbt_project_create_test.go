package fivetran_test

import (
	"context"
	"testing"
)

func TestNewDbtProjectCreateE2E(t *testing.T) {
	t.Skip("Skipping test until we get more info on dbt projects data which can be used for testing")

	created, err := Client.NewDbtProjectCreate().
		GroupID(PredefinedGroupId).
		DbtVersion("1.3.1").
		GitRemoteUrl("https://github.com/fivetran/dbt_demo").
		GitBranch("main").
		DefaultSchema("").
		FolderPath("").
		TargetName("").
		Threads(4).
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
	AssertNotEmpty(t, created.Data.CreatedAt)
	AssertEqual(t, created.Data.TargetName, "")
	AssertEqual(t, created.Data.GitRemoteUrl, "https://github.com/fivetran/dbt_demo")
	AssertEqual(t, created.Data.DefaultSchema, "")
	AssertEqual(t, created.Data.PublicKey, "")
	AssertEqual(t, created.Data.CreatedById, "")
	AssertEqual(t, created.Data.GitBranch, "main")
}
