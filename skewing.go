package chronos

import (
	"time"
)

// Tomorrow 明天的时刻
func Tomorrow[T MixedTime](at T) time.Time {
	src := time.Now()
	switch v := any(at).(type) {
	case time.Time:
		src = v
	case *time.Time:
		if v != nil {
			src = *v
		}
	}
	return src.Add(24 * time.Hour)
}

// Yesterday 昨天的时刻
func Yesterday[T MixedTime](at T) time.Time {
	src := time.Now()
	switch v := any(at).(type) {
	case time.Time:
		src = v
	case *time.Time:
		if v != nil {
			src = *v
		}
	}
	return src.Add(-24 * time.Hour)
}
