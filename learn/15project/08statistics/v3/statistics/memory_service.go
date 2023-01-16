package statistics

import (
	"fmt"
	"runtime"
	serror "v3/statistics/error"
)

type memoryService struct {
	memoryInfo map[string][]uint64
}

func newMemoryService() *memoryService {
	return &memoryService{
		memoryInfo: make(map[string][]uint64),
	}
}

func (ms *memoryService) add(taskName string) error {
	if _, ok := ms.memoryInfo[taskName]; !ok {
		ms.memoryInfo[taskName] = []uint64{nowMemory()}
	}
	return nil
}

func (ms *memoryService) end(taskName string) error {
	task, ok := ms.memoryInfo[taskName]
	if !ok {
		return serror.TaskNotExistsErr(taskName)
	}
	if len(task) == 2 {
		return nil
	}
	ms.memoryInfo[taskName] = append(task, nowMemory())
	return nil
}

func (ms *memoryService) print(taskName string) (string, error) {
	task, ok := ms.memoryInfo[taskName]
	if !ok {
		return "", serror.TaskNotExistsErr(taskName)
	}
	if len(task) != 2 {
		return "", serror.TaskLengthError(taskName)
	}
	startMemory := task[0]
	endMemory := task[1]
	return fmt.Sprintf("任务：【%s】，一共消耗内存：【%vMiB】", taskName, bToMb(endMemory-startMemory)), nil
}

func (ms *memoryService) getAllTask() map[string][]uint64 {
	return ms.memoryInfo
}

func nowMemory() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc
}

//	func PrintMemUsage() {
//		var m runtime.MemStats
//		runtime.ReadMemStats(&m)
//		// For info on each, see: https://golang.org/pkg/runtime/#MemStats
//		fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
//		fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
//		fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
//		fmt.Printf("\tNumGC = %v\n", m.NumGC)
//	}
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
