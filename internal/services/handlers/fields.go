package handlers

import (
	"ga/pkg/genshin_core/models/languages"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

func HasAllFields(fieldsStruct interface{}) gin.H {
	v := reflect.ValueOf(fieldsStruct)

	if v.Kind() != reflect.Struct {
		return buildError("argument must be a struct", nil)
	}

	var emptyFields []string

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		jsonName := strings.Split(v.Type().Field(i).Tag.Get("json"), ",")[0]
		tag := v.Type().Field(i).Tag.Get("ga")

		if strings.Contains(tag, "required") {
			if field.IsZero() || (field.Kind() == reflect.Ptr && field.IsNil()) {
				emptyFields = append(emptyFields, jsonName)
			}
		}
	}

	if len(emptyFields) > 0 {
		return buildError("empty required fields", emptyFields)
	}

	return nil
}

func HasAnyFields(fieldsStruct interface{}) gin.H {
	v := reflect.ValueOf(fieldsStruct)

	if v.Kind() != reflect.Struct {
		return buildError("argument must be a struct", nil)
	}

	var emptyFields []string
	var requiredFields int

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		jsonName := strings.Split(v.Type().Field(i).Tag.Get("json"), ",")[0]
		tag := v.Type().Field(i).Tag.Get("ga")

		if strings.Contains(tag, "required") {
			requiredFields++
			if field.IsZero() || (field.Kind() == reflect.Ptr && field.IsNil()) {
				emptyFields = append(emptyFields, jsonName)
			} else {
				return nil
			}
		}
	}

	if len(emptyFields) == requiredFields {
		return buildError("at least one field is required", emptyFields)
	}

	return nil
}

func HasLocalizedDefault(fieldsStruct any, defaultLanguage languages.Language) gin.H {
	v := reflect.ValueOf(fieldsStruct)

	if v.Kind() != reflect.Struct {
		return buildError("argument must be a struct", nil)
	}

	var missedLanguages []string

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		jsonName := strings.Split(v.Type().Field(i).Tag.Get("json"), ",")[0]
		tag := v.Type().Field(i).Tag.Get("ga")

		if strings.Contains(tag, "localized") {
			if field.Kind() == reflect.Map && !field.MapIndex(reflect.ValueOf(defaultLanguage)).IsValid() {
				missedLanguages = append(missedLanguages, jsonName)
			}
		}
	}

	if len(missedLanguages) > 0 {
		return buildError("default language is not provided for fields", missedLanguages)
	}

	return nil
}
