package chronos

import (
	"strconv"
	"time"
)

// DiffValue 差值，单位：纳秒
type DiffValue int64

// Nanoseconds 转成纳秒
func (d DiffValue) Nanoseconds() int64 {
	time.Duration(d).Minutes()
	return int64(d)
}

// Microseconds 转成微妙
func (d DiffValue) Microseconds() int64 {
	return int64(d) / 1e3
}

// Milliseconds 转成毫秒
func (d DiffValue) Milliseconds() int64 {
	return int64(d) / 1e6
}

// Seconds 转成秒
func (d DiffValue) Seconds() int64 {
	return int64(d) / 1e9
}

// Minutes 转成分
func (d DiffValue) Minutes() int64 {
	return int64(d) / (60 * 1e9)
}

// Hours 转成小时
func (d DiffValue) Hours() int64 {
	return int64(d) / (3600 * 1e9)
}

// Days 转成天
func (d DiffValue) Days() int64 {
	return int64(d) / (24 * 3600 * 1e9)
}

// Weeks 转成星期
func (d DiffValue) Weeks() int64 {
	return int64(d) / (7 * 24 * 3600 * 1e9)
}

// Months 转成月，默认每月30天
// 可使用 DiffWithDaysPer 设置每月天数
func (d DiffValue) Months(opts ...func(*int)) int64 {
	daysPerMonth := 30 // 平均值 30.44
	for _, opt := range opts {
		opt(&daysPerMonth)
	}
	return int64(d) / (int64(daysPerMonth) * 24 * 3600 * 1e9)
}

// DiffWithDaysPer 设置每月/每年天数
func DiffWithDaysPer(daysPer int) func(*int) {
	return func(value *int) {
		*value = daysPer
	}
}

// Years 转成年，默认每年365天
// 可使用 DiffWithDaysPer 设置每年天数
func (d DiffValue) Years(opts ...func(*int)) int64 {
	daysPerYear := 365 // 平均值 365.24
	for _, opt := range opts {
		opt(&daysPerYear)
	}
	return int64(d) / (int64(daysPerYear) * 24 * 3600 * 1e9)
}

// String 返回人类可读的时间差格式，无时间差则返回空字符串
func (d DiffValue) String() string {
	absNanos := int64(d)
	if absNanos < 0 {
		absNanos = -absNanos
	}
	sign := ""
	if d < 0 {
		sign = "-"
	}

	var formatTime = func(value int, unit string) string {
		if value == 0 {
			return ""
		}
		return strconv.Itoa(value) + unit
	}

	switch {
	case absNanos < 1e3:
		return sign + formatTime(int(absNanos), "ns")
	case absNanos < 1e6:
		v := int(absNanos / 1e3)
		return sign + formatTime(v, "μs")
	case absNanos < 1e9:
		ms := int(absNanos / 1e6)
		return sign + formatTime(ms, "ms")
	case absNanos < 60*1e9:
		sec := int(absNanos / 1e9)
		return sign + formatTime(sec, "s")
	case absNanos < 3600*1e9:
		val := int(absNanos / (60 * 1e9))
		sec := int((absNanos % (60 * 1e9)) / 1e9)
		return sign + formatTime(val, "m") + formatTime(sec, "s")
	default:
		hours := int(absNanos / (3600 * 1e9))
		remaining := absNanos % (3600 * 1e9)
		val := int(remaining / (60 * 1e9))
		sec := int(remaining % (60 * int64(1e9)) / 1e9)
		return sign + formatTime(hours, "h") + formatTime(val, "m") + formatTime(sec, "s")
	}
}
