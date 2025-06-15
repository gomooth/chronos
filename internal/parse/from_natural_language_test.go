package parse_test

import (
	"testing"
	"time"

	"github.com/gomooth/chronos/internal/parse"
)

func TestFromNaturalLanguage(t *testing.T) {
	// 固定基准时间，便于测试
	baseTime := time.Date(2023, 5, 15, 12, 0, 0, 0, time.UTC)
	locNY, _ := time.LoadLocation("America/New_York")
	locTokyo, _ := time.LoadLocation("Asia/Tokyo")

	tests := []struct {
		name     string
		expr     string
		opts     []func(*parse.FromNaturalLanguageOption)
		expected time.Time
		wantErr  bool
	}{
		// 基本功能测试
		{"a hour ago", "a hour ago", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, time.Date(2023, 5, 15, 11, 0, 0, 0, time.UTC), false},
		{"an hour ago", "an hour ago", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, time.Date(2023, 5, 15, 11, 0, 0, 0, time.UTC), false},
		{"2 hours later", "2 hours later", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, time.Date(2023, 5, 15, 14, 0, 0, 0, time.UTC), false},
		{"30 minutes ago", "30 minutes ago", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, time.Date(2023, 5, 15, 11, 30, 0, 0, time.UTC), false},
		{"5 seconds before", "5 seconds before", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, time.Date(2023, 5, 15, 11, 59, 55, 0, time.UTC), false},
		{"3 days after", "3 days after", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, time.Date(2023, 5, 18, 12, 0, 0, 0, time.UTC), false},
		{"1 week ago", "1 week ago", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, time.Date(2023, 5, 8, 12, 0, 0, 0, time.UTC), false},
		{"2 weeks later", "2 weeks later", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, time.Date(2023, 5, 29, 12, 0, 0, 0, time.UTC), false},

		{"1 nanosecond ago", "1 nanosecond ago", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, time.Date(2023, 5, 15, 11, 59, 59, 999999999, time.UTC), false},
		{"500 nanoseconds later", "500 nanoseconds later", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, time.Date(2023, 5, 15, 12, 0, 0, 500, time.UTC), false},
		{"a microsecond ago", "a microsecond ago", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, time.Date(2023, 5, 15, 11, 59, 59, 999999000, time.UTC), false},
		{"2 microseconds later", "2 microseconds later", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, time.Date(2023, 5, 15, 12, 0, 0, 2000, time.UTC), false},
		{"a millisecond ago", "a millisecond ago", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, time.Date(2023, 5, 15, 11, 59, 59, 999000000, time.UTC), false},
		{"5 milliseconds later", "5 milliseconds later", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, time.Date(2023, 5, 15, 12, 0, 0, 5000000, time.UTC), false},
		{"1 nanosecond before", "1 nanosecond before", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, time.Date(2023, 5, 15, 11, 59, 59, 999999999, time.UTC), false},
		{"1 microsecond after", "1 microsecond after", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, time.Date(2023, 5, 15, 12, 0, 0, 1000, time.UTC), false},

		// 月份和年份测试（边界情况）
		{"a month ago", "a month ago", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, time.Date(2023, 4, 15, 12, 0, 0, 0, time.UTC), false},
		{"12 months ago", "12 months ago", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, time.Date(2022, 5, 15, 12, 0, 0, 0, time.UTC), false},
		{"1 month later", "1 month later", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, time.Date(2023, 6, 15, 12, 0, 0, 0, time.UTC), false},
		{"a year ago", "a year ago", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, time.Date(2022, 5, 15, 12, 0, 0, 0, time.UTC), false},
		{"5 years later", "5 years later", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, time.Date(2028, 5, 15, 12, 0, 0, 0, time.UTC), false},

		// 特殊表达式测试
		{"now", "now", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, baseTime, false},
		{"today", "today", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, time.Date(2023, 5, 15, 0, 0, 0, 0, time.UTC), false},
		{"yesterday", "yesterday", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, time.Date(2023, 5, 14, 0, 0, 0, 0, time.UTC), false},
		{"tomorrow", "tomorrow", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, time.Date(2023, 5, 16, 0, 0, 0, 0, time.UTC), false},

		// 边界值测试
		{"0 hours ago", "0 hours ago", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, baseTime, false},
		{"1 nanosecond ago", "1 nanosecond ago", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, time.Date(2023, 5, 15, 11, 59, 59, 999999999, time.UTC), false},
		{"large value", "10000 years later", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime)}, time.Date(12023, 5, 15, 12, 0, 0, 0, time.UTC), false},
		{"leap year test", "1 year ago", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(time.Date(2024, 3, 1, 12, 0, 0, 0, time.UTC))}, time.Date(2023, 3, 1, 12, 0, 0, 0, time.UTC), false},

		// 时区测试
		{"New York timezone", "1 hour ago", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime), parse.WithFromNaturalLanguageLocation(locNY)}, time.Date(2023, 5, 15, 7, 0, 0, 0, locNY), false},
		{"Tokyo timezone", "2 hours later", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime), parse.WithFromNaturalLanguageLocation(locTokyo)}, time.Date(2023, 5, 15, 23, 0, 0, 0, locTokyo), false},
		{"Today in Tokyo", "today", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime), parse.WithFromNaturalLanguageLocation(locTokyo)}, time.Date(2023, 5, 15, 0, 0, 0, 0, locTokyo), false},
		{"Month calculation in different timezone", "1 month ago", []func(*parse.FromNaturalLanguageOption){parse.WithFromNaturalLanguageBaseTime(baseTime), parse.WithFromNaturalLanguageLocation(locNY)}, baseTime.AddDate(0, -1, 0).In(locNY), false},

		// 异常情况测试
		{"empty string", "", nil, time.Time{}, true},
		{"invalid format", "invalid time string", nil, time.Time{}, true},
		{"unknown unit", "1 century ago", nil, time.Time{}, true},
		{"invalid direction", "1 hour sideways", nil, time.Time{}, true},
		{"negative quantity", "-1 hour ago", nil, time.Time{}, true},
		{"invalid quantity", "one hour ago", nil, time.Time{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse.FromNaturalLanguage(tt.expr, tt.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromNaturalLanguage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !got.Equal(tt.expected) {
				t.Errorf("FromNaturalLanguage() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// 测试默认基准时间为当前时间
func TestFromNaturalLanguageDefaultBase(t *testing.T) {
	before := time.Now()
	time.Sleep(10 * time.Millisecond) // 确保时间有变化

	got, err := parse.FromNaturalLanguage("now")
	if err != nil {
		t.Fatalf("FromNaturalLanguage() with empty base time failed: %v", err)
	}

	time.Sleep(10 * time.Millisecond)
	after := time.Now()

	if got.Before(before) || got.After(after) {
		t.Errorf("FromNaturalLanguage() with empty base time returned %v, expected between %v and %v", got, before, after)
	}
}

// 测试复数形式
func TestPluralForms(t *testing.T) {
	baseTime := time.Date(2023, 5, 15, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		expr     string
		expected time.Time
	}{
		{"1 hour ago", time.Date(2023, 5, 15, 11, 0, 0, 0, time.UTC)},
		{"1 hours ago", time.Date(2023, 5, 15, 11, 0, 0, 0, time.UTC)},
		{"2 hour ago", time.Date(2023, 5, 15, 10, 0, 0, 0, time.UTC)},
		{"2 hours ago", time.Date(2023, 5, 15, 10, 0, 0, 0, time.UTC)},
	}

	for _, tt := range tests {
		t.Run(tt.expr, func(t *testing.T) {
			got, err := parse.FromNaturalLanguage(tt.expr, parse.WithFromNaturalLanguageBaseTime(baseTime))
			if err != nil {
				t.Fatalf("FromNaturalLanguage() failed: %v", err)
			}
			if !got.Equal(tt.expected) {
				t.Errorf("FromNaturalLanguage() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// 测试大小写不敏感
func TestCaseInsensitivity(t *testing.T) {
	baseTime := time.Date(2023, 5, 15, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		expr     string
		expected time.Time
	}{
		{"A HOUR AGO", time.Date(2023, 5, 15, 11, 0, 0, 0, time.UTC)},
		{"An HoUr AgO", time.Date(2023, 5, 15, 11, 0, 0, 0, time.UTC)},
		{"2 DAYS LATER", time.Date(2023, 5, 17, 12, 0, 0, 0, time.UTC)},
		{"3 MONTHS BEFORE", time.Date(2023, 2, 15, 12, 0, 0, 0, time.UTC)},
	}

	for _, tt := range tests {
		t.Run(tt.expr, func(t *testing.T) {
			got, err := parse.FromNaturalLanguage(tt.expr, parse.WithFromNaturalLanguageBaseTime(baseTime))
			if err != nil {
				t.Fatalf("FromNaturalLanguage() failed: %v", err)
			}
			if !got.Equal(tt.expected) {
				t.Errorf("FromNaturalLanguage() = %v, want %v", got, tt.expected)
			}
		})
	}
}
