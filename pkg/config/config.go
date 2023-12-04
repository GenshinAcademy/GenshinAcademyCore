package config

import (
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	defaultTypeParsers = map[reflect.Type]ParserFunc{
		reflect.TypeOf(time.Duration(0)): func(v string) (interface{}, error) {
			return time.ParseDuration(v)
		},
	}

	defaultKindParsers = map[reflect.Kind]ParserFunc{
		reflect.Bool: func(v string) (interface{}, error) {
			return strconv.ParseBool(v)
		},
		reflect.String: func(v string) (interface{}, error) {
			return v, nil
		},
		reflect.Int: func(v string) (interface{}, error) {
			i, err := strconv.ParseInt(v, 10, 32)
			return int(i), err
		},
		reflect.Int16: func(v string) (interface{}, error) {
			i, err := strconv.ParseInt(v, 10, 16)
			return int16(i), err
		},
		reflect.Int32: func(v string) (interface{}, error) {
			i, err := strconv.ParseInt(v, 10, 32)
			return int32(i), err
		},
		reflect.Int64: func(v string) (interface{}, error) {
			return strconv.ParseInt(v, 10, 64)
		},
		reflect.Int8: func(v string) (interface{}, error) {
			i, err := strconv.ParseInt(v, 10, 8)
			return int8(i), err
		},
		reflect.Uint: func(v string) (interface{}, error) {
			i, err := strconv.ParseUint(v, 10, 32)
			return uint(i), err
		},
		reflect.Uint16: func(v string) (interface{}, error) {
			i, err := strconv.ParseUint(v, 10, 16)
			return uint16(i), err
		},
		reflect.Uint32: func(v string) (interface{}, error) {
			i, err := strconv.ParseUint(v, 10, 32)
			return uint32(i), err
		},
		reflect.Uint64: func(v string) (interface{}, error) {
			i, err := strconv.ParseUint(v, 10, 64)
			return i, err
		},
		reflect.Uint8: func(v string) (interface{}, error) {
			i, err := strconv.ParseUint(v, 10, 8)
			return uint8(i), err
		},
		reflect.Float64: func(v string) (interface{}, error) {
			return strconv.ParseFloat(v, 64)
		},
		reflect.Float32: func(v string) (interface{}, error) {
			f, err := strconv.ParseFloat(v, 32)
			return float32(f), err
		},
	}
)

type ParserFunc func(v string) (interface{}, error)

// UnmarshalFromEnv looking for environment variables and assigns their values
// to corresponding fields of a structure using "mapstructure" tags.
//
// It was designed for the maximum compatibility with https://github.com/spf13/viper.
//
// The 'cfg' argument is a pointer to the structure where values will be placed.
// The function recursively traverses the fields of the structure and its nested structures,
// which means structure can contain as much nested structures as you want.
//
// You may use omitempty tags to allow empty fields. STRINGS ONLY!
//
// Example:
//
//	type TestConfig struct {
//		BoolField				bool			`mapstructure:"BOOL_FIELD"`
//		StringField				string			`mapstructure:"STRING_FIELD"`
//		IntField				int				`mapstructure:"INT_FIELD"`
//		Int8Field				int8			`mapstructure:"INT8_FIELD"`
//		Int16Field				int16			`mapstructure:"INT16_FIELD"`
//		Int32Field				int32			`mapstructure:"INT32_FIELD"`
//		Int64Field				int64			`mapstructure:"INT64_FIELD"`
//		UintField				uint			`mapstructure:"UINT_FIELD"`
//		Uint8Field				uint8			`mapstructure:"UINT8_FIELD"`
//		Uint16Field				uint16			`mapstructure:"UINT16_FIELD"`
//		Uint32Field				uint32			`mapstructure:"UINT32_FIELD"`
//		Uint64Field				uint64			`mapstructure:"UINT64_FIELD"`
//		Float32Field			float32			`mapstructure:"FLOAT32_FIELD"`
//		Float64Field			float64			`mapstructure:"FLOAT64_FIELD"`
//		TimeDurationField 		time.Duration 	`mapstructure:"TIME_DURATION_FIELD"`
//		EmptyField				string			`mapstructure:"EMPTY_FIELD,omitempty"`
//	}
//
//	var cfg Config
//	err := UnmarshalFromEnv(&cfg)
func UnmarshalFromEnv(cfg any, withDefaults bool) error {
	val := reflect.ValueOf(cfg).Elem()

	for i := 0; i < val.NumField(); i++ {
		var (
			field        = val.Field(i)
			tag          = val.Type().Field(i).Tag.Get("mapstructure")
			key          = strings.Split(tag, ",")[0]
			mustBeSet    = !strings.Contains(tag, "omitempty")
			envValue     = os.Getenv(key)
			defaultValue = val.Type().Field(i).Tag.Get("default")
		)

		if field.Kind() == reflect.Struct {
			if err := UnmarshalFromEnv(field.Addr().Interface(), withDefaults); err != nil {
				return fmt.Errorf("failed to parse %s: %w", val.Type().Field(i).Name, err)
			}
			continue
		}

		if mustBeSet && envValue == "" && (defaultValue == "" || !withDefaults) {
			return fmt.Errorf("%s cannot be empty", key)
		}

		parser, ok := defaultTypeParsers[field.Type()]
		if !ok {
			parser, ok = defaultKindParsers[field.Kind()]
			if !ok {
				return fmt.Errorf("unsupported field type for %s", tag)
			}
		}

		if envValue == "" && withDefaults {
			log.Warnf("value for %s not found, using default value: %s", key, defaultValue)
			envValue = defaultValue
		}

		val, err := parser(envValue)
		if err != nil {
			return fmt.Errorf("failed to parse %s: %w", key, err)
		}

		field.Set(reflect.ValueOf(val).Convert(field.Type()))
	}

	return nil
}

func UnmarshalFromEnvFile(cfg any, envFilePath string, withDefaults bool) error {
	if err := godotenv.Overload(envFilePath); err != nil {
		return err
	}

	return UnmarshalFromEnv(cfg, withDefaults)
}
