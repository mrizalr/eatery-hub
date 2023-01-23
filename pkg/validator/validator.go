package validator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validator *validator.Validate
	trans     ut.Translator
)

func init() {
	en := en.New()
	uni := ut.New(en, en)

	trans, _ = uni.GetTranslator("en")
	Validator = validator.New()

	en_translations.RegisterDefaultTranslations(Validator, trans)
}

func TranslateErrors(err error) []string {
	errs := []string{}
	validationErrors := err.(validator.ValidationErrors)
	errTranslate := validationErrors.Translate(trans)

	for _, msg := range errTranslate {
		errs = append(errs, msg)
	}

	return errs
}
