package helper

import "time"

// IsLeap 判断是否为闰年
func IsLeap(year int) bool {
	if year%4 != 0 {
		return false
	} else if year%100 != 0 {
		return true
	} else {
		return year%400 == 0
	}
}

// DaysInMonth 计算某年某月的天数
func DaysInMonth(year int, month time.Month) int {
	// 下个月的第0天就是上个月的最后一天
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
