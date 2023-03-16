package asynchronous

import (
	"errors"
	"fmt"
	"sync"
)

type Tasker interface {
	TaskName() string
	Run() error
	Stop() error
	ExitSignal() chan struct{}
}

type Manager struct {
	list map[string]Tasker
	mu   sync.Mutex
}

func NewManager() *Manager {
	return &Manager{list: map[string]Tasker{}}
}

func (m *Manager) Register(tasks ...Tasker) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, task := range tasks {
		m.list[task.TaskName()] = task
	}
}

func (m *Manager) Stop(taskName string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if task, ok := m.list[taskName]; ok {
		return task.Stop()
	}
	return errors.New("this task is not registered:" + taskName)
}

func (m *Manager) Run() {

	mexit := make(chan struct{}, len(m.list))
	for _, task := range m.list {
		go task.Run()
	}
	for _, task := range m.list {
		go func() {
			for {
				select {
				case <-task.ExitSignal():
					fmt.Println("task stop:" + task.TaskName())
					task.Stop()
					mexit <- struct{}{}
				default:
				}
			}
		}()
	}
	for {
		if len(mexit) == len(m.list) {
			fmt.Println("all task is stop")
			return
		}
		fmt.Println("（（（（（（（（（（（（（")
	}
}
