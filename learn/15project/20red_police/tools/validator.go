package tools

import "github.com/go-playground/validator/v10"

var DefaultV = validator.New()

func Validator(dest interface{}) error {
	if err := DefaultV.Struct(dest); err != nil {
		return err
	}
	return nil
}
