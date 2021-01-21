package utilities

import (
	"time"

	"cloud.google.com/go/civil"
)

func StringNotNil(value *string) string {
	if value == nil {
		return ""
	}

	return *value
}

func IntNotNil(value *int) int {
	if value == nil {
		return 0
	}

	return *value
}

func Int32NotNil(value *int32) int32 {
	if value == nil {
		return int32(0)
	}

	return *value
}

func Int64NotNil(value *int64) int64 {
	if value == nil {
		return int64(0)
	}

	return *value
}

func BoolNotNil(value *bool) bool {
	if value == nil {
		return false
	}

	return *value
}

func DateToTime(date civil.Date) time.Time {
	t, _ := time.Parse("2006-01-02", date.String())

	return t
}
