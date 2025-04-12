package validators

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_tranlations "github.com/go-playground/validator/v10/translations/en"
)

var validate *validator.Validate
var trans ut.Translator

func init() {
	validate = validator.New()
	validate.RegisterTagNameFunc(jsonTagName)

	en := en.New()
	uni := ut.New(en, en)
	trans, _ = uni.GetTranslator("en")
	en_tranlations.RegisterDefaultTranslations(validate, trans)

	validate.RegisterValidation("fkcategoryid", fkcategoryIDExists)
	validate.RegisterTranslation("fkcategoryid", trans, func(ut ut.Translator) error {
		return ut.Add("fkcategoryid", "{0} is not a valid category", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("fkcategoryid", fe.Field())
		return t
	})
}

func jsonTagName(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	// skip if tag key says it should be ignored
	if name == "-" {
		return ""
	}
	return name
}

func ValidateStruct(data any) (map[string]any, bool) {
	errorFields := make(map[string][]string)
	fieldErrors := validate.Struct(data)
	valid := true
	if fieldErrors != nil {
		valid = false
		for _, e := range fieldErrors.(validator.ValidationErrors) {
			errmsg := e.Translate(trans)
			field := e.Field()

			if msgs, ok := errorFields[field]; !ok {
				errorFields[field] = []string{errmsg}
			} else {
				errorFields[field] = append(msgs, errmsg)
			}
		}
	}
	m := structToMap(data)
	m["valid"] = valid
	m["errors"] = errorFields
	return m, valid
}

func structToMap(c any) map[string]any {
	var m map[string]any
	b, _ := json.Marshal(c)
	_ = json.Unmarshal(b, &m)
	return m
	// return c.(map[string]any)
}
