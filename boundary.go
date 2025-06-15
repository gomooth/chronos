package chronos

import (
	"time"
)

// dayEnd 返回一天结束时的时分秒纳秒
func dayEnd() (hour, minute, sec, nsec int) {
	return 23, 59, 59, 1e9 - 1
}

// StartOfHour 返回当前小时的开始时间
func StartOfHour[T MixedTime](v T) time.Time {
	at := getTime(v)
	year, month, day := at.Date()
	return time.Date(year, month, day, at.Hour(), 0, 0, 0, at.Location())
}

// EndOfHour 返回当前小时的结束时间
func EndOfHour[T MixedTime](v T) time.Time {
	at := getTime(v)
	year, month, day := at.Date()
	_, minute, sec, nsec := dayEnd()
	return time.Date(year, month, day, at.Hour(), minute, sec, nsec, at.Location())
}

// StartOfDay 返回当天的开始时间
func StartOfDay[T MixedTime](v T) time.Time {
	at := getTime(v)
	year, month, day := at.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, at.Location())
}

// EndOfDay 返回当天的结束时间
func EndOfDay[T MixedTime](v T) time.Time {
	at := getTime(v)
	year, month, day := at.Date()
	hour, minute, sec, nsec := dayEnd()
	return time.Date(year, month, day, hour, minute, sec, nsec, at.Location())
}

// StartOfWeek 返回当周的开始时间
// 默认周一为一周的开始，可通过 WithWeekStartDay 选项修改
func StartOfWeek[T MixedTime](v T, opts ...func(*WeekStartOption)) time.Time {
	// 默认配置：周一为一周的开始
	config := &WeekStartOption{
		startDay: time.Monday,
	}

	// 应用选项
	for _, opt := range opts {
		opt(config)
	}

	at := getTime(v)
	weekday := at.Weekday()

	// 计算到周起始日的偏移量
	daysToSubtract := (weekday - config.startDay) % 7
	if daysToSubtract < 0 {
		daysToSubtract += 7
	}

	start := at.AddDate(0, 0, -int(daysToSubtract))
	year, month, day := start.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, at.Location())
}

// WeekStartOption 定义周起始日选项
type WeekStartOption struct {
	startDay time.Weekday
}

// WithWeekStartDay 设置周起始日
func WithWeekStartDay(day time.Weekday) func(*WeekStartOption) {
	return func(o *WeekStartOption) {
		o.startDay = day
	}
}

// EndOfWeek 返回当周的结束时间
// 默认周日为一周的结束，可通过 WithWeekStartDay 选项修改（结束日 = 起始日+6天）
func EndOfWeek[T MixedTime](v T, opts ...func(*WeekStartOption)) time.Time {
	// 默认配置：周一为一周的开始（周日为结束）
	config := &WeekStartOption{
		startDay: time.Monday,
	}

	// 应用选项
	for _, opt := range opts {
		opt(config)
	}

	at := getTime(v)
	weekday := at.Weekday()

	// 计算结束日（起始日+6天）
	endDay := (config.startDay + 6) % 7

	// 计算到周结束日的偏移量
	daysToAdd := (endDay - weekday) % 7
	if daysToAdd < 0 {
		daysToAdd += 7
	}

	end := at.AddDate(0, 0, int(daysToAdd))
	year, month, day := end.Date()
	hour, minute, sec, nsec := dayEnd()
	return time.Date(year, month, day, hour, minute, sec, nsec, at.Location())
}

// StartOfMonth 返回当月的开始时间
func StartOfMonth[T MixedTime](v T) time.Time {
	at := getTime(v)
	year, month, _ := at.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, at.Location())
}

// EndOfMonth 返回当月的结束时间
func EndOfMonth[T MixedTime](v T) time.Time {
	at := getTime(v)
	year, month, _ := at.Date()

	// 获取下个月的第一天，然后减去1纳秒
	nextMonth := month + 1
	if nextMonth > 12 {
		nextMonth = 1
		year++
	}

	return time.Date(year, nextMonth, 1, 0, 0, 0, 0, at.Location()).Add(-time.Nanosecond)
}

// StartOfQuarter 返回当前季度的开始时间
func StartOfQuarter[T MixedTime](v T) time.Time {
	at := getTime(v)
	year, month, _ := at.Date()
	quarter := (month-1)/3 + 1
	startMonth := (quarter-1)*3 + 1

	return time.Date(year, startMonth, 1, 0, 0, 0, 0, at.Location())
}

// EndOfQuarter 返回当前季度的结束时间
func EndOfQuarter[T MixedTime](v T) time.Time {
	at := getTime(v)
	year, month, _ := at.Date()
	quarter := (month-1)/3 + 1
	endMonth := time.Month(quarter * 3)

	// 获取下个季度的第一天，然后减去1纳秒
	nextQuarterMonth := endMonth + 1
	if nextQuarterMonth > 12 {
		nextQuarterMonth = 1
		year++
	}

	return time.Date(year, nextQuarterMonth, 1, 0, 0, 0, 0, at.Location()).Add(-time.Nanosecond)
}

// StartOfYear 返回当年的开始时间
func StartOfYear[T MixedTime](v T) time.Time {
	at := getTime(v)
	return time.Date(at.Year(), 1, 1, 0, 0, 0, 0, at.Location())
}

// EndOfYear 返回当年的结束时间
func EndOfYear[T MixedTime](v T) time.Time {
	at := getTime(v)
	// 获取下一年的第一天，然后减去1纳秒
	year := at.Year() + 1
	return time.Date(year, 1, 1, 0, 0, 0, 0, at.Location()).Add(-time.Nanosecond)
}
