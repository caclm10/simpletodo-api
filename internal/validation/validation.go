package validation

import (
	"net/http"
	"reflect"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo/v4"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationErrors = []ValidationError

type Validator struct {
	validate *validator.Validate
	uni      *ut.UniversalTranslator
	trans    ut.Translator
}

func NewValidator() *Validator {
	en := en.New()
	uni := ut.New(en, en)

	trans, _ := uni.GetTranslator("en")

	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	return &Validator{validate, uni, trans}
}

func (v *Validator) Validate(s any) error {
	if err := v.validate.Struct(s); err != nil {
		errs := err.(validator.ValidationErrors)

		result := make(ValidationErrors, len(errs))
		for i, ve := range errs {
			sf, _ := reflect.TypeOf(s).Elem().FieldByName(ve.Field())

			result[i] = ValidationError{Field: sf.Tag.Get("json"), Message: ve.Translate(v.trans)}
		}

		return echo.NewHTTPError(http.StatusUnprocessableEntity, echo.Map{
			"code":   http.StatusUnprocessableEntity,
			"status": "Error",
			"errors": result,
		})
	}

	return nil
}
