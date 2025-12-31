package config

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Load populates the exported fields of a struct from environment variables using
// struct tags. Supported tags:
//   - `env:"VAR_NAME"` (required to map a field)
//   - `default:"value"` (optional default when env var missing)
//   - `required:"true"` (treat missing as error)
//   - `expand:"true"` (apply os.ExpandEnv on the value)
//
// returns a single aggregated error listing missing env vars or parse errors.
func Load(e *Env) error {
	value := reflect.ValueOf(e)
	if value.Kind() != reflect.Ptr || value.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("Environment loader requires a pointer to a struct")
	}
	valElem := value.Elem()
	valType := valElem.Type()

	var missing []string
	var parseErrs []string

	for i := 0; i < valElem.NumField(); i++ {
		valField := valElem.Field(i)
		typeField := valType.Field(i)

		// only settable exported fields
		if !valField.CanSet() {
			continue
		}

		envKey := typeField.Tag.Get("env")
		if envKey == "" {
			// skip fields without env tag
			continue
		}

		raw, ok := os.LookupEnv(envKey)
		def := typeField.Tag.Get("default")
		req := typeField.Tag.Get("required") == "true"
		expand := typeField.Tag.Get("expand") == "true"

		if !ok || raw == "" {
			if def != "" {
				raw = def
			} else if req {
				missing = append(missing, envKey)
				continue
			} else {
				// leave zero value
				continue
			}
		}

		if expand {
			raw = os.ExpandEnv(raw)
		}

		switch valField.Kind() {
		case reflect.String:
			valField.SetString(raw)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			iv, err := strconv.ParseInt(raw, 10, 64)
			if err != nil {
				parseErrs = append(parseErrs, fmt.Sprintf("%s: %v", envKey, err))
				continue
			}
			valField.SetInt(iv)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			uv, err := strconv.ParseUint(raw, 10, 64)
			if err != nil {
				parseErrs = append(parseErrs, fmt.Sprintf("%s: %v", envKey, err))
				continue
			}
			valField.SetUint(uv)
		case reflect.Bool:
			bv, err := strconv.ParseBool(raw)
			if err != nil {
				parseErrs = append(parseErrs, fmt.Sprintf("%s: %v", envKey, err))
				continue
			}
			valField.SetBool(bv)
		default:
			// unsupported field type (arrays, slices, structs) are ignored for now
			continue
		}
	}

	if len(missing) > 0 || len(parseErrs) > 0 {
		var parts []string
		if len(missing) > 0 {
			parts = append(parts, "missing required environment variables: \n"+strings.Join(missing, ", "))
		}
		if len(parseErrs) > 0 {
			parts = append(parts, "parse errors: "+strings.Join(parseErrs, "; "))
		}
		return errors.New(strings.Join(parts, "; "))
	}

	return nil
}
