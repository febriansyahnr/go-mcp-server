package validatorExt

import (
	"regexp"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslator "github.com/go-playground/validator/v10/translations/en"
)

func New() *validator.Validate {
	validate := validator.New()

	enTranslator.RegisterDefaultTranslations(validate, GetTranslator())

	// Register custom validation for virtual account numbers that allows digits and spaces
	validate.RegisterValidation("va_number", validateVANumber)

	return validate
}

// validateVANumber validates that a field contains only numeric characters (0-9) 
// with optional spaces at the beginning or end only
func validateVANumber(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	
	// Check if field is not empty
	if field == "" {
		return false
	}
	
	// Check if field matches pattern: optional leading spaces + digits + optional trailing spaces
	matched, _ := regexp.MatchString(`^\s*[0-9]+\s*$`, field)
	return matched
}

func GetTranslator() ut.Translator {
	en := en.New()
	uni := ut.New(en, en)

	trans, _ := uni.GetTranslator("en")

	return trans
}
