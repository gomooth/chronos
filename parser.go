package chronos

import (
	"fmt"
	"time"

	"github.com/gomooth/chronos/internal/parse"
)

type ParseOption struct {
	fromStringOptions   []func(*parse.FromStringOption)
	fromNaturalLanguage struct {
		supported bool
		options   []func(*parse.FromNaturalLanguageOption)
	}
}

// Parse 时间解析
func Parse[T TimeValue](v T, opts ...func(*ParseOption)) (*time.Time, error) {
	cnf := new(ParseOption)
	for _, opt := range opts {
		opt(cnf)
	}

	now := time.Now()
	switch val := any(v).(type) {
	case time.Time:
		return &val, nil
	case *time.Time:
		return val, nil
	case int, int16, int32, int64, uint, uint16, uint32, uint64, uintptr:
		t, err := parse.FromUnixTime(val)
		if err != nil {
			return nil, fmt.Errorf("invalid unix time: %w", err)
		}
		return &t, nil
	default:
		str := val.(string)
		switch str {
		case "now":
			return &now, nil
		case "yesterday":
			at := Yesterday(now)
			return &at, nil
		case "tomorrow":
			at := Tomorrow(now)
			return &at, nil
		default:
			// 尝试解析字符串为时间
			parsed, err := parse.FromStringFormat(str, cnf.fromStringOptions...)
			if err != nil {
				if !cnf.fromNaturalLanguage.supported {
					return nil, fmt.Errorf("invalid time string: %w", err)
				}
				// 尝试解析自然语言
				parsed, err = parse.FromNaturalLanguage(str, cnf.fromNaturalLanguage.options...)
				if err != nil {
					return nil, fmt.Errorf("invalid time string: %w", err)
				}
			}
			return &parsed, nil
		}
	}
}

// ParseWithLayout 指定时间解析的自定义格式
func ParseWithLayout(layout string, others ...string) func(*ParseOption) {
	return func(p *ParseOption) {
		if p.fromStringOptions == nil {
			p.fromStringOptions = make([]func(*parse.FromStringOption), 0)
		}
		p.fromStringOptions = append(p.fromStringOptions, parse.WithFromStringLayout(layout))
		for _, other := range others {
			p.fromStringOptions = append(p.fromStringOptions, parse.WithFromStringLayout(other))
		}
	}
}

// ParseWithLocation 指定时间解析的时区
func ParseWithLocation(loc *time.Location) func(*ParseOption) {
	return func(p *ParseOption) {
		if p.fromStringOptions == nil {
			p.fromStringOptions = make([]func(*parse.FromStringOption), 0)
		}
		p.fromStringOptions = append(p.fromStringOptions, parse.WithFromStringLocation(loc))

		if p.fromNaturalLanguage.options == nil {
			p.fromNaturalLanguage.options = make([]func(*parse.FromNaturalLanguageOption), 0)
		}
		p.fromNaturalLanguage.options = append(p.fromNaturalLanguage.options, parse.WithFromNaturalLanguageLocation(loc))
	}
}

// ParseWithBaseTime 指定时间解析的时区
func ParseWithBaseTime(base time.Time) func(*ParseOption) {
	return func(p *ParseOption) {
		if p.fromNaturalLanguage.options == nil {
			p.fromNaturalLanguage.options = make([]func(*parse.FromNaturalLanguageOption), 0)
		}
		p.fromNaturalLanguage.options = append(p.fromNaturalLanguage.options, parse.WithFromNaturalLanguageBaseTime(base))
	}
}

// ParseWithNaturalLanguage 指定时间解析是否支持自然语言
func ParseWithNaturalLanguage(supported bool) func(*ParseOption) {
	return func(p *ParseOption) {
		p.fromNaturalLanguage.supported = supported
	}
}
