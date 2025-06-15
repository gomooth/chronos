package chronos

import (
	"github.com/gomooth/chronos/internal/helper"
)

// IsLeap 判断是否为闰年
func IsLeap[T MixedTime](v T) bool {
	at := getTime(v)
	return helper.IsLeap(at.Year())
}

// DaysInMonth 计算某年某月的天数
func DaysInMonth[T MixedTime](v T) int {
	at := getTime(v)
	return helper.DaysInMonth(at.Year(), at.Month())
}
