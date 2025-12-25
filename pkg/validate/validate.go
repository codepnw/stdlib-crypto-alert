package validate

import "github.com/go-playground/validator"

func Struct(input any) error {
	v := validator.New()
	return v.Struct(input)
}
