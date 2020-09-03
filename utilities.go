package utilities

import (
	"reflect"
	"strings"
)

// GetJsonTaggedFieldNames returns comma separated string of
// fieldnames of struct having a specified tag
//
func GetJsonTaggedFieldNames(tag string, model interface{}) string {
	val := reflect.ValueOf(model)
	list := ""
	for i := 0; i < val.Type().NumField(); i++ {
		field := val.Type().Field(i)
		tag := field.Tag.Get(tag)
		if tag != "" {
			list += "," + field.Name
		}
	}

	list = strings.Trim(list, ",")

	return list
}
