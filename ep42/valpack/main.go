package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func main() {
	user := struct {
		//step 1
		// Name     string `validate:"required,gte=6,lte=15"`
		// Age      int    `validate:"required,gte=18,lte=55"`
		// Password string `validate:"required,min=6,max=15"`
		//step 2
		Name     string `json:"name" validate:"required,gte=6,lte=15"`
		Age      int    `json:"age" validate:"required,gte=18,lte=55"`
		Password string `json:"password" validate:"required,min=6,max=15"`
	}{
		Name:     "balas",
		Age:      17,
		Password: "abcde",
	}

	//Step 1 & 2
	validateUser(user)

}

func init() {
	//step 1
	validate = validator.New()

	//step 2
	validate.RegisterTagNameFunc(jsonTagName)
}

func validateUser(user any) {
	fieldErrors := validate.Struct(user)
	//step 1
	printRawErrors(fieldErrors)
}

func printRawErrors(fieldErrors error) {
	if fieldErrors != nil {
		for _, err := range fieldErrors.(validator.ValidationErrors) {
			fmt.Println("Namespace: ", err.Namespace())
			fmt.Println("Field: ", err.Field())
			fmt.Println("StructNamespace: ", err.StructNamespace())
			fmt.Println("StructField: ", err.StructField())
			fmt.Println("Tag: ", err.Tag())
			fmt.Println("ActualTag: ", err.ActualTag())
			fmt.Println("Kind: ", err.Kind())
			fmt.Println("Type: ", err.Type())
			fmt.Println("Value: ", err.Value())
			fmt.Println("Param: ", err.Param())
			fmt.Println("----------------------")
		}
	}
}

// Step 2
func jsonTagName(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	// skip if tag key says it should be ignored
	if name == "-" {
		return ""
	}
	return name
}
