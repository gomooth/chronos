package chronos_test

import (
	"testing"
	"time"

	"github.com/gomooth/chronos"

	"github.com/stretchr/testify/assert"
)

func TestParse_StringInput(t *testing.T) {
	t.Run("now", func(t *testing.T) {
		at, err := chronos.Parse("now")
		assert.NoError(t, err)
		assert.Equal(t, at.Format("2006-01-02"), time.Now().Format("2006-01-02"))
	})

	t.Run("yesterday", func(t *testing.T) {
		at, err := chronos.Parse("yesterday")
		assert.NoError(t, err)
		assert.Equal(t, at.Format("2006-01-02"), time.Now().Add(-24*time.Hour).Format("2006-01-02"))
	})

	t.Run("tomorrow", func(t *testing.T) {
		at, err := chronos.Parse("tomorrow")
		assert.NoError(t, err)
		assert.Equal(t, at.Format("2006-01-02"), time.Now().Add(24*time.Hour).Format("2006-01-02"))
	})

	t.Run("RFC3339 format", func(t *testing.T) {
		v := "2023-04-22T18:22:15Z"
		at, err := chronos.Parse(v)
		assert.NoError(t, err)
		assert.Equal(t, at.Format("2006-01-02 15:04:05"), "2023-04-22 18:22:15")
	})

	t.Run("DateOnly format", func(t *testing.T) {
		v := "2023-04-22"
		at, err := chronos.Parse(v)
		assert.NoError(t, err)
		assert.Equal(t, at.Format("2006-01-02"), "2023-04-22")
	})

	t.Run("invalid format", func(t *testing.T) {
		v := "invalid-time-string"
		at, err := chronos.Parse(v)
		assert.Error(t, err)
		assert.Nil(t, at)
	})
}

func TestParse_IntegerInput(t *testing.T) {
	t.Run("int seconds", func(t *testing.T) {
		v := int(1672643045)
		at, err := chronos.Parse(v)
		assert.NoError(t, err)
		assert.Equal(t, *at, time.Unix(int64(v), 0))
	})

	t.Run("int64 milliseconds", func(t *testing.T) {
		v := int64(1672643045123)
		at, err := chronos.Parse(v)
		assert.NoError(t, err)
		assert.Equal(t, *at, time.Unix(1672643045, 123000000))
	})

	t.Run("int64 microseconds", func(t *testing.T) {
		v := int64(1672643045123456)
		at, err := chronos.Parse(v)
		assert.NoError(t, err)
		assert.Equal(t, *at, time.Unix(1672643045, 123456000))
	})

	t.Run("int64 nanoseconds", func(t *testing.T) {
		v := int64(1672643045123456789)
		at, err := chronos.Parse(v)
		assert.NoError(t, err)
		assert.Equal(t, *at, time.Unix(0, 1672643045123456789))
	})

	t.Run("uint32", func(t *testing.T) {
		v := uint32(1672643045)
		at, err := chronos.Parse(v)
		assert.NoError(t, err)
		assert.Equal(t, *at, time.Unix(int64(v), 0))
	})

	t.Run("int64 too large", func(t *testing.T) {
		v := uint64(1<<64 - 1)
		at, err := chronos.Parse(v)
		assert.Error(t, err)
		assert.Nil(t, at)
	})
}

func TestParse_TimeInput(t *testing.T) {
	now := time.Now()
	zeroTime := time.Time{}

	t.Run("time.Time", func(t *testing.T) {
		at, err := chronos.Parse(now)
		assert.NoError(t, err)
		assert.Equal(t, at, &now)
	})

	t.Run("*time.Time", func(t *testing.T) {
		at, err := chronos.Parse(&now)
		assert.NoError(t, err)
		assert.Equal(t, at, &now)
	})

	t.Run("zero time.Time", func(t *testing.T) {
		at, err := chronos.Parse(zeroTime)
		assert.NoError(t, err)
		assert.Equal(t, at, &zeroTime)
	})
}

func TestParse_WithOptions(t *testing.T) {
	t.Run("custom layout", func(t *testing.T) {
		input := "22/09/2023"
		at, err := chronos.Parse(input, chronos.ParseWithLayout("02/01/2006"))
		assert.NoError(t, err)
		assert.Equal(t, at.Format("2006-01-02"), "2023-09-22")
	})

	t.Run("custom more layout", func(t *testing.T) {
		input := "2023.09.22"
		at, err := chronos.Parse(input, chronos.ParseWithLayout("02/01/2006", "2006.01.02"))
		assert.NoError(t, err)
		assert.Equal(t, at.Format("2006-01-02"), "2023-09-22")
	})

	t.Run("custom timezone", func(t *testing.T) {
		input := "2023-09-22"
		loc := time.FixedZone("TEST", 3600)

		at, err := chronos.Parse(input, chronos.ParseWithLocation(loc))
		assert.NoError(t, err)
		assert.Equal(t,
			at.Format("2006-01-02 15:04:05"),
			time.Date(2023, 9, 22, 0, 0, 0, 0, loc).Format("2006-01-02 15:04:05"),
		)
	})
}

func TestParse_NaturalLanguage(t *testing.T) {
	// 固定基准时间，便于测试
	baseTime := time.Date(2023, 5, 15, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		input    string
		expected time.Time
		wantErr  bool
	}{
		// 基本自然语言表达式
		{
			"1 hour ago",
			"1 hour ago",
			time.Date(2023, 5, 15, 11, 0, 0, 0, time.UTC),
			false,
		},
		{
			"2 days later",
			"2 days later",
			time.Date(2023, 5, 17, 12, 0, 0, 0, time.UTC),
			false,
		},
		{
			"a month ago",
			"a month ago",
			time.Date(2023, 4, 15, 12, 0, 0, 0, time.UTC),
			false,
		},

		// 无效自然语言表达式
		{
			"invalid natural language",
			"invalid time string",
			time.Time{},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := chronos.Parse(tt.input,
				chronos.ParseWithNaturalLanguage(true),
				chronos.ParseWithBaseTime(baseTime),
			)
			if !tt.wantErr {
				assert.NoError(t, err)
				assert.Equal(t, &tt.expected, got)
			} else {
				assert.Error(t, err)
			}
		})
	}

	// 时区测试
	t.Run("New York timezone", func(t *testing.T) {
		input := "1 hour ago"
		locNY, _ := time.LoadLocation("America/New_York")

		at, err := chronos.Parse(input,
			chronos.ParseWithNaturalLanguage(true),
			chronos.ParseWithBaseTime(baseTime),
			chronos.ParseWithLocation(locNY),
		)
		assert.NoError(t, err)
		assert.Equal(t,
			*at,
			time.Date(2023, 5, 15, 7, 0, 0, 0, locNY), // UTC-5 in May (EDT),
		)
	})

	// 自然语言解析关闭的情况
	t.Run("natural language disabled", func(t *testing.T) {
		at, err := chronos.Parse("1 hour ago", chronos.ParseWithBaseTime(baseTime))
		assert.Error(t, err)
		assert.Nil(t, at)
	})
}
