package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type IconType uint

const (
	FrontFace IconType = iota
)

type IconsUrl map[IconType]string

func (w *IconsUrl) Scan(value interface{}) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	err := json.Unmarshal(bytes, w)
	return err
}

func (w *IconsUrl) Value() (driver.Value, error) {
	bytes, err := json.Marshal(w)
	return string(bytes), err
}
