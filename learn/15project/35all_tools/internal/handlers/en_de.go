package handlers

import (
	cerr "35all_tools/error"
	"35all_tools/internal/model"
	"35all_tools/internal/service"
	"github.com/gin-gonic/gin"
)

type EnDeHandler struct {
	*Base
	endeService *service.EnDeService
}

func NewEnDeHandler(eds *service.EnDeService) *EnDeHandler {
	return &EnDeHandler{endeService: eds}
}

func (edh *EnDeHandler) MD5Encode(ctx *gin.Context) {
	req := model.NewMD5Encode()
	if err := edh.ShouldBindJSON(ctx, req); err != nil {
		edh.Err(ctx, cerr.PARAMS_BIND_ERR_CODE, err.Error(), nil)
		return
	}
	md5Encode, err := edh.endeService.Md5Encode(req)
	if err != nil {
		edh.Err(ctx, cerr.MD5_ENCODE_ERR, err.Error(), nil)
	}
	edh.Success(ctx, md5Encode)
}

func (edh *EnDeHandler) Url16Encode(ctx *gin.Context) {
	req := model.NewUrl16Encode()
	if err := edh.ShouldBindJSON(ctx, req); err != nil {
		edh.Err(ctx, cerr.PARAMS_BIND_ERR_CODE, err.Error(), nil)
		return
	}
	encode, err := edh.endeService.UrlEncode(req)
	if err != nil {
		edh.Err(ctx, cerr.URL_ENCODE_ERR, err.Error(), nil)
		return
	}
	edh.Success(ctx, encode)
}

func (edh *EnDeHandler) Base64Encode(ctx *gin.Context) {
	req := model.NewBase64Encode()
	if err := edh.ShouldBindJSON(ctx, req); err != nil {
		edh.Err(ctx, cerr.PARAMS_BIND_ERR_CODE, err.Error(), nil)
		return
	}
	encode, err := edh.endeService.Base64Encode(req)
	if err != nil {
		edh.Err(ctx, cerr.BASE64_ENCODE_ERR, err.Error(), nil)
	}
	edh.Success(ctx, encode)
}

func (edh *EnDeHandler) Base64Decode(ctx *gin.Context) {
	req := model.NewBase64Decode()
	if err := edh.ShouldBindJSON(ctx, req); err != nil {
		edh.Err(ctx, cerr.PARAMS_BIND_ERR_CODE, err.Error(), nil)
		return
	}
	encode, err := edh.endeService.Base64Decode(req)
	if err != nil {
		edh.Err(ctx, cerr.BASE64_ENCODE_ERR, err.Error(), nil)
	}
	edh.Success(ctx, encode)
}

func (edh *EnDeHandler) Escape(ctx *gin.Context) {
	req := model.NewEscape()
	if err := edh.ShouldBindJSON(ctx, req); err != nil {
		edh.Err(ctx, cerr.PARAMS_BIND_ERR_CODE, err.Error(), nil)
		return
	}
	encode, err := edh.endeService.Escape(req)
	if err != nil {
		edh.Err(ctx, cerr.ESCAPE_ERR, err.Error(), nil)
	}
	edh.Success(ctx, encode)
}

func (edh *EnDeHandler) DeEscape(ctx *gin.Context) {
	req := model.NewDeEscape()
	if err := edh.ShouldBindJSON(ctx, req); err != nil {
		edh.Err(ctx, cerr.PARAMS_BIND_ERR_CODE, err.Error(), nil)
		return
	}
	encode, err := edh.endeService.DeEscape(req)
	if err != nil {
		edh.Err(ctx, cerr.DEESCAPE_ERR, err.Error(), nil)
	}
	edh.Success(ctx, encode)
}
