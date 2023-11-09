package handlers

import (
	"fmt"
	"ga/internal/types"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

func HasAllFields(fieldsStruct interface{}) gin.H {
	v := reflect.ValueOf(fieldsStruct)

	if v.Kind() != reflect.Struct {
		return BuildError("argument must be a struct", nil)
	}

	var emptyFields []string

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		jsonName := strings.Split(v.Type().Field(i).Tag.Get("json"), ",")[0]
		tag := v.Type().Field(i).Tag.Get("ga")

		if strings.Contains(tag, "required") {
			if field.Kind() == reflect.Ptr && field.IsNil() {
				emptyFields = append(emptyFields, jsonName)
			}
		}
		// TODO: not null
	}

	if len(emptyFields) > 0 {
		return BuildError("empty required fields", emptyFields)
	}

	return nil
}

func HasAnyFields(fieldsStruct interface{}) gin.H {
	v := reflect.ValueOf(fieldsStruct)

	if v.Kind() != reflect.Struct {
		return BuildError("argument must be a struct", nil)
	}

	var emptyFields []string
	var requiredFields int

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		jsonName := strings.Split(v.Type().Field(i).Tag.Get("json"), ",")[0]
		tag := v.Type().Field(i).Tag.Get("ga")

		if strings.Contains(tag, "required") {
			fmt.Println(tag)
			requiredFields++
			if field.IsZero() || (field.Kind() == reflect.Ptr && field.IsNil()) {
				emptyFields = append(emptyFields, jsonName)
			} else {
				return nil
			}
		}
	}

	if len(emptyFields) == requiredFields && requiredFields > 0 {
		return BuildError("at least one field is required", emptyFields)
	}

	return nil
}

func HasLocalizedDefault(fieldsStruct any, defaultLanguage types.Language) gin.H {
	v := reflect.ValueOf(fieldsStruct)

	if v.Kind() != reflect.Struct {
		return BuildError("argument must be a struct", nil)
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
		return BuildError("default language is not provided for fields", missedLanguages)
	}

	return nil
}

var Languages = map[types.Language]string{
	types.Russian: "Russian",
	types.English: "English",
}

func GetLanguage(c *gin.Context) types.Language {
	languageCodes := strings.Split(strings.ToLower(c.GetHeader("Accept-Languages")), ",")

	for _, lCode := range languageCodes {
		if _, ok := Languages[types.Language(lCode)]; ok {
			return types.Language(lCode)
		}
	}

	return types.DefaultLanguage
}
