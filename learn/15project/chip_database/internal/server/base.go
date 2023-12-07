package server

import (
	"chip_database/internal/request"
	"chip_database/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Base struct{}

func (b *Base) CheckJson(ctx *gin.Context, dest interface{}) error {
	return request.CheckJson(ctx, dest)
}

func (b *Base) Check(ctx *gin.Context, dest interface{}) error {
	return request.Check(ctx, dest)
}

func (b *Base) ValidateCondition(condition interface{}) error {
	v := validator.New()
	if err := v.Struct(condition); err != nil {
		return err
	}
	return nil
}

func (b *Base) Success(ctx *gin.Context, data interface{}) {
	response.Success(ctx, data)
}

func (b *Base) Err500(ctx *gin.Context) {
	response.Err500(ctx)
}

func (b *Base) Err404(ctx *gin.Context) {
	response.Err404(ctx)
}

func (b *Base) ParamErr(ctx *gin.Context, msg string, data interface{}) {
	response.ParamErr(ctx, msg, data)
}

func (b Base) DBErr(ctx *gin.Context, msg string, data interface{}) {
	response.DBErr(ctx, msg, data)
}

func (b *Base) Err(ctx *gin.Context, code int32, msg string, data interface{}) {
	response.Err(ctx, code, msg, data)
}
