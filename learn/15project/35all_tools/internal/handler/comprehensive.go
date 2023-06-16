package handler

import (
	cerr "35all_tools/error"
	"35all_tools/internal/model"
	"35all_tools/internal/service"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"io"
	"log"
)

type ComprehensiveHandler struct {
	*Base
	comprehensiveService *service.ComprehensiveService
}

func NewComprehensiveHandler(cs *service.ComprehensiveService) *ComprehensiveHandler {
	return &ComprehensiveHandler{comprehensiveService: cs}
}

func (ch *ComprehensiveHandler) IpInfo(ctx *gin.Context) {
	req := model.NewIpInfo()
	if err := ch.ShouldBindJSON(ctx, req); err != nil {
		ch.Err(ctx, cerr.PARAMS_BIND_ERR_CODE, err.Error(), nil)
		return
	}
	info, err := ch.comprehensiveService.IpInfo(req)
	if err != nil {
		ch.Err(ctx, cerr.IPINFO_ERR, err.Error(), nil)
		return
	}
	ch.Success(ctx, info)
}

func (ch *ComprehensiveHandler) DomainMapIp(ctx *gin.Context) {
	req := model.NewDomainMapIp()
	if err := ch.ShouldBindJSON(ctx, req); err != nil {
		ch.Err(ctx, cerr.PARAMS_BIND_ERR_CODE, err.Error(), nil)
		return
	}
	info, err := ch.comprehensiveService.DomainMapIp(req)
	if err != nil {
		ch.Err(ctx, cerr.DOMAIN_MAP_IP_ERR, err.Error(), nil)
		return
	}
	ch.Success(ctx, info)
}

func (ch *ComprehensiveHandler) ImageBase64(ctx *gin.Context) {
	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		ch.Err(ctx, cerr.IMAGE_BASE64_ERR, err.Error(), nil)
		return
	}

	data, err := io.ReadAll(file)
	if err != nil {
		ch.Err(ctx, cerr.IMAGE_BASE64_ERR, err.Error(), nil)
		return
	}
	log.Println("Image Base64..............")
	ch.Success(ctx, base64.StdEncoding.EncodeToString(data))
}
