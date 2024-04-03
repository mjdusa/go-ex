package reflectex

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	DefaultIndentSize int = 2
)

func Reflector(a any) string {
	val := reflect.ValueOf(a)

	return reflector(val, val.Type().Name(), 0, DefaultIndentSize)
}

func makePrefix(prefix string, val reflect.Value) string {
	if val.Kind() == reflect.Ptr {
		return prefix + ".*" + val.Type().Name()
	}
	return prefix + "." + val.Type().Name()
}

func makeStructPrefix(prefix string, val reflect.StructField) string {
	if val.Type.Kind() == reflect.Ptr {
		return prefix + ".*" + val.Name
	}
	return prefix + "." + val.Name
}

//nolint:funlen  // function is long but it's a big switch and breaking out cases would make it unnecessarily complex
func reflector(val reflect.Value, prefix string, depth int, indentSize int) string {
	msg := ""
	indent := indentSize * depth
	spaces := strings.Repeat(" ", indent)

	valType := val.Type()

	//nolint:exhaustive  // handling of remaining types is not necessary
	switch val.Kind() {
	case reflect.Array:
		msg += fmt.Sprintf("%s%s (Array: %d items)\n", spaces, prefix, val.Len())
		for idx := 0; idx < val.Len(); idx++ {
			f := val.Index(idx)
			msg += reflector(f, makePrefix(prefix, f), depth+1, indentSize)
		}
	case reflect.Bool:
		msg += fmt.Sprintf("%s%s: %v (%s)\n", spaces, prefix, val.Bool(), val.Kind())
	case reflect.Complex64:
		msg += fmt.Sprintf("%s%s: %v (%s)\n", spaces, prefix, val.Complex(), val.Kind())
	case reflect.Complex128:
		msg += fmt.Sprintf("%s%s: %v (%s)\n", spaces, prefix, val.Complex(), val.Kind())
	case reflect.Float32:
		msg += fmt.Sprintf("%s%s: %v (%s)\n", spaces, prefix, val.Float(), val.Kind())
	case reflect.Float64:
		msg += fmt.Sprintf("%s%s: %v (%s)\n", spaces, prefix, val.Float(), val.Kind())
	case reflect.Int:
		msg += fmt.Sprintf("%s%s: %v (%s)\n", spaces, prefix, val.Int(), val.Kind())
	case reflect.Int8:
		msg += fmt.Sprintf("%s%s: %v (%s)\n", spaces, prefix, val.Int(), val.Kind())
	case reflect.Int16:
		msg += fmt.Sprintf("%s%s: %v (%s)\n", spaces, prefix, val.Int(), val.Kind())
	case reflect.Int32:
		msg += fmt.Sprintf("%s%s: %v (%s)\n", spaces, prefix, val.Int(), val.Kind())
	case reflect.Int64:
		msg += fmt.Sprintf("%s%s: %v (%s)\n", spaces, prefix, val.Int(), val.Kind())
	case reflect.Map:
		msg += fmt.Sprintf("%s%s (Map: %d items)\n", spaces, prefix, val.Len())
		keys := val.MapKeys()
		for idx := 0; idx < val.Len(); idx++ {
			key := keys[idx]
			f := val.MapIndex(key)
			msg += reflector(f, makePrefix(prefix, f), depth+1, indentSize)
		}
	case reflect.Ptr:
		if !val.IsNil() {
			f := val.Elem()
			msg += reflector(f, prefix, depth, indentSize) // pointer shouldn't change the prefix or depth
		}
	case reflect.Slice:
		msg += fmt.Sprintf("%s%s (Slice: %d items)\n", spaces, prefix, val.Len())
		for idx := 0; idx < val.Len(); idx++ {
			f := val.Index(idx)
			msg += reflector(f, makePrefix(prefix, f), depth+1, indentSize)
		}
	case reflect.Struct:
		msg += fmt.Sprintf("%s%s (Struct: %d fields)\n", spaces, prefix, val.NumField())
		for idx := 0; idx < val.NumField(); idx++ {
			f := val.Field(idx)
			ft := valType.Field(idx)
			msg += reflector(f, makeStructPrefix(prefix, ft), depth+1, indentSize)
		}
	case reflect.String:
		msg += fmt.Sprintf("%s%s: %v (%s)\n", spaces, prefix, val.String(), val.Kind())
	case reflect.Uint:
		msg += fmt.Sprintf("%s%s: %v (%s)\n", spaces, prefix, val.Uint(), val.Kind())
	case reflect.Uint8:
		msg += fmt.Sprintf("%s%s: %v (%s)\n", spaces, prefix, val.Uint(), val.Kind())
	case reflect.Uint16:
		msg += fmt.Sprintf("%s%s: %v (%s)\n", spaces, prefix, val.Uint(), val.Kind())
	case reflect.Uint32:
		msg += fmt.Sprintf("%s%s: %v (%s)\n", spaces, prefix, val.Uint(), val.Kind())
	case reflect.Uint64:
		msg += fmt.Sprintf("%s%s: %v (%s)\n", spaces, prefix, val.Uint(), val.Kind())
	default:
		msg += fmt.Sprintf("%s%s: %v (%s)\n", spaces, prefix, val.String(), val.Kind())
	}

	return msg
}
