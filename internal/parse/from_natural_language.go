package parse

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// FromNaturalLanguage 解析自然语言时间表达式
// 支持格式如: "an hour ago", "2 days later", "5 minutes ago" 等
func FromNaturalLanguage(expr string, opts ...func(*FromNaturalLanguageOption)) (time.Time, error) {
	cnf := new(FromNaturalLanguageOption)
	for _, opt := range opts {
		opt(cnf)
	}

	baseTime := time.Now()
	if cnf.baseTime != nil && !cnf.baseTime.IsZero() {
		baseTime = *cnf.baseTime
	}
	loc := baseTime.Location()
	if cnf.loc != nil {
		loc = cnf.loc
	}

	// 处理常见特殊表达式
	switch strings.ToLower(expr) {
	case "now":
		return baseTime.In(loc), nil
	case "today":
		return time.Date(baseTime.Year(), baseTime.Month(), baseTime.Day(), 0, 0, 0, 0, loc), nil
	case "yesterday":
		return time.Date(baseTime.Year(), baseTime.Month(), baseTime.Day()-1, 0, 0, 0, 0, loc), nil
	case "tomorrow":
		return time.Date(baseTime.Year(), baseTime.Month(), baseTime.Day()+1, 0, 0, 0, 0, loc), nil
	}

	// 正则表达式匹配模式
	re := regexp.MustCompile(`(?i)^\s*(a|an|\d+)\s+(nanosecond|microsecond|millisecond|second|minute|hour|day|week|month|year)s?\s+(ago|later|before|after)\s*$`)
	matches := re.FindStringSubmatch(expr)
	if matches == nil {
		return time.Time{}, fmt.Errorf("unsurpported time expression: %s", expr)
	}

	// 解析数量
	var quantity int
	if matches[1] == "a" || matches[1] == "an" {
		quantity = 1
	} else {
		var err error
		quantity, err = strconv.Atoi(matches[1])
		if err != nil {
			return time.Time{}, fmt.Errorf("invalid quantity: %s", matches[1])
		}
	}

	// 解析时间单位
	var duration time.Duration
	var days, months, years int
	switch strings.ToLower(matches[2]) {
	case "nanosecond":
		duration = time.Duration(quantity) * time.Nanosecond
	case "microsecond":
		duration = time.Duration(quantity) * time.Microsecond
	case "millisecond":
		duration = time.Duration(quantity) * time.Millisecond
	case "second":
		duration = time.Duration(quantity) * time.Second
	case "minute":
		duration = time.Duration(quantity) * time.Minute
	case "hour":
		duration = time.Duration(quantity) * time.Hour
	case "day":
		days = quantity
	case "week":
		days = quantity * 7
	case "month":
		months = quantity
	case "year":
		years = quantity
	default:
		return time.Time{}, fmt.Errorf("unknown time unit: %s", matches[2])
	}

	// 计算时间
	switch strings.ToLower(matches[3]) {
	case "ago", "before":
		duration = -duration
		days = -days
		months = -months
		years = -years
	case "later", "after":
		// pass
	default:
		return time.Time{}, fmt.Errorf("unknown direction: %s", matches[3])
	}

	if duration != 0 {
		return baseTime.Add(duration).In(loc), nil
	}
	if days != 0 || months != 0 || years != 0 {
		return baseTime.AddDate(years, months, days).In(loc), nil
	}

	return baseTime.In(loc), nil
}
