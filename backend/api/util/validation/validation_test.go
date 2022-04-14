package validation

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestIsAlphaUnicodeNumericSpaceHyphen(t *testing.T) {
	type MyStruct struct {
		String string `validate:"alphaUnicodeNumericSpaceHyphen"`
	}

	validate := validator.New()
	validate.RegisterValidation("alphaUnicodeNumericSpaceHyphen", isAlphaUnicodeNumericSpaceHyphen)

	testInputs := map[string]bool{
		"hello 1":      true,
		"; drop table": false,
		"\t\n":         true,
		"ngày mai \t":  true,
		"123":          true,
		"123\t":        true,
		"123;":         false,
		"123--":        true,
		"--123":        true,
		"; --":         false,
	}
	for k, v := range testInputs {
		s := MyStruct{String: k}
		err := validate.Struct(s)
		if v {
			assert.NoError(t, err, k)
		} else {
			assert.Error(t, err, k)
		}
	}
}
