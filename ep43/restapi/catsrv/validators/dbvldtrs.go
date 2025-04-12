package validators

import (
	"rfcrud/dal/catalog"

	"github.com/go-playground/validator/v10"
)

func fkcategoryIDExists(fl validator.FieldLevel) bool {
	catid := fl.Field().Int()
	found := true

	_, err := catalog.GetCategory(int(catid))
	if err != nil {
		found = false
	}
	return found
}
