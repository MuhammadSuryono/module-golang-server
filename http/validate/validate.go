package validate

import (
	"github.com/MuhammadSuryono/module-golang-server/http/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"reflect"
)

type Validate struct {
	Validator  *validator.Validate
	Translator ut.Translator
	S          interface{}
}

type Context struct {
	ctx *gin.Context
}

type ResponseError struct {
	Error         interface{}
	dataStruct    interface{}
	errorValidate error
}

func translateError(err error, trans ut.Translator) (errs []string) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := e.Translate(trans)
		errs = append(errs, translatedErr)
	}
	return errs
}

func registerTranslation(v *validator.Validate) ut.Translator {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(v, trans)
	return trans
}

func NewValidate() Validate {
	validateForm := validator.New()
	trans := registerTranslation(validateForm)
	return Validate{
		Validator:  validateForm,
		Translator: trans,
	}
}

func (v Validate) ValidationStruct(s interface{}) ResponseError {
	err := v.Validator.Struct(s)
	return ResponseError{
		Error:         translateError(err, v.Translator),
		dataStruct:    s,
		errorValidate: err,
	}
}

func (r ResponseError) JsonResponse(c *gin.Context) {
	if r.errorValidate != nil {
		c.AbortWithStatusJSON(response.BAD_REQUEST_CODE, response.FailureResponse(
			response.BAD_REQUEST_STATUS,
			response.BAD_REQUEST_MESSAGE,
			r.Error,
		))
		return

	} else {
		c.Set("data", r.dataStruct)
		c.Next()
	}

}

func KindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

func (v Validate) RegisterStruct(s interface{}) Validate {
	v.S = s
	return v
}
