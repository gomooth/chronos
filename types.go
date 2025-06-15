package chronos

import "time"

type TimeValue interface {
	~string |
		~int | ~int16 | ~int32 | ~int64 |
		~uint | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		time.Time | *time.Time
}

type MixedTime interface {
	time.Time | *time.Time
}
