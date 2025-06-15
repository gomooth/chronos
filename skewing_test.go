package chronos_test

import (
	"testing"
	"time"

	"github.com/gomooth/chronos"

	"github.com/stretchr/testify/assert"
)

func TestTomorrow(t *testing.T) {
	// 准备测试数据
	now := time.Date(2023, 5, 15, 14, 30, 15, 123456789, time.UTC)
	tomorrow := time.Date(2023, 5, 16, 14, 30, 15, 123456789, time.UTC)

	tests := []struct {
		name     string
		input    interface{}
		expected time.Time
	}{
		{
			"time.Time 类型",
			now,
			tomorrow,
		},
		{
			"*time.Time 类型",
			&now,
			tomorrow,
		},
		{
			"nil *time.Time 类型",
			(*time.Time)(nil),
			time.Now().Add(24 * time.Hour).Truncate(time.Second), // 忽略纳秒差异
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result time.Time

			switch v := tt.input.(type) {
			case time.Time:
				result = chronos.Tomorrow(v)
			case *time.Time:
				result = chronos.Tomorrow(v)
			}

			// 对于涉及当前时间的测试，忽略秒级以下的差异
			if tt.name == "nil *time.Time 类型" || tt.name == "其他类型（使用当前时间）" {
				assert.Equal(t, tt.expected.Truncate(time.Second), result.Truncate(time.Second))
			} else {
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestYesterday(t *testing.T) {
	// 准备测试数据
	now := time.Date(2023, 5, 15, 14, 30, 15, 123456789, time.UTC)
	yesterday := time.Date(2023, 5, 14, 14, 30, 15, 123456789, time.UTC)

	tests := []struct {
		name     string
		input    interface{}
		expected time.Time
	}{
		{
			"time.Time 类型",
			now,
			yesterday,
		},
		{
			"*time.Time 类型",
			&now,
			yesterday,
		},
		{
			"nil *time.Time 类型",
			(*time.Time)(nil),
			time.Now().Add(-24 * time.Hour).Truncate(time.Second),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result time.Time

			switch v := tt.input.(type) {
			case time.Time:
				result = chronos.Yesterday(v)
			case *time.Time:
				result = chronos.Yesterday(v)
			}

			// 对于涉及当前时间的测试，忽略秒级以下的差异
			if tt.name == "nil *time.Time 类型" || tt.name == "其他类型（使用当前时间）" {
				assert.Equal(t, tt.expected.Truncate(time.Second), result.Truncate(time.Second))
			} else {
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

// 测试跨月边界
func TestMonthBoundary(t *testing.T) {
	// 测试跨月
	t.Run("Tomorrow across month", func(t *testing.T) {
		lastDay := time.Date(2023, 4, 30, 0, 0, 0, 0, time.UTC)
		expected := time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, expected, chronos.Tomorrow(lastDay))
	})

	t.Run("Yesterday across month", func(t *testing.T) {
		firstDay := time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC)
		expected := time.Date(2023, 4, 30, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, expected, chronos.Yesterday(firstDay))
	})
}

// 测试闰年
func TestLeapYear(t *testing.T) {
	t.Run("Tomorrow in leap year", func(t *testing.T) {
		feb28 := time.Date(2024, 2, 28, 0, 0, 0, 0, time.UTC)
		expected := time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, expected, chronos.Tomorrow(feb28))
	})

	t.Run("Yesterday after leap day", func(t *testing.T) {
		mar1 := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
		expected := time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, expected, chronos.Yesterday(mar1))
	})
}
