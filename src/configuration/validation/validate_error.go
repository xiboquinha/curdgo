package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/xiboquinha/curdgo/src/configuration/erros"
)

var (
	validate = validator.New()
	traduz   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		un := ut.New(en, en)
		traduz, _ = un.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, traduz)
	}
}

func ValidateUserError(
	validation_err error,
) *erros.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return erros.NewBadRequestErr("erro na tipagem do campo")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorCauses := []erros.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := erros.Causes{
				Message: e.Translate(traduz),
				Field:   e.Field(),
			}

			errorCauses = append(errorCauses, cause)

		}

		return erros.NewBadRequestValidationErr("alguns campos estao invalidos: ", errorCauses)
	} else {
		return erros.NewBadRequestErr("erro tentando converter campos")
	}

}
