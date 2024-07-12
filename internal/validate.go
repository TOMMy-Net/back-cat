package internal

import (
	"github.com/go-playground/validator/v10"
)

func validate(s any) error {
	val := validator.New()

	var err = val.Struct(s)
	if err != nil {
		return err
	}
	return nil
}
