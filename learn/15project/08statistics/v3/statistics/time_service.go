package statistics

import (
	"fmt"
	"time"
	serror "v3/statistics/error"
)

type timeService struct {
	timeInfo map[string][]time.Time
}

func newTimeService() *timeService {
	return &timeService{
		timeInfo: make(map[string][]time.Time),
	}
}

func (ts *timeService) add(taskName string) error {
	if _, ok := ts.timeInfo[taskName]; !ok {
		ts.timeInfo[taskName] = []time.Time{nowTime()}
	}
	return nil
}

func (ts *timeService) end(taskName string) error {
	task, ok := ts.timeInfo[taskName]
	if !ok {
		return serror.TaskNotExistsErr(taskName)
	}
	if len(task) == 2 {
		return nil
	}
	ts.timeInfo[taskName] = append(task, nowTime())
	return nil
}

func (ts *timeService) print(taskName string) (string, error) {
	task, ok := ts.timeInfo[taskName]
	if !ok {
		return "", serror.TaskNotExistsErr(taskName)
	}
	if len(task) != 2 {
		return "", serror.TaskLengthError(taskName)
	}
	startTime := task[0]
	endTime := task[1]
	return fmt.Sprintf("任务：【%s】，一共耗时：【%v】", taskName, endTime.Sub(startTime).String()), nil
}

func (ts *timeService) getAllTask() map[string][]time.Time {
	return ts.timeInfo
}

func nowTime() time.Time {
	return time.Now()
}
