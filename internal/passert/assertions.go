package passert

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func FieldsInitialized(t *testing.T, value interface{}, msg string) {
	t.Helper()

	rawValue := reflect.ValueOf(value)

	var reflectValue reflect.Value

	switch rawValue.Kind() { //nolint:exhaustive
	case reflect.Struct:
		reflectValue = rawValue
	case reflect.Ptr:
		reflectValue = rawValue.Elem()
	default:
		return
	}

	for i := 0; i < reflectValue.NumField(); i++ {
		f := reflectValue.Field(i)
		assert.False(t, f.IsZero(), "%s: must be not empty %s", msg, f.String())
	}
}
