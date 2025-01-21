package transformations

import "github.com/fivetran/go-fivetran/utils"

func (elc *TransformationSchedule) Request() *transformationScheduleRequest {
	return &transformationScheduleRequest{
		Cron:  				elc.cron,
		ConnectionIds:  	elc.connectionIds,
		DaysOfWeek: 		elc.daysOfWeek,
		TimeOfDay:			elc.timeOfDay,
		ScheduleType:		elc.scheduleType,
		Interval:			elc.interval,
		SmartSyncing:		elc.smartSyncing,
	}
}

func (elc *TransformationSchedule) Merge(customConfig *map[string]interface{}) (*map[string]interface{}, error) {
	err := utils.MergeIntoMap(elc.Request(), customConfig)
	if err != nil {
		return nil, err
	}
	return customConfig, nil
}

func (elc *TransformationSchedule) Cron(value []string) *TransformationSchedule {
	elc.cron = &value
	return elc
}

func (elc *TransformationSchedule) ConnectionIds(value []string) *TransformationSchedule {
	elc.connectionIds = &value
	return elc
}

func (elc *TransformationSchedule) DaysOfWeek(value []string) *TransformationSchedule {
	elc.daysOfWeek = &value
	return elc
}

func (elc *TransformationSchedule) TimeOfDay(value string) *TransformationSchedule {
	elc.timeOfDay = &value
	return elc
}

func (elc *TransformationSchedule) ScheduleType(value string) *TransformationSchedule {
	elc.scheduleType = &value
	return elc
}

func (elc *TransformationSchedule) SmartSyncing(value bool) *TransformationSchedule {
	elc.smartSyncing = &value
	return elc
}

func (elc *TransformationSchedule) Interval(value int) *TransformationSchedule {
	elc.interval = &value
	return elc
}