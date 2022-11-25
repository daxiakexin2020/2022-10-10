package manage_center

import (
	"errors"
)

type manager struct {
	Boots  []func() error
	KS     KernelServer
	IsBoot bool
}

type KernelServer interface {
	Name() string
	Run() error
}

func NewManager() *manager {
	return &manager{
		Boots: make([]func() error, 0),
	}
}

func (m *manager) Register(boots ...func() error) {
	m.Boots = append(m.Boots, boots...)
}

func (m *manager) RegisterKernel(ks KernelServer) {
	m.KS = ks
}

func (m *manager) Run() error {
	if m.IsBoot {
		return errors.New(">>>>>>>>>>>>>>>>>>>>>>框架已经启动，不要重复启动<<<<<<<<<<<<<<<<<<<<<<<<<")
	}
	m.IsBoot = true
	for _, boot := range m.Boots {
		if err := boot(); err != nil {
			return err
		}
	}
	return m.run()
}

func (m *manager) run() error {
	if m.KS == nil {
		return errors.New(">>>>>>>>>>>>>>>>>>>>>>没有注册服务，请先注册服务<<<<<<<<<<<<<<<<<<<<<<<<<")
	}
	return m.KS.Run()
}
