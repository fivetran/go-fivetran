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
	AssertEqual(t, details.Data.Status, "PAUSED")
	AssertEmpty(t, details.Data.Schedule.DaysOfWeek)
	AssertEqual(t, details.Data.Schedule.Interval, 0)
	AssertEqual(t, details.Data.Schedule.ScheduleType, "INTEGRATED")
	AssertEmpty(t, details.Data.Schedule.TimeOfDay)
	AssertEmpty(t, details.Data.LastRun)
	AssertNotEmpty(t, details.Data.OutputModelName)
	AssertNotEmpty(t, details.Data.DbtProjectId)
	AssertNotEmpty(t, details.Data.DbtModelId)
	AssertNotEmpty(t, details.Data.NextRun)
	AssertNotEmpty(t, details.Data.CreatedAt)
	AssertNotEmpty(t, details.Data.ModelIds)
	AssertEmpty(t, details.Data.ConnectorIds)
}
