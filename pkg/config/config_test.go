package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalFromEnv(t *testing.T) {
	type TestConfig struct {
		BoolField bool `mapstructure:"BOOL_FIELD"`

		StringField string `mapstructure:"STRING_FIELD"`

		IntField   int   `mapstructure:"INT_FIELD"`
		Int8Field  int8  `mapstructure:"INT8_FIELD"`
		Int16Field int16 `mapstructure:"INT16_FIELD"`
		Int32Field int32 `mapstructure:"INT32_FIELD"`
		Int64Field int64 `mapstructure:"INT64_FIELD"`

		UintField   uint   `mapstructure:"UINT_FIELD"`
		Uint8Field  uint8  `mapstructure:"UINT8_FIELD"`
		Uint16Field uint16 `mapstructure:"UINT16_FIELD"`
		Uint32Field uint32 `mapstructure:"UINT32_FIELD"`
		Uint64Field uint64 `mapstructure:"UINT64_FIELD"`

		Float32Field float32 `mapstructure:"FLOAT32_FIELD"`
		Float64Field float64 `mapstructure:"FLOAT64_FIELD"`

		EmptyField string `mapstructure:"EMPTY_FIELD,omitempty"`
	}

	os.Setenv("BOOL_FIELD", "true")

	os.Setenv("STRING_FIELD", "test")

	os.Setenv("INT_FIELD", "-2147483648")
	os.Setenv("INT8_FIELD", "-128")
	os.Setenv("INT16_FIELD", "-32768")
	os.Setenv("INT32_FIELD", "-2147483648")
	os.Setenv("INT64_FIELD", "-9223372036854775808")

	os.Setenv("UINT_FIELD", "4294967295")
	os.Setenv("UINT8_FIELD", "255")
	os.Setenv("UINT16_FIELD", "65535")
	os.Setenv("UINT32_FIELD", "4294967295")
	os.Setenv("UINT64_FIELD", "18446744073709551615")

	os.Setenv("FLOAT32_FIELD", "3.14")
	os.Setenv("FLOAT64_FIELD", "3.14159265359")

	cfg := new(TestConfig)
	err := UnmarshalFromEnv(cfg, false)

	assert.NoError(t, err)
	assert.Equal(t, true, cfg.BoolField)
	assert.Equal(t, "test", cfg.StringField)
	assert.Equal(t, -2147483648, cfg.IntField)
	assert.Equal(t, int8(-128), cfg.Int8Field)
	assert.Equal(t, int16(-32768), cfg.Int16Field)
	assert.Equal(t, int32(-2147483648), cfg.Int32Field)
	assert.Equal(t, int64(-9223372036854775808), cfg.Int64Field)
	assert.Equal(t, uint(4294967295), cfg.UintField)
	assert.Equal(t, uint8(255), cfg.Uint8Field)
	assert.Equal(t, uint16(65535), cfg.Uint16Field)
	assert.Equal(t, uint32(4294967295), cfg.Uint32Field)
	assert.Equal(t, uint64(18446744073709551615), cfg.Uint64Field)
	assert.Equal(t, float32(3.14), cfg.Float32Field)
	assert.Equal(t, float64(3.14159265359), cfg.Float64Field)
}

func TestUnmarshalFromEnv_EmptyField(t *testing.T) {
	type TestConfig struct {
		BoolField   bool   `mapstructure:"BOOL_FIELD"`
		StringField string `mapstructure:"STRING_FIELD"`
	}

	os.Setenv("BOOL_FIELD", "true")
	os.Setenv("STRING_FIELD", "")

	cfg := new(TestConfig)
	err := UnmarshalFromEnv(cfg, false)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "STRING_FIELD cannot be empty")
}

func TestUnmarshalFromEnv_InvalidType(t *testing.T) {
	type TestConfig struct {
		BoolField bool `mapstructure:"BOOL_FIELD"`
		IntField  int  `mapstructure:"INT_FIELD"`
	}

	os.Setenv("BOOL_FIELD", "true")
	os.Setenv("INT_FIELD", "invalid")

	cfg := new(TestConfig)
	err := UnmarshalFromEnv(cfg, false)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse INT_FIELD")
}

func TestUnmarshalFromEnv_StructField(t *testing.T) {
	type TestConfig struct {
		StructField struct {
			NestedField int `mapstructure:"NESTED_FIELD"`
		}
	}

	os.Setenv("NESTED_FIELD", "123")

	cfg := new(TestConfig)
	err := UnmarshalFromEnv(cfg, false)

	assert.NoError(t, err)
	assert.Equal(t, 123, cfg.StructField.NestedField)
}

func TestUnmarshalFromEnv_OmitEmpty(t *testing.T) {
	type TestConfig struct {
		StringField string `mapstructure:"STRING_FIELD,omitempty"`
		BoolField   bool   `mapstructure:"BOOL_FIELD"`
		IntField    int    `mapstructure:"INT_FIELD"`
	}

	os.Setenv("BOOL_FIELD", "false")
	os.Setenv("STRING_FIELD", "")
	os.Setenv("INT_FIELD", "0")

	cfg := new(TestConfig)
	err := UnmarshalFromEnv(cfg, false)

	assert.NoError(t, err)
}
