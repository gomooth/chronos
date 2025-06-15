package chronos_test

import (
	"testing"
	"time"

	"github.com/gomooth/chronos"
)

func TestIsLeap(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  bool
	}{
		// time.Time 测试
		{
			name:  "time.Time leap year",
			input: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			want:  true,
		},
		{
			name:  "time.Time non-leap year",
			input: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			want:  false,
		},
		{
			name:  "time.Time century leap year (2000)",
			input: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			want:  true,
		},
		{
			name:  "time.Time century non-leap year (1900)",
			input: time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC),
			want:  false,
		},

		// *time.Time 测试
		{
			name:  "*time.Time leap year",
			input: ptr(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
			want:  true,
		},
		{
			name:  "*time.Time non-leap year",
			input: ptr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got bool
			switch v := tt.input.(type) {
			case time.Time:
				got = chronos.IsLeap(v)
			case *time.Time:
				got = chronos.IsLeap(v)
			default:
				t.Fatalf("unsupported input type: %T", v)
			}

			if got != tt.want {
				t.Errorf("IsLeap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDaysInMonth(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  int
	}{
		// time.Time 测试
		{
			name:  "time.Time January",
			input: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			want:  31,
		},
		{
			name:  "time.Time February non-leap",
			input: time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC),
			want:  28,
		},
		{
			name:  "time.Time February leap",
			input: time.Date(2020, 2, 1, 0, 0, 0, 0, time.UTC),
			want:  29,
		},
		{
			name:  "time.Time April",
			input: time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC),
			want:  30,
		},

		// *time.Time 测试
		{
			name:  "*time.Time December",
			input: ptr(time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC)),
			want:  31,
		},
		{
			name:  "*time.Time November",
			input: ptr(time.Date(2023, 11, 1, 0, 0, 0, 0, time.UTC)),
			want:  30,
		},
		{
			name:  "*time.Time February 1900",
			input: ptr(time.Date(1900, 2, 1, 0, 0, 0, 0, time.UTC)),
			want:  28,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got int
			switch v := tt.input.(type) {
			case time.Time:
				got = chronos.DaysInMonth(v)
			case *time.Time:
				got = chronos.DaysInMonth(v)
			default:
				t.Fatalf("unsupported input type: %T", v)
			}

			if got != tt.want {
				t.Errorf("DaysInMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 辅助函数，创建时间指针
func ptr(t time.Time) *time.Time {
	return &t
}
