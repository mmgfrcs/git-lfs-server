package models

import (
	"errors"
	enLocale "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type Validator struct {
	validator  *validator.Validate
	translator ut.Translator
}

func NewValidator() *Validator {
	valid := validator.New()
	enLang := enLocale.New()
	trans, _ := ut.New(enLang, enLang).GetTranslator("en")
	_ = enTrans.RegisterDefaultTranslations(valid, trans)

	return &Validator{validator: valid, translator: trans}
}

func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		var validationErrors validator.ValidationErrors
		errors.As(err, &validationErrors)
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, "Validation Error: "+mapValueToString(validationErrors.Translate(v.translator)))
	}
	return nil
}

func mapValueToString(m map[string]string) string {
	arr := make([]string, 0, len(m))
	for _, v := range m {
		arr = append(arr, v)
	}

	return strings.Join(arr, ",")
}
