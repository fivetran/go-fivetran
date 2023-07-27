package fivetran_test

import (
	"context"
	"testing"
)

func TestNewDbtTransformationDetailsE2E(t *testing.T) {
	t.Skip("Skipping test until we get more info on dbt transformations data which can be used for testing")

	dbtTransformationId := CreateTempDbtTransformation(t)

	details, err := Client.NewDbtTransformationDetailsService().
		TransformationId(dbtTransformationId).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", details)
		t.Error(err)
	}

	AssertEqual(t, details.Code, "Success")
	AssertNotEmpty(t, details.Data.ID)
	AssertEqual(t, details.Data.Status, "")
	AssertEqual(t, details.Data.Schedule.DaysOfWeek, "")
	AssertEqual(t, details.Data.Schedule.Interval, "")
	AssertEqual(t, details.Data.Schedule.ScheduleType, "")
	AssertEqual(t, details.Data.Schedule.TimeOfDay, "")
	AssertEqual(t, details.Data.LastRun, "")
	AssertEqual(t, details.Data.OutputModelName, "")
	AssertEqual(t, details.Data.DbtProjectId, "")
	AssertEqual(t, details.Data.DbtModelId, "")
	AssertEqual(t, details.Data.NextRun, "")
	AssertEqual(t, details.Data.CreatedAt, "")
	AssertEqual(t, details.Data.RunTests, "")
	AssertEqual(t, details.Data.ModelIds, "")
	AssertEqual(t, details.Data.ConnectorIds, "")
}
