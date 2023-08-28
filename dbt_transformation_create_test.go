package fivetran_test

import (
	"context"
	"testing"

	"github.com/fivetran/go-fivetran"
)

const (
	dbtModelId   = "real_dbt_model_id"
	dbtProjectId = "real_dbt_project_id"
)

func TestNewDbtTransformationCreateE2E(t *testing.T) {
	t.Skip("Skipping test until we get more info on dbt transformations data which can be used for testing")

	created, err := Client.NewDbtTransformationCreateService().
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

	AssertEqual(t, created.Code, "Success")
	AssertNotEmpty(t, created.Message)
	AssertNotEmpty(t, created.Data.ID)

	// check managable fields
	AssertEmpty(t, created.Data.Schedule.DaysOfWeek)
	AssertEqual(t, created.Data.DbtModelId, dbtModelId)
	AssertEqual(t, created.Data.Schedule.Interval, 0)
	AssertEqual(t, created.Data.Schedule.ScheduleType, "INTEGRATED")
	AssertEqual(t, created.Data.Schedule.TimeOfDay, "")
	AssertEqual(t, created.Data.RunTests, true)
	AssertEqual(t, created.Data.Paused, true)

	// check readonly fields
	AssertEqual(t, created.Data.Status, "PAUSED")
	AssertNotEmpty(t, created.Data.CreatedAt)
	AssertEmpty(t, created.Data.LastRun)
	AssertNotEmpty(t, created.Data.OutputModelName)
	AssertEqual(t, created.Data.DbtProjectId, dbtProjectId)
	AssertNotEmpty(t, created.Data.NextRun)
	AssertNotEmpty(t, created.Data.ModelIds)
	AssertEmpty(t, created.Data.ConnectorIds)

	t.Cleanup(func() { DeleteDbtTransformation(t, created.Data.ID) })
}
