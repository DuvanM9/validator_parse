package validatestruct

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

// camelToSnake convierte un nombre en camelCase a snake_case
func camelToSnake(name string) string {
	var re = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := re.ReplaceAllString(name, "${1}_${2}")
	return strings.ToLower(snake)
}

func FormatValidationError(err error) string {
	errs := err.(validator.ValidationErrors)
	fieldName := camelToSnake(errs[0].Field())
	jsonTag := camelToSnake(errs[0].StructField())
	actualTag := errs[0].ActualTag()
	param := errs[0].Param()

	switch actualTag {
	case "oneof":
		return fmt.Sprintf("%s: must be one of [%s]", jsonTag, param)
	case "required":
		return fmt.Sprintf("%s: %s is required", jsonTag, fieldName)
	default:
		return err.Error()
	}
}
