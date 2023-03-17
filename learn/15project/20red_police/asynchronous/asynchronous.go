package asynchronous

import (
	"errors"
	"log"
	"sync"
)

type Tasker interface {
	TaskName() string
	Run() error
	Stop() error
	ExitSignal() chan struct{}
}

type manager struct {
	list map[string]Tasker
	mu   sync.Mutex
	exit chan struct{}
	boot bool
}

var (
	gmanager *manager
	monce    sync.Once
)

func newmanager() *manager {
	return &manager{list: map[string]Tasker{}}
}

func Manager() *manager {
	monce.Do(func() {
		gmanager = newmanager()
	})
	return gmanager
}

func (m *manager) Register(tasks ...Tasker) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.boot {
		return errors.New("asynchronous already is runing,can not register task")
	}
	for _, task := range tasks {
		m.list[task.TaskName()] = task
	}
	return nil
}

func (m *manager) StopTask(taskName string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if !m.boot {
		return errors.New("asynchronous is stoped")
	}
	if task, ok := m.list[taskName]; ok {
		delete(m.list, taskName)
		return task.Stop()
	}
	return errors.New("this task is not registered:" + taskName)
}

func (m *manager) Run() {
	if m.boot {
		return
	}
	m.exit = make(chan struct{}, len(m.list))
	for _, task := range m.list {
		go task.Run()
	}
	m.boot = true
	for {
		if len(m.list) == 0 {
			log.Println("all task is stoped,asynchronous stoping..........")
			m.boot = false
			return
		}
	}
}
