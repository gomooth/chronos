package parse

import (
	"fmt"
	"time"
)

// FromUnixTime 解析Unix时间戳
func FromUnixTime(v any) (time.Time, error) {
	var sec, nsec int64

	switch val := v.(type) {
	case int:
		sec, nsec = int64(val), 0
	case int16:
		sec, nsec = int64(val), 0
	case int32:
		sec, nsec = int64(val), 0
	case int64:
		if val > 1e18 { // 纳秒
			sec, nsec = 0, val
		} else if val > 1e15 { // 微秒
			sec, nsec = 0, val*1e3
		} else if val > 1e12 { // 毫秒
			sec, nsec = 0, val*1e6
		} else {
			sec, nsec = val, 0
		}
	case uint:
		sec, nsec = int64(val), 0
	case uint16:
		sec, nsec = int64(val), 0
	case uint32:
		sec, nsec = int64(val), 0
	case uint64:
		if val > 1<<63-1 {
			return time.Time{}, fmt.Errorf("uint64 value too large: %d", val)
		}
		sec, nsec = int64(val), 0
	case uintptr:
		sec, nsec = int64(val), 0
	}

	return time.Unix(sec, nsec), nil
}
