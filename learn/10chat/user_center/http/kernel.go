package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"sync"
)

type ginKernel struct {
	Logger *logrus.Logger
	Router *gin.Engine
}

var (
	R  *gin.Engine
	gk *ginKernel
)

func (gk *ginKernel) Name() string {
	return "gin_kernel"
}

func (gk *ginKernel) Run() error {
	return gk.Router.Run(":8888")
}

var once sync.Once

func Initkernel() *ginKernel {
	once.Do(func() {
		if gk == nil {
			gk = &ginKernel{
				Logger: new(logrus.Logger),
				Router: R,
			}
		}
	})
	return gk
}
