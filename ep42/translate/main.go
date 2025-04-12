package main

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_tranlations "github.com/go-playground/validator/v10/translations/en"
)

var validate *validator.Validate
var trans ut.Translator

func main() {
	user := struct {
		Name     string `json:"name" validate:"required,gte=6,lte=15"`
		Age      int    `json:"age" validate:"required,gte=18,lte=55"`
		Password string `json:"password" validate:"required,nontrivial,password,min=6,max=15"`
	}{
		Name:     "balas",
		Age:      17,
		Password: "abcde",
	}

	fe := validateUser(user)
	for field, errmsg := range fe {
		for _, msg := range errmsg {
			println(field, msg)
		}
	}
}

func init() {

	validate = validator.New()
	validate.RegisterTagNameFunc(jsonTagName)

	en := en.New()
	uni := ut.New(en, en)
	trans, _ = uni.GetTranslator("en")
	en_tranlations.RegisterDefaultTranslations(validate, trans)

	validate.RegisterValidation("password", PasswordValidator)
	validate.RegisterTranslation("password",
		trans,
		func(ut ut.Translator) error {
			return ut.Add("password", "{0} is not valid. Should have a lowercase, uppercase, number and any one of !@#$%^&*()", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("password", fe.Field())
			return t
		})
	validate.RegisterValidation("nontrivial", TrivialValidator)
	validate.RegisterTranslation("nontrivial",
		trans,
		func(ut ut.Translator) error {
			return ut.Add("nontrivial", "{0} is the same as user name", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("nontrivial", fe.Field())
			return t
		})
}

func validateUser(user any) map[string][]string {
	errorFields := make(map[string][]string)
	fieldErrors := validate.Struct(user)

	if fieldErrors != nil {
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
	return errorFields
}

func jsonTagName(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	// skip if tag key says it should be ignored
	if name == "-" {
		return ""
	}
	return name
}

func PasswordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	//check if password has uppercase letter
	hasUpper := false
	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			hasUpper = true
			break
		}
	}
	//check if password has lowercase letter
	hasLower := false
	for _, char := range password {
		if char >= 'a' && char <= 'z' {
			hasLower = true
			break
		}
	}
	//check if password has number
	hasNumber := false
	for _, char := range password {
		if char >= '0' && char <= '9' {
			hasNumber = true
			break
		}
	}

	hasSpecial := strings.ContainsAny(password, "!@#$%^&*()")

	return hasUpper && hasLower && hasNumber && hasSpecial
}

// validator function to check if password is not equal to the name

func TrivialValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	name := fl.Parent().FieldByName("Name").String()
	return password != name
}
