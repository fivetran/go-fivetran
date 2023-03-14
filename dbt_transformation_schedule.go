package fivetran

import "time"

// DbtTransformationSchedule builds dbt transformation management, dbt transformation schedule
type DbtTransformationSchedule struct {
	scheduleType *string
	daysOfWeek   []string
	timeOfDay    *time.Time
}

type dbtTransformationScheduleRequest struct {
	ScheduleType *string
	DaysOfWeek   []string
	TimeOfDay    *time.Time
}

type DbtTransformationScheduleResponse struct {
	ScheduleType *string
	DaysOfWeek   []string
	Interval     *string
	TimeOfDay    *time.Time
}

func NewDbtTransformationSchedule() *DbtTransformationSchedule {
	return &DbtTransformationSchedule{}
}

func (cc *DbtTransformationSchedule) request() *dbtTransformationScheduleRequest {

	return &dbtTransformationScheduleRequest{
		ScheduleType: cc.scheduleType,
		DaysOfWeek:   cc.daysOfWeek,
		TimeOfDay:    cc.timeOfDay,
	}
}

func (cc *DbtTransformationSchedule) ScheduleType(value string) *DbtTransformationSchedule {
	cc.scheduleType = &value
	return cc
}

func (cc *DbtTransformationSchedule) DaysOfWeek(value []string) *DbtTransformationSchedule {
	cc.daysOfWeek = value
	return cc
}

func (cc *DbtTransformationSchedule) TimeOfDay(value time.Time) *DbtTransformationSchedule {
	cc.timeOfDay = &value
	return cc
}
