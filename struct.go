package utilities

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/bigquery"
	errortools "github.com/leapforce-libraries/go_errortools"
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

func StringArrayToStruct(records *[][]string, model interface{}) *errortools.Error {
	if records == nil {
		return nil
	}

	if reflect.TypeOf(model).Kind() != reflect.Ptr {
		return errortools.ErrorMessage("The interface is not a pointer.")
	}

	v := reflect.ValueOf(model).Elem()
	if v.Kind() != reflect.Slice {
		return errortools.ErrorMessage("The interface is not a pointer to a slice.")
	}

	rv := reflect.ValueOf(model)

	structType := reflect.TypeOf(model).Elem().Elem()

	numFields := structType.NumField()

	fields := make(map[string]int)

	for index, record := range *records {
		for j, v := range record {
			//remove inivisible characters and trim
			v = strings.ReplaceAll(v, string([]byte{byte(239)}), "")
			v = strings.ReplaceAll(v, string([]byte{byte(187)}), "")
			v = strings.ReplaceAll(v, string([]byte{byte(191)}), "")
			v = strings.Trim(v, " ")

			(*records)[index][j] = v
		}

		if index == 0 {
			for cellIndex, cellValue := range record {
				fields[strings.Trim(cellValue, " ")] = cellIndex
			}

			continue
		}

		new := reflect.New(structType).Elem()

		for i := 0; i < numFields; i++ {
			fieldName := structType.Field(i).Name
			fieldTag := structType.Field(i).Tag.Get("csv")

			if fieldTag == "" {
				continue
			}
			fieldIndex, ok := fields[fieldTag]

			if ok {
				value := strings.Trim(record[fieldIndex], " ")

				switch new.FieldByName(fieldName).Kind() {
				case reflect.String:
					new.FieldByName(fieldName).SetString(value)
					break
				case reflect.Int:
					i, err := strconv.ParseInt(value, 10, 64)
					if err == nil {
						new.FieldByName(fieldName).SetInt(i)
					}
					break
				case reflect.Float64:
					i, err := strconv.ParseFloat(value, 64)
					if err == nil {
						new.FieldByName(fieldName).SetFloat(i)
					}
					break
				}

			}
		}

		rv.Elem().Set(reflect.Append(rv.Elem(), new))
	}

	return nil
}

func StructToStringArray(model interface{}, includeHeaders bool) (*[][]string, *errortools.Error) {
	if reflect.TypeOf(model).Kind() != reflect.Ptr {
		return nil, errortools.ErrorMessage("The interface is not a pointer.")
	}

	v := reflect.ValueOf(model).Elem()
	if v.Kind() != reflect.Slice {
		return nil, errortools.ErrorMessage("The interface is not a pointer to a slice.")
	}

	structType := reflect.TypeOf(model).Elem().Elem()

	records := [][]string{}

	if includeHeaders {
		record := []string{}
		for i := 0; i < structType.NumField(); i++ {
			fieldName := structType.Field(i).Tag.Get("csv")
			if fieldName == "" {
				fieldName = structType.Field(i).Name
			}
			record = append(record, fieldName)
		}

		records = append(records, record)
	}

	for i := 0; i < v.Len(); i++ {

		record := []string{}
		v1 := v.Index(i)
		for j := 0; j < v1.NumField(); j++ {
			switch v1.Field(j).Kind() {
			case reflect.String:
				record = append(record, v1.Field(j).String())
				break
			case reflect.Int:
				record = append(record, strconv.FormatInt(v1.Field(j).Int(), 10))
				break
			case reflect.Float64:
				record = append(record, strconv.FormatFloat(v1.Field(j).Float(), 'f', 5, 64))
				break
			default:
				record = append(record, "")
				break
			}
		}

		records = append(records, record)
	}

	return &records, nil
}

func StructToURL(model interface{}, tag *string) (*string, *errortools.Error) {
	if IsNil(model) {
		return nil, nil
	}

	if reflect.TypeOf(model).Kind() != reflect.Ptr {
		return nil, errortools.ErrorMessage("The interface is not a pointer.")
	}

	p := reflect.ValueOf(model) //pointer
	s := p.Elem()               //interface

	if s.Kind() != reflect.Struct {
		s = s.Elem()
	}

	if s.Kind() != reflect.Struct {
		return nil, errortools.ErrorMessage("The interface is not a pointer to a struct.")
	}

	values := url.Values{}

	for j := 0; j < s.NumField(); j++ {
		fieldName := s.Type().Field(j).Name

		if tag != nil {
			tagValue := s.Type().Field(j).Tag.Get(*tag)

			if tagValue == "" {
				continue
			}
			fieldName = tagValue
		}

		field := s.Field(j)

		if field.Kind() == reflect.Ptr {
			if field.IsNil() {
				continue
			}

			field = field.Elem()
		}

		switch field.Kind() {
		case reflect.String:
			values.Set(fieldName, field.String())
			break
		case reflect.Int:
			values.Set(fieldName, strconv.FormatInt(field.Int(), 10))
			break
		case reflect.Float64:
			values.Set(fieldName, strconv.FormatFloat(field.Float(), 'f', 5, 64))
			break
		default:
			break
		}
	}

	url := values.Encode()

	return &url, nil
}

func SetStructField(model interface{}, fieldName string, value interface{}) *errortools.Error {
	if reflect.TypeOf(model).Kind() != reflect.Ptr {
		return errortools.ErrorMessage("Model is not a pointer.")
	}

	val := reflect.ValueOf(model)
	s := val.Elem()

	if s.Kind() != reflect.Struct {
		return errortools.ErrorMessage("Model is not a pointer to a struct.")
	}

	f := s.FieldByNameFunc(func(name string) bool {
		return strings.ToLower(name) == strings.ToLower(fieldName)
	})

	if f.IsValid() {
		if f.CanSet() {
			f.Set(reflect.ValueOf(value))
		}
	}

	return nil
}

func GetStructFieldString(model interface{}, fieldName string) string {

	f := reflect.ValueOf(model).Elem().FieldByName(fieldName)
	if f.IsZero() {
		return ""
	}

	fieldValue := f.Interface()
	value := ""
	switch v := fieldValue.(type) {

	case bigquery.NullFloat64:
		if v.Valid {
			value = strconv.FormatFloat(v.Float64, 'f', -1, 64)
		} else {
			value = ""
		}
	case bigquery.NullInt64:
		if v.Valid {
			value = strconv.FormatInt(v.Int64, 10)
		} else {
			value = ""
		}
	case int64:
		value = strconv.FormatInt(v, 10)
	case int32:
		value = strconv.FormatInt(int64(v), 10)
	case int:
		value = strconv.FormatInt(int64(v), 10)
	case string:
		value = v
	case bool:
		if v {
			value = "TRUE"
		}
		value = "FALSE"
	case bigquery.NullTimestamp:
		if v.Valid {
			value = v.Timestamp.Format("02-01-2006")
		} else {
			value = ""
		}
	case bigquery.NullDate:
		if v.Valid {
			if v.Date.Day == 1 && v.Date.Month == 1 && v.Date.Year == 1800 {
				value = ""
			} else {
				value = fmt.Sprintf("%02d-%02d-%04d", v.Date.Day, v.Date.Month, v.Date.Year)
			}
		} else {
			value = ""
		}
	case time.Time:
		value = v.Format("02-01-2006")
	case bigquery.NullString:
		value = ""
		if v.Valid {
			value = v.StringVal
		}
	default:
		value = ""
	}

	return value
}
