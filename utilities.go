package utilities

import (
	"reflect"
	"strings"
)

// GetTaggedFieldNames returns comma separated string of
// fieldnames of struct having a specified tag
//
func GetTaggedFieldNames(tag string, model interface{}) string {
	return getTaggedNames(tag, model, "field")
}

// GetTaggedTagNames returns comma separated string of
// fieldnames of struct having a specified tag
//
func GetTaggedTagNames(tag string, model interface{}) string {
	return getTaggedNames(tag, model, "tag")
}

func getTaggedNames(tag string, model interface{}, fieldOrTag string) string {
	val := reflect.ValueOf(model)
	list := ""
	for i := 0; i < val.Type().NumField(); i++ {
		field := val.Type().Field(i)
		value, ok := field.Tag.Lookup(tag)
		if ok {
			if fieldOrTag == "field" {
				list += "," + field.Name
			} else if fieldOrTag == "tag" {
				list += "," + value
			}
		}
	}

	list = strings.Trim(list, ",")

	return list
}
