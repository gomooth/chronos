package parse

import (
	"fmt"
	"time"

	"github.com/gomooth/chronos/timelayout"
)

var stringFormats = []string{
	time.Layout, time.ANSIC, time.UnixDate, time.RubyDate,
	time.RFC822, time.RFC822Z, time.RFC850, time.RFC1123, time.RFC1123Z, time.RFC3339, time.RFC3339Nano,
	time.Kitchen,
	time.Stamp, time.StampMilli, time.StampMicro, time.StampNano,
	time.DateTime, time.DateOnly, time.TimeOnly,

	timelayout.RFC1036, timelayout.RFC7231,
	timelayout.ISO8601, timelayout.ISO8601Zulu,
	timelayout.Cookie,

	"2006/01/02",
}

// FromStringFormat 尝试解析各种格式的时间字符串
func FromStringFormat(s string, opts ...func(*FromStringOption)) (time.Time, error) {
	cnf := new(FromStringOption)
	for _, opt := range opts {
		opt(cnf)
	}

	formats := stringFormats
	if len(cnf.layouts) > 0 {
		formats = append(cnf.layouts, formats...)
	}

	loc := time.Local
	if cnf.loc != nil {
		loc = cnf.loc
	}

	for _, format := range formats {
		if t, err := time.ParseInLocation(format, s, loc); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("could not parse time string: %s", s)
}
