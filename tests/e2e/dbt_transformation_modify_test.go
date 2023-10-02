package fivetran_test

import (
	"context"
	"testing"

	"github.com/fivetran/go-fivetran"
	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewDbtTransformationModifyE2E(t *testing.T) {
	t.Skip("Skipping test until we get more info on dbt transformations data which can be used for testing")

	dbtTransformationId := testutils.CreateTempDbtTransformation(t)
	details, err := testutils.Client.NewDbtTransformationModifyService().
		DbtTransformationId(dbtTransformationId).
		Schedule(fivetran.NewDbtTransformationSchedule().
			ScheduleType("INTEGRATED").
			DaysOfWeek([]string{}).
			Interval(0).
			TimeOfDay("")).
		RunTests(true).
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", details)
		t.Error(err)
	}

	// These checks are inconsistent currently, will update it when enable the test
	testutils.AssertEqual(t, details.Code, "Success")
	testutils.AssertNotEmpty(t, details.Message)
	testutils.AssertNotEmpty(t, details.Data.ID)
	testutils.AssertEqual(t, details.Data.Status, "")
	testutils.AssertEmpty(t, details.Data.Schedule.DaysOfWeek)
	testutils.AssertEqual(t, details.Data.Schedule.Interval, "")
	testutils.AssertEqual(t, details.Data.Schedule.ScheduleType, "")
	testutils.AssertEqual(t, details.Data.Schedule.TimeOfDay, "")
	testutils.AssertEqual(t, details.Data.LastRun, "")
	testutils.AssertEqual(t, details.Data.OutputModelName, "")
	testutils.AssertEqual(t, details.Data.DbtProjectId, "")
	testutils.AssertEqual(t, details.Data.DbtModelId, "")
	testutils.AssertEqual(t, details.Data.NextRun, "")
	testutils.AssertEqual(t, details.Data.CreatedAt, "")
	testutils.AssertNotEmpty(t, details.Data.ModelIds)
	testutils.AssertEmpty(t, details.Data.ConnectorIds)
}
