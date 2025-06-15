package calculator

import (
	"time"

	"github.com/gomooth/chronos/internal/helper"
)

func AddDate(at time.Time, year, month, day int, opts ...func(*AddDateOption)) time.Time {
	cnf := new(AddDateOption)
	for _, opt := range opts {
		opt(cnf)
	}

	if !cnf.notOverflow {
		return at.AddDate(year, month, day)
	}

	// 获取原始时间部分
	originalYear, originalMonth, originalDay := at.Date()

	// 计算新的年月
	newYear := originalYear + year
	newMonth := int(originalMonth) + month
	newDay := originalDay + day

	// 处理月份溢出（可能跨年）
	for newMonth > 12 {
		newYear++
		newMonth -= 12
	}
	for newMonth < 1 {
		newYear--
		newMonth += 12
	}

	// 如果只是增加月份或年份（day=0）
	if day == 0 &&
		// 特殊日期处理
		(originalDay == 31 ||
			(originalMonth == time.January && originalDay >= 29) ||
			(originalMonth == time.February && originalDay == 29)) {

		// 计算新月份的天数
		daysInNewMonth := helper.DaysInMonth(newYear, time.Month(newMonth))

		// 确保新日期不超过新月份的天数
		newDay = originalDay
		if newDay > daysInNewMonth {
			newDay = daysInNewMonth
		}
	}

	return time.Date(
		newYear, time.Month(newMonth), newDay,
		at.Hour(), at.Minute(), at.Second(), at.Nanosecond(), at.Location())
}
