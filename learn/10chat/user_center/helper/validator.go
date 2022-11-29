package helper

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
	if err := v.Struct(sourceData); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, msg := range validationErrors.Translate(trans) {
			return errors.New(msg)
		}
	}
	return nil
}

func make() {
	once.Do(func() {
		v = validator.New()
		zh := zh.New()
		uni := ut.New(zh)
		trans, _ = uni.GetTranslator("zh")
		_ = zh_translations.RegisterDefaultTranslations(v, trans)
	})
}
