package chronos_test

import (
	"testing"
	"time"

	"github.com/gomooth/chronos"

	"github.com/stretchr/testify/assert"
)

func TestMaxMinFunctions(t *testing.T) {
	// 准备测试数据
	earlyTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	midTime := time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC)
	lateTime := time.Date(2023, 1, 3, 0, 0, 0, 0, time.UTC)

	var nilTimePtr *time.Time

	// 测试Max函数
	t.Run("Max with time.Time", func(t *testing.T) {
		assert.Equal(t, lateTime, chronos.Max(earlyTime, midTime, lateTime))
		assert.Equal(t, midTime, chronos.Max(earlyTime, midTime))
		assert.Equal(t, lateTime, chronos.Max(lateTime, midTime, earlyTime))
	})

	t.Run("Max with *time.Time", func(t *testing.T) {
		assert.Equal(t, &lateTime, chronos.Max(&earlyTime, &midTime, &lateTime))
		assert.Equal(t, &midTime, chronos.Max(&earlyTime, &midTime))
		assert.Equal(t, &lateTime, chronos.Max(&lateTime, &midTime, &earlyTime))
	})

	t.Run("Max with nil pointers", func(t *testing.T) {
		assert.Equal(t, &midTime, chronos.Max(nilTimePtr, &midTime))
		//assert.Equal(t, midChronos, chronos.Max(nilChronosPtr, midChronos))
		//assert.Equal(t, earlyChronos, chronos.Max(nilChronosPtr, earlyChronos, nilChronosPtr))
	})

	t.Run("Max with all nil", func(t *testing.T) {
		assert.Equal(t, nilTimePtr, chronos.Max(nilTimePtr, nilTimePtr))
		//assert.Equal(t, nilChronosPtr, chronos.Max(nilChronosPtr, nilChronosPtr))
	})

	// 测试Min函数
	t.Run("Min with time.Time", func(t *testing.T) {
		assert.Equal(t, earlyTime, chronos.Min(earlyTime, midTime, lateTime))
		assert.Equal(t, earlyTime, chronos.Min(midTime, earlyTime))
		assert.Equal(t, earlyTime, chronos.Min(lateTime, midTime, earlyTime))
	})

	t.Run("Min with *time.Time", func(t *testing.T) {
		assert.Equal(t, &earlyTime, chronos.Min(&earlyTime, &midTime, &lateTime))
		assert.Equal(t, &earlyTime, chronos.Min(&midTime, &earlyTime))
		assert.Equal(t, &earlyTime, chronos.Min(&lateTime, &midTime, &earlyTime))
	})

	t.Run("Min with nil pointers", func(t *testing.T) {
		zeroTime := time.Time{}

		assert.Equal(t, nilTimePtr, chronos.Min(nilTimePtr, &zeroTime))
	})

	t.Run("Min with all nil", func(t *testing.T) {
		assert.Equal(t, nilTimePtr, chronos.Min(nilTimePtr, nilTimePtr))
	})

	// 测试相同时间
	t.Run("Equal times", func(t *testing.T) {
		sameTime1 := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		sameTime2 := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

		assert.Equal(t, sameTime1, chronos.Max(sameTime1, sameTime2))
		assert.Equal(t, sameTime1, chronos.Min(sameTime1, sameTime2))
	})

	// 测试不同时区
	t.Run("Different timezones", func(t *testing.T) {
		utcTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		localTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.FixedZone("CST", 8*3600))

		// UTC时间比CST时间早8小时
		assert.Equal(t, utcTime, chronos.Max(utcTime, localTime))
		assert.Equal(t, localTime, chronos.Min(utcTime, localTime))
	})
}

func TestDiff(t *testing.T) {
	// 准备测试数据
	t1 := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2023, 1, 1, 1, 30, 0, 0, time.UTC) // 比t1晚1小时30分钟

	// 基本功能测试
	t.Run("Basic time diff", func(t *testing.T) {
		d := chronos.Diff(t2, t1)
		assert.Equal(t, int64(5400*1e9), d.Nanoseconds())
		assert.Equal(t, int64(5400*1e6), d.Microseconds())
		assert.Equal(t, int64(5400*1e3), d.Milliseconds())
		assert.Equal(t, int64(5400), d.Seconds())
		assert.Equal(t, int64(90), d.Minutes())
		assert.Equal(t, int64(1), d.Hours())
		assert.Equal(t, "1h30m", d.String())
	})

	// 负时间差测试
	t.Run("Negative diff", func(t *testing.T) {
		d := chronos.Diff(t1, t2)
		assert.Equal(t, int64(-5400*1e9), d.Nanoseconds())
		assert.Equal(t, "-1h30m", d.String())
	})

	// 零值时间测试
	t.Run("Zero time", func(t *testing.T) {
		d := chronos.Diff(time.Time{}, t1)
		assert.Equal(t, int64(0), d.Nanoseconds())
		assert.Equal(t, "", d.String())
	})

	// 相同时间测试
	t.Run("Same time", func(t *testing.T) {
		d := chronos.Diff(t1, t1)
		assert.Equal(t, int64(0), d.Nanoseconds())
		assert.Equal(t, "", d.String())
	})

	// 长时间测试
	t.Run("Long duration", func(t *testing.T) {
		t3 := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		t4 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
		d := chronos.Diff(t4, t3)

		assert.Equal(t, int64(1), d.Years())
		assert.Equal(t, int64(12), d.Months())
		assert.Equal(t, int64(52), d.Weeks())
		assert.Equal(t, int64(365), d.Days())
		assert.Equal(t, "8760h", d.String())
	})

	// 自定义年月天数测试
	t.Run("Custom days per year/month", func(t *testing.T) {
		t3 := time.Date(1914, 1, 1, 0, 0, 0, 0, time.UTC)
		t4 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
		d := chronos.Diff(t4, t3)

		years365 := d.Years()
		assert.Equal(t, int64(110), years365)
		// 使用360天/年
		years360 := d.Years(chronos.DiffWithDaysPer(360))
		assert.Equal(t, int64(111), years360)

		// 使用31天/月
		months30 := d.Months()
		assert.Equal(t, int64(1339), months30)
		months31 := d.Months(chronos.DiffWithDaysPer(31))
		assert.Equal(t, int64(1296), months31)
	})

	// 微秒级测试
	t.Run("Microsecond precision", func(t *testing.T) {
		t5 := time.Date(2023, 1, 1, 0, 0, 0, 500*1000, time.UTC) // 500微秒
		d := chronos.Diff(t5, t1)
		assert.Equal(t, "500μs", d.String())
	})

	// 毫秒级测试
	t.Run("Millisecond precision", func(t *testing.T) {
		t6 := time.Date(2023, 1, 1, 0, 0, 0, 500*1000*1000, time.UTC) // 500毫秒
		d := chronos.Diff(t6, t1)
		assert.Equal(t, "500ms", d.String())
	})

	// 边界条件测试
	t.Run("Boundary conditions", func(t *testing.T) {
		// 刚好1秒
		t7 := time.Date(2023, 1, 1, 0, 0, 1, 0, time.UTC)
		d := chronos.Diff(t7, t1)
		assert.Equal(t, "1s", d.String())

		// 刚好1分钟
		t8 := time.Date(2023, 1, 1, 0, 1, 0, 0, time.UTC)
		d = chronos.Diff(t8, t1)
		assert.Equal(t, "1m", d.String())

		// 刚好1小时
		t9 := time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC)
		d = chronos.Diff(t9, t1)
		assert.Equal(t, "1h", d.String())
	})

	// 不同时区测试
	t.Run("Different timezones", func(t *testing.T) {
		localTime := time.Date(2023, 1, 1, 8, 0, 0, 0, time.FixedZone("CST", 8*3600))
		d := chronos.Diff(localTime, t1)
		assert.Equal(t, int64(0), d.Nanoseconds())
		assert.Equal(t, "", d.String())
	})
}
