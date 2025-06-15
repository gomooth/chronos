package chronos

import "time"

// 获取时间值的实际时间
func getTime[T MixedTime](t T) time.Time {
	switch v := any(t).(type) {
	case *time.Time:
		if v == nil {
			return time.Time{}
		}
		return *v
	default:
		return v.(time.Time)
	}
}

// Max 返回时间最晚的 chronos 实例
func Max[T MixedTime](v1, v2 T, others ...T) T {
	maxTime := v1
	times := append([]T{v2}, others...)
	for _, c := range times {
		if getTime(c).After(getTime(maxTime)) {
			maxTime = c
		}
	}
	return maxTime
}

// Min 返回时间最早的 chronos 实例
func Min[T MixedTime](v1, v2 T, others ...T) T {
	minTime := v1
	times := append([]T{v2}, others...)
	for _, c := range times {
		if getTime(c).Before(getTime(minTime)) {
			minTime = c
		}
	}
	return minTime
}

// Diff 计算两个时间的差值
func Diff[T1, T2 MixedTime](v1 T1, v2 T2) DiffValue {
	t1 := getTime(v1)
	t2 := getTime(v2)

	if t1.IsZero() || t2.IsZero() {
		return DiffValue(0)
	}

	return DiffValue(t1.Sub(t2).Nanoseconds())
}
