package fivetran_test

import (
	"context"
	"testing"

	"github.com/fivetran/go-fivetran"
	testutils "github.com/fivetran/go-fivetran/test_utils"
)

const (
	dbtModelId   = "real_dbt_model_id"
	dbtProjectId = "real_dbt_project_id"
)

func TestNewDbtTransformationCreateE2E(t *testing.T) {
	t.Skip("Skipping test until we get more info on dbt transformations data which can be used for testing")

	created, err := testutils.Client.NewDbtTransformationCreateService().
		DbtModelId(dbtModelId).
		Schedule(fivetran.NewDbtTransformationSchedule().
			ScheduleType("INTEGRATED").
			DaysOfWeek([]string{}).
			Interval(0).
			TimeOfDay("")).
		RunTests(true).
		Paused(true).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	testutils.AssertEqual(t, created.Code, "Success")
	testutils.AssertNotEmpty(t, created.Message)
	testutils.AssertNotEmpty(t, created.Data.ID)

	// check managable fields
	testutils.AssertEmpty(t, created.Data.Schedule.DaysOfWeek)
	testutils.AssertEqual(t, created.Data.DbtModelId, dbtModelId)
	testutils.AssertEqual(t, created.Data.Schedule.Interval, 0)
	testutils.AssertEqual(t, created.Data.Schedule.ScheduleType, "INTEGRATED")
	testutils.AssertEqual(t, created.Data.Schedule.TimeOfDay, "")
	testutils.AssertEqual(t, created.Data.RunTests, true)
	testutils.AssertEqual(t, created.Data.Paused, true)

	// check readonly fields
	testutils.AssertEqual(t, created.Data.Status, "PAUSED")
	testutils.AssertNotEmpty(t, created.Data.CreatedAt)
	testutils.AssertEmpty(t, created.Data.LastRun)
	testutils.AssertNotEmpty(t, created.Data.OutputModelName)
	testutils.AssertEqual(t, created.Data.DbtProjectId, dbtProjectId)
	testutils.AssertNotEmpty(t, created.Data.NextRun)
	testutils.AssertNotEmpty(t, created.Data.ModelIds)
	testutils.AssertEmpty(t, created.Data.ConnectorIds)

	t.Cleanup(func() { testutils.DeleteDbtTransformation(t, created.Data.ID) })
}
