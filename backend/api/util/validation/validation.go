package validation

import (
	"errors"
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

const (
	alphaUnicodeNumericSpaceHyphenRegexString = "^[\\p{L}\\p{N}\\-\\s]+$"
)

var (
	alphaUnicodeNumericSpaceHyphenRegex = regexp.MustCompile(alphaUnicodeNumericSpaceHyphenRegexString)
)

// isAlphaUnicodeNumericSpaceHyphen is the validation function for validating if the current field's value contains only: unicode alphanumeric characters, whitespace (tab, space, etc) and hyphen (-).
func isAlphaUnicodeNumericSpaceHyphen(fl validator.FieldLevel) bool {
	return alphaUnicodeNumericSpaceHyphenRegex.MatchString(fl.Field().String())
}

func Register() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		return v.RegisterValidation("alphaUnicodeNumericSpaceHyphen", isAlphaUnicodeNumericSpaceHyphen)
	}
	return errors.New("failed on registering custom validation rules")
}
