package server

import (
	"chip_database/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Info struct {
	infoSrv *service.InfoService
}

func NewInfo(infoSrv *service.InfoService) *Info {
	return &Info{infoSrv: infoSrv}
}

func (s *Info) List(ctx *gin.Context) {
	fmt.Println("api all :", ctx)
}
