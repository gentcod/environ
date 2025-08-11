package environ

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	envFileErr      = "error reading env file: %w"
	envVariableErr  = "error parsing line in env file: %q"
	unsupportedErr  = "unsupported type for config struct field %s"
	notStructPtrErr = "config must be a non-nil pointer to a struct"
)

// Init initializes the environment configurations and returns an error if it occurs.
func Init(path string, conc any) error {
	return loadConfig(path, conc)
}

// loadConfig takes the file path of the .env file and parses the values to the config struct
func loadConfig(filepath string, config any) error {
	val := reflect.ValueOf(config)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return errors.New(notStructPtrErr)
	}
	elem := val.Elem()
	if elem.Kind() != reflect.Struct {
		return errors.New(notStructPtrErr)
	}

	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf(envFileErr, err)
	}
	defer file.Close()

	envMap := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf(envVariableErr, line)
		}
		envMap[parts[0]] = strings.Trim(parts[1], `"'`)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf(envFileErr, err)
	}

	structType := elem.Type()
	for i := 0; i < elem.NumField(); i++ {
		fieldVal := elem.Field(i)
		fieldType := structType.Field(i)

		envKey := fieldType.Tag.Get("json")
		if envKey == "" {
			continue
		}

		if value, ok := envMap[envKey]; ok {
			if err := setField(fieldVal, value); err != nil {
				return fmt.Errorf("error setting field %s: %w", fieldType.Name, err)
			}
		}
	}
	return nil
}

// setField casts the string value to as the appropriate field value type
func setField(field reflect.Value, value string) error {
	if !field.CanSet() {
		return nil
	}

	if field.Type() == reflect.TypeOf(time.Duration(0)) {
		d, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		field.SetInt(int64(d))
		return nil
	}

	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		field.SetInt(intValue)
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		field.SetBool(boolValue)
	default:
		return fmt.Errorf(unsupportedErr, field.Type().String())
	}
	return nil
}
