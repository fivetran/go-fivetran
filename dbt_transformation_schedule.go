package fivetran

type DbtTransformationSchedule struct {
	scheduleType *string
	daysOfWeek   []string
	interval     int
	timeOfDay    *string
}

type dbtTransformationScheduleRequest struct {
	ScheduleType *string  `json:"schedule_type,omitempty"`
	DaysOfWeek   []string `json:"days_of_week,omitempty"`
	Interval     int      `json:"interval,omitempty"`
	TimeOfDay    *string  `json:"time_of_day,omitempty"`
}

type dbtTransformationScheduleResponse struct {
	ScheduleType *string  `json:"schedule_type"`
	DaysOfWeek   []string `json:"days_of_week"`
	Interval     int      `json:"interval"`
	TimeOfDay    *string  `json:"time_of_day"`
}

func NewDbtTransformationSchedule() *DbtTransformationSchedule {
	return &DbtTransformationSchedule{}
}

func (dbtTransformationSchedule *DbtTransformationSchedule) request() *dbtTransformationScheduleRequest {
	return &dbtTransformationScheduleRequest{
		ScheduleType: dbtTransformationSchedule.scheduleType,
		DaysOfWeek:   dbtTransformationSchedule.daysOfWeek,
		Interval:     dbtTransformationSchedule.interval,
		TimeOfDay:    dbtTransformationSchedule.timeOfDay,
	}
}

func (dbtTransformationSchedule *DbtTransformationSchedule) ScheduleType(value string) *DbtTransformationSchedule {
	dbtTransformationSchedule.scheduleType = &value
	return dbtTransformationSchedule
}

func (dbtTransformationSchedule *DbtTransformationSchedule) DaysOfWeek(value []string) *DbtTransformationSchedule {
	dbtTransformationSchedule.daysOfWeek = value
	return dbtTransformationSchedule
}

func (dbtTransformationSchedule *DbtTransformationSchedule) Interval(value int) *DbtTransformationSchedule {
	dbtTransformationSchedule.interval = value
	return dbtTransformationSchedule
}

func (dbtTransformationSchedule *DbtTransformationSchedule) TimeOfDay(value string) *DbtTransformationSchedule {
	dbtTransformationSchedule.timeOfDay = &value
	return dbtTransformationSchedule
}
