package validator

import (
	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"sync"
)

var (
	v     *validator.Validate
	once  sync.Once
	trans ut.Translator
)

func Check(sourceData interface{}) error {

	if v == nil {
		make()
	}

	var msg string
	if err := v.Struct(sourceData); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, msg = range validationErrors.Translate(trans) { //map  取第一条错误信息提示...
			return errors.New(msg)
		}
	}
	return nil
}

func make() {
	once.Do(func() {
		if v == nil {
			v = validator.New()
			zh := zh.New()
			uni := ut.New(zh)
			trans, _ = uni.GetTranslator("zh")
			_ = zh_translations.RegisterDefaultTranslations(v, trans)
		}
	})
}

func inCheck(sourceData interface{}) (error, interface{}) {

	validate := validator.New()

	zh := zh.New()
	uni := ut.New(zh)
	ts, _ := uni.GetTranslator("zh")
	_ = zh_translations.RegisterDefaultTranslations(validate, ts)

	if err := validate.Struct(sourceData); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return err, validationErrors.Translate(trans)
	}

	return nil, nil
}
