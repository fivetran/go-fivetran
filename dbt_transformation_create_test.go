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

	t.Cleanup(func() { DeleteDbtTransformation(t, created.Data.ID) })
}
