package utilities

import (
	"strconv"
	"time"
)

type Dictionary map[string]string

func (dictionary *Dictionary) get(key string) *string {
	if dictionary == nil {
		return nil
	}

	s, ok := map[string]string(*dictionary)[key]
	if ok {
		return &s
	}

	return nil
}

func (dictionary *Dictionary) GetString(key string) *string {
	return dictionary.get(key)
}

func (dictionary *Dictionary) GetInt(key string) *int {
	s := dictionary.GetString(key)

	if s == nil {
		return nil
	}

	i64, err := strconv.ParseInt(*s, 10, 0)
	if err != nil {
		return nil
	}

	i := int(i64)

	return &i
}

func (dictionary *Dictionary) GetInt64(key string) *int64 {
	s := dictionary.GetString(key)

	if s == nil {
		return nil
	}

	i64, err := strconv.ParseInt(*s, 10, 64)
	if err != nil {
		return nil
	}

	return &i64
}

func (dictionary *Dictionary) GetFloat64(key string) *float64 {
	s := dictionary.GetString(key)

	if s == nil {
		return nil
	}

	f64, err := strconv.ParseFloat(*s, 64)
	if err != nil {
		return nil
	}

	return &f64
}

func (dictionary *Dictionary) GetBool(key string) *bool {
	s := dictionary.GetString(key)

	if s == nil {
		return nil
	}

	b, err := strconv.ParseBool(*s)
	if err != nil {
		return nil
	}

	return &b
}

func (dictionary *Dictionary) GetTime(key string, layout string) *time.Time {
	s := dictionary.GetString(key)

	if s == nil {
		return nil
	}

	time, err := time.Parse(layout, *s)
	if err != nil {
		return nil
	}

	return &time
}
