package chronos

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDayEnd(t *testing.T) {
	h, m, s, ns := dayEnd()
	assert.Equal(t, 23, h)
	assert.Equal(t, 59, m)
	assert.Equal(t, 59, s)
	assert.Equal(t, 999999999, ns)
}

func TestStartOfHour(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			"normal case",
			time.Date(2023, 5, 15, 14, 30, 15, 123456789, time.UTC),
			time.Date(2023, 5, 15, 14, 0, 0, 0, time.UTC),
		},
		{
			"start of hour",
			time.Date(2023, 5, 15, 14, 0, 0, 0, time.UTC),
			time.Date(2023, 5, 15, 14, 0, 0, 0, time.UTC),
		},
		{
			"end of hour",
			time.Date(2023, 5, 15, 14, 59, 59, 999999999, time.UTC),
			time.Date(2023, 5, 15, 14, 0, 0, 0, time.UTC),
		},
		{
			"timezone test",
			time.Date(2023, 5, 15, 14, 30, 15, 123456789, time.FixedZone("CST", 8*3600)),
			time.Date(2023, 5, 15, 14, 0, 0, 0, time.FixedZone("CST", 8*3600)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StartOfHour(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestEndOfHour(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			"normal case",
			time.Date(2023, 5, 15, 14, 30, 15, 123456789, time.UTC),
			time.Date(2023, 5, 15, 14, 59, 59, 999999999, time.UTC),
		},
		{
			"start of hour",
			time.Date(2023, 5, 15, 14, 0, 0, 0, time.UTC),
			time.Date(2023, 5, 15, 14, 59, 59, 999999999, time.UTC),
		},
		{
			"end of hour",
			time.Date(2023, 5, 15, 14, 59, 59, 999999999, time.UTC),
			time.Date(2023, 5, 15, 14, 59, 59, 999999999, time.UTC),
		},
		{
			"timezone test",
			time.Date(2023, 5, 15, 14, 30, 15, 123456789, time.FixedZone("CST", 8*3600)),
			time.Date(2023, 5, 15, 14, 59, 59, 999999999, time.FixedZone("CST", 8*3600)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EndOfHour(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestStartOfDay(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			"normal case",
			time.Date(2023, 5, 15, 14, 30, 15, 123456789, time.UTC),
			time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			"start of day",
			time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			"end of day",
			time.Date(2023, 5, 15, 23, 59, 59, 999999999, time.UTC),
			time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			"timezone test",
			time.Date(2023, 5, 15, 14, 30, 15, 123456789, time.FixedZone("CST", 8*3600)),
			time.Date(2023, 5, 15, 0, 0, 0, 0, time.FixedZone("CST", 8*3600)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StartOfDay(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestEndOfDay(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			"normal case",
			time.Date(2023, 5, 15, 14, 30, 15, 123456789, time.UTC),
			time.Date(2023, 5, 15, 23, 59, 59, 999999999, time.UTC),
		},
		{
			"start of day",
			time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 5, 15, 23, 59, 59, 999999999, time.UTC),
		},
		{
			"end of day",
			time.Date(2023, 5, 15, 23, 59, 59, 999999999, time.UTC),
			time.Date(2023, 5, 15, 23, 59, 59, 999999999, time.UTC),
		},
		{
			"timezone test",
			time.Date(2023, 5, 15, 14, 30, 15, 123456789, time.FixedZone("CST", 8*3600)),
			time.Date(2023, 5, 15, 23, 59, 59, 999999999, time.FixedZone("CST", 8*3600)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EndOfDay(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestStartOfWeek(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		opts     []func(*WeekStartOption)
		expected time.Time
	}{
		{
			"Monday",
			time.Date(2023, 5, 15, 14, 30, 15, 123456789, time.UTC), // Monday
			nil,
			time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			"Tuesday",
			time.Date(2023, 5, 16, 14, 30, 15, 123456789, time.UTC), // Tuesday
			nil,
			time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			"Sunday",
			time.Date(2023, 5, 21, 14, 30, 15, 123456789, time.UTC), // Sunday
			nil,
			time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			"start of week (Monday)",
			time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC), // Monday
			nil,
			time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			"end of week (Sunday)",
			time.Date(2023, 5, 21, 23, 59, 59, 999999999, time.UTC), // Sunday
			nil,
			time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			"cross month",
			time.Date(2023, 5, 1, 14, 30, 15, 123456789, time.UTC), // Monday (April 24 is Sunday)
			nil,
			time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC),
		},

		{
			"Set Sunday as the beginning of the week",
			time.Date(2023, 5, 18, 15, 30, 0, 0, time.UTC),
			[]func(*WeekStartOption){WithWeekStartDay(time.Sunday)},
			time.Date(2023, 5, 14, 0, 0, 0, 0, time.UTC), // 上周日
		},
		{
			"Sunday day (when the starting date of the week is Sunday)",
			time.Date(2023, 5, 14, 10, 0, 0, 0, time.UTC), // 周日
			[]func(*WeekStartOption){WithWeekStartDay(time.Sunday)},
			time.Date(2023, 5, 14, 0, 0, 0, 0, time.UTC), // 同一天
		},
		{
			"cross month",
			time.Date(2023, 6, 1, 10, 0, 0, 0, time.UTC), // 周四（6月1日）
			[]func(*WeekStartOption){WithWeekStartDay(time.Sunday)},
			time.Date(2023, 5, 28, 0, 0, 0, 0, time.UTC), // 上周日（5月28日）
		},
		{
			"cross year",
			time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC), // 周一
			[]func(*WeekStartOption){WithWeekStartDay(time.Sunday)},
			time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC), // 上周日
		},
		{
			"influence of time zones",
			time.Date(2023, 5, 18, 15, 30, 0, 0, time.FixedZone("CST", 8*60*60)), // UTC+8
			nil,
			time.Date(2023, 5, 15, 0, 0, 0, 0, time.FixedZone("CST", 8*60*60)), // 周一（同时区）
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StartOfWeek(tt.input, tt.opts...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestEndOfWeek(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		opts     []func(*WeekStartOption)
		expected time.Time
	}{
		{
			"Monday",
			time.Date(2023, 5, 15, 14, 30, 15, 123456789, time.UTC), // Monday
			nil,
			time.Date(2023, 5, 21, 23, 59, 59, 999999999, time.UTC),
		},
		{
			"Tuesday",
			time.Date(2023, 5, 16, 14, 30, 15, 123456789, time.UTC), // Tuesday
			nil,
			time.Date(2023, 5, 21, 23, 59, 59, 999999999, time.UTC),
		},
		{
			"Sunday",
			time.Date(2023, 5, 21, 14, 30, 15, 123456789, time.UTC), // Sunday
			nil,
			time.Date(2023, 5, 21, 23, 59, 59, 999999999, time.UTC),
		},
		{
			"start of week (Monday)",
			time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC), // Monday
			nil,
			time.Date(2023, 5, 21, 23, 59, 59, 999999999, time.UTC),
		},
		{
			"end of week (Sunday)",
			time.Date(2023, 5, 21, 23, 59, 59, 999999999, time.UTC), // Sunday
			nil,
			time.Date(2023, 5, 21, 23, 59, 59, 999999999, time.UTC),
		},
		{
			"cross month",
			time.Date(2023, 4, 30, 14, 30, 15, 123456789, time.UTC), // Sunday (April 30)
			nil,
			time.Date(2023, 4, 30, 23, 59, 59, 999999999, time.UTC),
		},

		{
			"sunday as beginning (Saturday as end)",
			time.Date(2023, 5, 18, 15, 30, 0, 0, time.UTC),
			[]func(*WeekStartOption){WithWeekStartDay(time.Sunday)},
			time.Date(2023, 5, 20, 23, 59, 59, 999999999, time.UTC), // 周六
		},
		{
			"tuesday as beginning (Monday as end)",
			time.Date(2023, 5, 18, 15, 30, 0, 0, time.UTC),
			[]func(*WeekStartOption){WithWeekStartDay(time.Tuesday)},
			time.Date(2023, 5, 22, 23, 59, 59, 999999999, time.UTC), // 周一
		},
		{
			"sunday (default end date)",
			time.Date(2023, 5, 21, 10, 0, 0, 0, time.UTC), // 周日
			nil,
			time.Date(2023, 5, 21, 23, 59, 59, 999999999, time.UTC), // 同一天
		},
		{
			"saturday (saturday as end)",
			time.Date(2023, 5, 20, 10, 0, 0, 0, time.UTC), // 周六
			[]func(*WeekStartOption){WithWeekStartDay(time.Sunday)},
			time.Date(2023, 5, 20, 23, 59, 59, 999999999, time.UTC), // 同一天
		},
		{
			"cross month",
			time.Date(2023, 5, 31, 10, 0, 0, 0, time.UTC), // 周三（5月31日）
			nil,
			time.Date(2023, 6, 4, 23, 59, 59, 999999999, time.UTC), // 周日（6月4日）
		},
		{
			"cross year",
			time.Date(2023, 12, 30, 10, 0, 0, 0, time.UTC), // 周六
			[]func(*WeekStartOption){WithWeekStartDay(time.Monday)},
			time.Date(2023, 12, 31, 23, 59, 59, 999999999, time.UTC), // 周日
		},
		{
			"influence of time zones",
			time.Date(2023, 5, 18, 15, 30, 0, 0, time.FixedZone("CST", 8*60*60)), // UTC+8
			nil,
			time.Date(2023, 5, 21, 23, 59, 59, 999999999, time.FixedZone("CST", 8*60*60)), // 周日（同时区）
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EndOfWeek(tt.input, tt.opts...)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestStartOfMonth(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			"normal case",
			time.Date(2023, 5, 15, 14, 30, 15, 123456789, time.UTC),
			time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			"start of month",
			time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			"end of month",
			time.Date(2023, 5, 31, 23, 59, 59, 999999999, time.UTC),
			time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			"February non-leap year",
			time.Date(2023, 2, 15, 14, 30, 15, 123456789, time.UTC),
			time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			"February leap year",
			time.Date(2024, 2, 15, 14, 30, 15, 123456789, time.UTC),
			time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StartOfMonth(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestEndOfMonth(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			"normal case",
			time.Date(2023, 5, 15, 14, 30, 15, 123456789, time.UTC),
			time.Date(2023, 5, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			"start of month",
			time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 5, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			"end of month",
			time.Date(2023, 5, 31, 23, 59, 59, 999999999, time.UTC),
			time.Date(2023, 5, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			"February non-leap year",
			time.Date(2023, 2, 15, 14, 30, 15, 123456789, time.UTC),
			time.Date(2023, 2, 28, 23, 59, 59, 999999999, time.UTC),
		},
		{
			"February leap year",
			time.Date(2024, 2, 15, 14, 30, 15, 123456789, time.UTC),
			time.Date(2024, 2, 29, 23, 59, 59, 999999999, time.UTC),
		},
		{
			"December",
			time.Date(2023, 12, 15, 14, 30, 15, 123456789, time.UTC),
			time.Date(2023, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EndOfMonth(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestStartOfQuarter(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			"Q1",
			time.Date(2023, 1, 15, 14, 30, 15, 123456789, time.UTC),
			time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			"Q2",
			time.Date(2023, 5, 15, 14, 30, 15, 123456789, time.UTC),
			time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			"Q3",
			time.Date(2023, 8, 15, 14, 30, 15, 123456789, time.UTC),
			time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			"Q4",
			time.Date(2023, 11, 15, 14, 30, 15, 123456789, time.UTC),
			time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			"start of quarter",
			time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			"end of quarter",
			time.Date(2023, 6, 30, 23, 59, 59, 999999999, time.UTC),
			time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StartOfQuarter(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestEndOfQuarter(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			"Q1",
			time.Date(2023, 1, 15, 14, 30, 15, 123456789, time.UTC),
			time.Date(2023, 3, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			"Q2",
			time.Date(2023, 5, 15, 14, 30, 15, 123456789, time.UTC),
			time.Date(2023, 6, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			"Q3",
			time.Date(2023, 8, 15, 14, 30, 15, 123456789, time.UTC),
			time.Date(2023, 9, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			"Q4",
			time.Date(2023, 11, 15, 14, 30, 15, 123456789, time.UTC),
			time.Date(2023, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			"start of quarter",
			time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 6, 30, 23, 59, 59, 999999999, time.UTC),
		},
		{
			"end of quarter",
			time.Date(2023, 6, 30, 23, 59, 59, 999999999, time.UTC),
			time.Date(2023, 6, 30, 23, 59, 59, 999999999, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EndOfQuarter(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestStartOfYear(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			"normal case",
			time.Date(2023, 5, 15, 14, 30, 15, 123456789, time.UTC),
			time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			"start of year",
			time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			"end of year",
			time.Date(2023, 12, 31, 23, 59, 59, 999999999, time.UTC),
			time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			"leap year",
			time.Date(2024, 2, 29, 14, 30, 15, 123456789, time.UTC),
			time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StartOfYear(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestEndOfYear(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			"normal case",
			time.Date(2023, 5, 15, 14, 30, 15, 123456789, time.UTC),
			time.Date(2023, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			"start of year",
			time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			"end of year",
			time.Date(2023, 12, 31, 23, 59, 59, 999999999, time.UTC),
			time.Date(2023, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
		{
			"leap year",
			time.Date(2024, 2, 29, 14, 30, 15, 123456789, time.UTC),
			time.Date(2024, 12, 31, 23, 59, 59, 999999999, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EndOfYear(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
