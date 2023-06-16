package handler

import (
	cerr "35all_tools/error"
	"35all_tools/internal/model"
	"35all_tools/internal/service"
	"github.com/gin-gonic/gin"
)

type JsonHanler struct {
	*Base
	jsonservice *service.JsonSerive
}

func NewJsonHandler(jsonservice *service.JsonSerive, base *Base) *JsonHanler {
	return &JsonHanler{jsonservice: jsonservice, Base: base}
}

func (jh *JsonHanler) CheckJson(ctx *gin.Context) {
	req := model.NewJsonCheck()
	if err := jh.ShouldBindJSON(ctx, req); err != nil {
		jh.Err(ctx, cerr.PARAMS_BIND_ERR_CODE, err.Error(), nil)
		return
	}
	check, err := jh.jsonservice.JsonCheck(req)
	if err != nil {
		jh.Err(ctx, cerr.NOT_JSON_CODE, err.Error(), nil)
		return
	}
	jh.Success(ctx, check)
}

func (jh *JsonHanler) JsonToGolangStruct(ctx *gin.Context) {
	req := model.NewJsonToGolangStruct()
	if err := jh.ShouldBindJSON(ctx, req); err != nil {
		jh.Err(ctx, cerr.PARAMS_BIND_ERR_CODE, err.Error(), nil)
		return
	}
	checkModel := model.NewJsonCheck()
	checkModel.Str = req.Str
	_, err := jh.jsonservice.JsonCheck(checkModel)
	if err != nil {
		jh.Err(ctx, cerr.NOT_JSON_CODE, err.Error(), nil)
		return
	}
	golangStruct, err := jh.jsonservice.JsonToGolangStruct(req)
	if err != nil {
		jh.Err(ctx, cerr.JSON_TO_GOALNG_ERR_CODE, err.Error(), nil)
		return
	}
	ctx.String(200, "%s", golangStruct)
}
