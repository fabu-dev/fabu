package utils

import (
	"strconv"
	"strings"
	"time"
)

// 获取当前的时间字符串
func GetCurrentDateTime() string {
	nowTime := time.Now()
	localZone, _ := time.LoadLocation("Timezone")
	nowTime = nowTime.In(localZone)
	dateTime := nowTime.Format(TimeStdFormat)

	return dateTime
}

//获取每月的几号
func GetMonthDay() string {
	localZone, _ := time.LoadLocation(Timezone)
	dayInt := time.Now().In(localZone).Day()
	return strconv.Itoa(dayInt)
}

/*
获取星期几
注意：如果是周日，系统返回的是0，mms系统的规则是：1,2,3,4,5,6,7
即把0重新赋为7
*/
func GetWeekday() string {
	localZone, _ := time.LoadLocation(Timezone)
	buf := int(time.Now().In(localZone).Weekday())
	wday := strconv.Itoa(buf)
	if wday == "0" {
		wday = "7"
	}

	return wday
}

//获取本周周几的日期
func GetDateOfWeek(now time.Time, weekDayStr string, hour, minute, second int) (weekStartDate time.Time) {
	currentWeek, _ := strconv.Atoi(weekDayStr)
	offset := currentWeek - int(now.Weekday())
	weekStartDate = time.Date(now.Year(), now.Month(), now.Day(), hour, minute, second, 0, time.Local).AddDate(0, 0, offset)

	return weekStartDate
}

func GetTimeByString(timeStr string) []int {
	timeSlice := strings.Split(timeStr, ":")
	timeSliceInt := make([]int, 0, 3)
	for i := range timeSlice {
		val, _ := strconv.Atoi(timeSlice[i])
		timeSliceInt = append(timeSliceInt, val)
	}
	return timeSliceInt
}
