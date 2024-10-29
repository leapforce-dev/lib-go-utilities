package utilities

import (
	"cloud.google.com/go/civil"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
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

func Int64ArrayNotNil(value *[]int64) []int64 {
	if value == nil {
		return []int64{}
	}

	return *value
}

func Float32NotNil(value *float32) float32 {
	if value == nil {
		return float32(0)
	}

	return *value
}

func Float64NotNil(value *float64) float64 {
	if value == nil {
		return float64(0)
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

func TimeToTime(t civil.Time) time.Time {
	t_, _ := time.Parse("15:04:05", t.String())

	return t_
}

func MonthStartDate(date civil.Date) civil.Date {
	t, _ := time.Parse("2006-01-02", fmt.Sprintf("%04d-%02d-01", date.Year, date.Month))

	return civil.DateOf(t)
}

func MonthEndDate(date civil.Date) civil.Date {
	year := date.Year
	month := date.Month + 1
	if month == 13 {
		year++
		month = 1
	}
	t, _ := time.Parse("2006-01-02", fmt.Sprintf("%04d-%02d-01", year, month))

	return civil.DateOf(t).AddDays(-1)
}

func ParseFloat(str string) (float64, error) {
	val, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return val, nil
	}

	//Some number may be seperated by comma, for example, 23,120,123, so remove the comma firstly
	str = strings.Replace(str, ",", "", -1)

	//Some number is specifed in scientific notation
	pos := strings.IndexAny(str, "eE")
	if pos < 0 {
		return strconv.ParseFloat(str, 64)
	}

	var baseVal float64
	var expVal int64

	baseStr := str[0:pos]
	baseVal, err = strconv.ParseFloat(baseStr, 64)
	if err != nil {
		return 0, err
	}

	expStr := str[(pos + 1):]
	expVal, err = strconv.ParseInt(expStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return baseVal * math.Pow10(int(expVal)), nil
}
