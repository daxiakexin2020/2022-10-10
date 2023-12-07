package request

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Request struct {
}

func CheckJson(ctx *gin.Context, dest interface{}) error {
	if err := ShouldBindJSON(ctx, &dest); err != nil {
		return err
	}
	return ValidateCondition(dest)
}

func Check(ctx *gin.Context, dest interface{}) error {
	if err := Bind(ctx, dest); err != nil {
		return err
	}
	return ValidateCondition(dest)
}

func ShouldBindJSON(ctx *gin.Context, dest interface{}) error {
	return ctx.ShouldBindJSON(dest)
}

func Bind(ctx *gin.Context, dest interface{}) error {
	return ctx.Bind(dest)
}

func ValidateCondition(condition interface{}) error {
	v := validator.New()
	if err := v.Struct(condition); err != nil {
		return err
	}
	return nil
}
