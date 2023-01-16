package statistics

import (
	"fmt"
	"runtime"
	"time"
	serror "v2/statistics/error"
)

type service struct {
	bucket map[string]*bdata
}

type bdata struct {
	timeInfo   []time.Time
	memoryInfo []uint64
}

func newBdata() *bdata {
	return &bdata{
		timeInfo:   make([]time.Time, 0),
		memoryInfo: make([]uint64, 0),
	}
}

func newService() *service {
	return &service{bucket: make(map[string]*bdata)}
}

func (s *service) add(taskName string) error {
	if _, ok := s.bucket[taskName]; !ok {
		bd := newBdata()
		bd.timeInfo = []time.Time{nowTime()}
		bd.memoryInfo = []uint64{nowMemory()}
		s.bucket[taskName] = bd
	}
	return nil
}

func (s *service) end(taskName string) error {
	task, ok := s.bucket[taskName]
	if !ok {
		return serror.TaskNotExistsErr(taskName)
	}
	if len(task.timeInfo) == 2 {
		return nil
	}
	s.bucket[taskName].timeInfo = append(s.bucket[taskName].timeInfo, nowTime())
	s.bucket[taskName].memoryInfo = append(s.bucket[taskName].memoryInfo, nowMemory())
	return nil
}

func (s *service) print(taskName string) (string, error) {
	task, ok := s.bucket[taskName]
	if !ok {
		return "", serror.TaskNotExistsErr(taskName)
	}
	if len(task.timeInfo) != 2 || len(task.memoryInfo) != 2 {
		return "", serror.TaskLengthError(taskName)
	}

	startTime := task.timeInfo[0]
	endTime := task.timeInfo[1]
	startMemory := task.memoryInfo[0]
	endMemory := task.memoryInfo[1]
	timeInfo := fmt.Sprintf("任务：【%s】，一共耗时：【%v】", taskName, endTime.Sub(startTime).String())
	memoryInfo := fmt.Sprintf("任务：【%s】，一共消耗内存：【%vMiB】", taskName, bToMb(endMemory-startMemory))

	return timeInfo + "=========" + memoryInfo, nil
}

func nowTime() time.Time {
	return time.Now()
}

func nowMemory() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
