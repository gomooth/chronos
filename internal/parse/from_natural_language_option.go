package parse

import "time"

type FromNaturalLanguageOption struct {
	baseTime *time.Time
	loc      *time.Location
}

func WithFromNaturalLanguageBaseTime(base time.Time) func(*FromNaturalLanguageOption) {
	return func(o *FromNaturalLanguageOption) {
		if !base.IsZero() {
			o.baseTime = &base
		}
	}
}

func WithFromNaturalLanguageLocation(loc *time.Location) func(*FromNaturalLanguageOption) {
	return func(o *FromNaturalLanguageOption) {
		o.loc = loc
	}
}
