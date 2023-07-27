package fivetran_test

import (
	"context"
	"testing"

	"github.com/fivetran/go-fivetran"
)

func TestNewDbtTransformationCreateE2E(t *testing.T) {
	created, err := Client.NewDbtTransformationCreateService().
		DbtModelId("").
		Schedule(*fivetran.NewDbtTransformationSchedule().
			ScheduleType("INTEGRATED").
			DaysOfWeek([]string{}).
			Interval(0).
			TimeOfDay("")).
		RunTests(true).
		ProjectId("").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", created)
		t.Error(err)
	}

	AssertEqual(t, created.Code, "Success")
	AssertNotEmpty(t, created.Message)
	AssertNotEmpty(t, created.Data.ID)
	AssertEqual(t, created.Data.Status, "")
	AssertEqual(t, created.Data.Schedule.DaysOfWeek, "")
	AssertEqual(t, created.Data.Schedule.Interval, "")
	AssertEqual(t, created.Data.Schedule.ScheduleType, "")
	AssertEqual(t, created.Data.Schedule.TimeOfDay, "")
	AssertEqual(t, created.Data.LastRun, "")
	AssertEqual(t, created.Data.OutputModelName, "")
	AssertEqual(t, created.Data.DbtProjectId, "")
	AssertEqual(t, created.Data.DbtModelId, "")
	AssertEqual(t, created.Data.NextRun, "")
	AssertEqual(t, created.Data.CreatedAt, "")
	AssertEqual(t, created.Data.RunTests, "")
	AssertEqual(t, created.Data.ModelIds, "")
	AssertEqual(t, created.Data.ConnectorIds, "")

	t.Cleanup(func() { DeleteDbtTransformation(t, created.Data.ID) })
}
