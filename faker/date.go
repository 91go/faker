package faker

import (
	"github.com/91go/faker/utils"
	"strconv"
	"time"
)

// 当前时间戳
func NowTimeStamp() string {
	t := time.Now()
	return strconv.FormatInt(t.Unix(), 10)
}

func NowDate() string {
	dateStr := time.Now().Format("20060102,15:04:05")
	return dateStr
}

// Date will generate a random time.Time struct
func Date() time.Time {
	return time.Date(Year(), time.Month(utils.Number(0, 12)), Day(), Hour(), Minute(), Second(), NanoSecond(), time.UTC)
}

// DateRange will generate a random time.Time struct between a start and end date
func DateRange(start, end time.Time) time.Time {
	return time.Unix(0, int64(utils.Number(int(start.UnixNano()), int(end.UnixNano())))).UTC()
}

// Month will generate a random month string
func Month() string {
	return time.Month(utils.Number(1, 12)).String()
}

// Day will generate a random day between 1 - 31
func Day() int {
	return utils.Number(1, 31)
}

// WeekDay will generate a random weekday string (Monday-Sunday)
func WeekDay() string {
	return time.Weekday(utils.Number(0, 6)).String()
}

// Year will generate a random year between 1900 - current year
func Year() int {
	return utils.Number(1900, time.Now().Year())
}

// Hour will generate a random hour - in military time
func Hour() int {
	return utils.Number(0, 23)
}

// Minute will generate a random minute
func Minute() int {
	return utils.Number(0, 59)
}

// Second will generate a random second
func Second() int {
	return utils.Number(0, 59)
}

// NanoSecond will generate a random nano second
func NanoSecond() int {
	return utils.Number(0, 999999999)
}
