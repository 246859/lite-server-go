package validateutils

import validator2 "github.com/go-playground/validator/v10"

var validator = validator2.New()

func EmailCheck(email string) bool {
	return validator.Var(email, "email") == nil
}
