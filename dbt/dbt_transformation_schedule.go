package dbt

func (dbtTransformationSchedule *DbtTransformationSchedule) Request() *dbtTransformationScheduleRequest {
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
	dbtTransformationSchedule.interval = &value
	return dbtTransformationSchedule
}

func (dbtTransformationSchedule *DbtTransformationSchedule) TimeOfDay(value string) *DbtTransformationSchedule {
	dbtTransformationSchedule.timeOfDay = &value
	return dbtTransformationSchedule
}
