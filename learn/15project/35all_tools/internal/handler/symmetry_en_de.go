package handler

import (
	cerr "35all_tools/error"
	"35all_tools/internal/model"
	"35all_tools/internal/service"
	"github.com/gin-gonic/gin"
)

type SymmetryEnDeHandler struct {
	*Base
	symmetryEnDeService *service.SymmetryEnDeService
}

func NewSymmetryEnDeHandler(seds *service.SymmetryEnDeService) *SymmetryEnDeHandler {
	return &SymmetryEnDeHandler{symmetryEnDeService: seds}
}

func (sedh *SymmetryEnDeHandler) Encode(ctx *gin.Context) {
	req := model.NewSymmetryEnDeEncode()
	if err := sedh.ShouldBindJSON(ctx, req); err != nil {
		sedh.Err(ctx, cerr.PARAMS_BIND_ERR_CODE, err.Error(), nil)
		return
	}
	encode, err := sedh.symmetryEnDeService.Encode(req)
	if err != nil {
		sedh.Err(ctx, cerr.SYMMETRY_ENCODE_ERR, err.Error(), nil)
		return
	}
	sedh.Success(ctx, encode)
}

func (sedh *SymmetryEnDeHandler) Decode(ctx *gin.Context) {

}
