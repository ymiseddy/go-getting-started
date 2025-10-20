package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func ReadConfigInto(target any) error {
	// Load environment variables from .env file if it exists
	_ = godotenv.Load()

	val := reflect.ValueOf(target).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		structField := typ.Field(i)

		tag := structField.Tag.Get("env")
		if tag == "" {
			continue
		}

		// Parse the tag
		parts := strings.Split(tag, ",")
		envVar := parts[0]
		var defaultValue string
		required := false

		for _, part := range parts[1:] {
			if part == "required" {
				required = true
			} else if after, ok := strings.CutPrefix(part, "default="); ok {
				defaultValue = after
			}
		}

		// Get environment variable value
		envValue, exists := os.LookupEnv(envVar)
		if !exists {
			if required {
				return fmt.Errorf("required environment variable %s is not set", envVar)
			}
			envValue = defaultValue
		}

		// Set the field based on its type
		if !field.CanSet() {
			return fmt.Errorf("cannot set field %s", structField.Name)
		}

		switch field.Kind() {
		case reflect.String:
			field.SetString(envValue)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intVal, err := strconv.ParseInt(envValue, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid integer value for %s: %s", envVar, envValue)
			}
			field.SetInt(intVal)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			uintVal, err := strconv.ParseUint(envValue, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid unsigned integer value for %s: %s", envVar, envValue)
			}
			field.SetUint(uintVal)
		case reflect.Float32, reflect.Float64:
			floatVal, err := strconv.ParseFloat(envValue, 64)
			if err != nil {
				return fmt.Errorf("invalid float value for %s: %s", envVar, envValue)
			}
			field.SetFloat(floatVal)
		case reflect.Bool:
			boolVal, err := strconv.ParseBool(envValue)
			if err != nil {
				return fmt.Errorf("invalid boolean value for %s: %s", envVar, envValue)
			}
			field.SetBool(boolVal)
		default:
			return fmt.Errorf("unsupported field type %s for %s", field.Kind(), structField.Name)
		}
	}
	return nil
}
