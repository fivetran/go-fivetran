package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewDbtTransformationDetailsE2E(t *testing.T) {
	t.Skip("Skipping test until we get more info on dbt transformations data which can be used for testing")

	dbtTransformationId := testutils.CreateTempDbtTransformation(t)

	details, err := testutils.Client.NewDbtTransformationDetailsService().
		TransformationId(dbtTransformationId).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", details)
		t.Error(err)
	}

	testutils.AssertEqual(t, details.Code, "Success")
	testutils.AssertNotEmpty(t, details.Data.ID)
	testutils.AssertEqual(t, details.Data.Status, "PAUSED")
	testutils.AssertEmpty(t, details.Data.Schedule.DaysOfWeek)
	testutils.AssertEqual(t, details.Data.Schedule.Interval, 0)
	testutils.AssertEqual(t, details.Data.Schedule.ScheduleType, "INTEGRATED")
	testutils.AssertEmpty(t, details.Data.Schedule.TimeOfDay)
	testutils.AssertEmpty(t, details.Data.LastRun)
	testutils.AssertNotEmpty(t, details.Data.OutputModelName)
	testutils.AssertNotEmpty(t, details.Data.DbtProjectId)
	testutils.AssertNotEmpty(t, details.Data.DbtModelId)
	testutils.AssertNotEmpty(t, details.Data.NextRun)
	testutils.AssertNotEmpty(t, details.Data.CreatedAt)
	testutils.AssertNotEmpty(t, details.Data.ModelIds)
	testutils.AssertEmpty(t, details.Data.ConnectorIds)
}
