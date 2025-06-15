# Chronos

克洛诺斯，掌管时间的工具。
> ps: 提效后请到前台免费领鸡蛋。

目前，已支持时间解析、比较、边界问题


[English](README.md) | 简体中文

[![Static Badge](https://img.shields.io/badge/Releases-v0.1.0-green)](https://github.com/gomooth/chronos/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/gomooth/chronos)](https://goreportcard.com/report/github.com/gomooth/chronos)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)


## 安装

> go version 1.21+

```shell
go get -u github.com/gomooth/chronos

```
或
```golang
import "github.com/gomooth/chronos"
```

## 使用

### 一、昨天、明天

```golang
// 昨天此刻
chronos.Yesterday(time.Now())
// 明天此刻
chronos.Tomorrow(time.Now())
```

### 二、时间解析

#### 2.1 特殊表达式
```golang
// 此时此刻
at, err := chronos.Parse("now")
// 昨日此刻
at, err := chronos.Parse("yesterday")
// 明天此刻
at, err := chronos.Parse("tomorrow")
```

#### 2.2 时间戳

对于数字的解析，会自动根据时间戳数字的位数来区分 `秒`，`毫秒`，`微秒`，`纳秒`

```golang
// seconds
at, err := chronos.Parse(1672643045)
// milliseconds
at, err := chronos.Parse(1672643045123)
// microseconds
at, err := chronos.Parse(1672643045123456)
// nanoseconds
at, err := chronos.Parse(1672643045123456789)
```

#### 2.3 标准格式

```golang
at, err := chronos.Parse("2023-04-22T18:22:15Z")
at, err := chronos.Parse("2023-04-22 18:22:15")
at, err := chronos.Parse("2023-04-22")
```

默认支持表示时间格式有：
  - `Unix`    - `Mon Jan _2 15:04:05 MST 2006`
  - `Cookie`  - `Monday, 02-Jan-2006 15:04:05 MST`
  - `Ruby`    - `Mon Jan 02 15:04:05 -0700 2006`

  - `ANSIC`       - `Mon Jan _2 15:04:05 2006`
  - `ISO8601`     - `2006-01-02T15:04:05-07:00`，`2006-01-02T15:04:05Z`
  - `RFC822`      - `02 Jan 06 15:04 MST`，`02 Jan 06 15:04 -0700`
  - `RFC850`      - `Monday, 02-Jan-06 15:04:05 MST`
  - `RFC1036`     - `Mon, 02 Jan 06 15:04:05 -0700`
  - `RFC1123`     - `Mon, 02 Jan 2006 15:04:05 MST`，`Mon, 02 Jan 2006 15:04:05 -0700`
  - `RFC3339`     - `2006-01-02T15:04:05Z07:00`，`2006-01-02T15:04:05.999999999Z07:00`
  - `RFC7231`     - `Mon, 02 Jan 2006 15:04:05 MST`

  - `Stamp`      - `Jan _2 15:04:05`，`Jan _2 15:04:05.000`，`Jan _2 15:04:05.000000`，`Jan _2 15:04:05.000000000`
  - `DateTime`   - `2006-01-02 15:04:05`
  - `Date`       - `2006-01-02`，`2006/01/02`
  - `Time`       - `15:04:05`，`3:04PM`

对于非标准的自定义格式，可以通过 `chronos.ParseWithLayout(layout)` 来设置

```golang
input := "22/09/2023"
at, err := chronos.Parse(input, chronos.ParseWithLayout("02/01/2006"))
```

#### 2.4 自定义时区
时间解析时，默认使用本地时区。可以通过 `chronos.ParseWithLocation(loc)` 设置时区

```golang
input := "2023-09-22"
loc := time.FixedZone("TEST", 3600)
at, err := chronos.Parse(input, chronos.ParseWithLocation(loc))
```

#### 2.5 自然语言表达式
默认未开启自然语言表达式解析。需要通过 `chronos.ParseWithNaturalLanguage(true)` 开启该功能。

```golang
at, err := chronos.Parse(
	"an hour age", 
	chronos.ParseWithNaturalLanguage(true), // 启用自然语言表达式解析
)
```

自然语言解析时，默认是以当前时间基准来计算的。可以通过 `chronos.ParseWithBaseTime(at)` 来指定基准时间

```golang
// 自定义基准时间
base := time.Date(2023, 5, 15, 12, 0, 0, 0, time.UTC)
at, err := chronos.Parse(
	"an hour age", 
	chronos.ParseWithNaturalLanguage(true),
	chronos.ParseWithBaseTime(base), // 设置基准时间
)
```

- 支持的时间单位:
  - `nanosecond` - 纳秒
  - `microsecond` - 微秒
  - `millisecond` - 毫秒
  - `second` - 秒
  - `minute` - 分钟
  - `hour` - 小时
  - `day` - 天
  - `week` - 周
  - `month` - 月
  - `year` - 年
- 支持的方向:
  - `ago/before` - 过去
  - `later/after` - 未来
- 数量表示:
  - 可以使用数字 (如 `2 hours ago`)
  - 可以使用 `a` 或 `an` (如 `a hour ago`， `an hour ago`)

### 三、时间比较

#### 3.1 最值
多个时间的最大值、最小值。支持 `time.Time`，`*time.Time`，不支持混合比较
```golang
// 最大值
chronos.Max(time1, time2, time3)
// 最小值
chronos.Min(time1, time2, time3)
```

#### 3.2 差值
比较2个时间的差值，返回相差的纳秒 `DiffValue`。
```golang
diff := chronos.Diff(t2, t1)

// 将差值转换单位
diff.Nanoseconds()
diff.Microseconds()
diff.Milliseconds()
diff.Seconds()
diff.Minutes()
diff.Hours()
diff.Days()
diff.Weeks()
diff.Months()
diff.Months(chronos.DiffWithDaysPer(31))
diff.Years()
diff.Years(chronos.DiffWithDaysPer(360))
// 将差值显示人类友好字符串
diff.String()
```

## 四、时间边界

```golang
// 指定时间所在小时的起止时间
chronos.StartOfHour(at)
chronos.EndOfHour(at)

// 指定时间所在日的起止时间
chronos.StartOfDay(at)
chronos.EndOfDay(at)

// 指定时间所在周的起止时间
// 默认周一为一周的起始日，可以通过 `chronos.WithWeekStartDay()` 选项修改
chronos.StartOfWeek(at)
chronos.EndOfWeek(at)
// 设置周日为一周起始日
chronos.StartOfWeek(at, chronos.WithWeekStartDay(time.Sunday))
chronos.EndOfWeek(at, chronos.WithWeekStartDay(time.Sunday))

// 指定时间所在月的起止时间
chronos.StartOfMonth(at)
chronos.EndOfMonth(at)

// 指定时间所在季度的起止时间
chronos.StartOfQuarter(at)
chronos.EndOfQuarter(at)

// 指定时间所在年的起止时间
chronos.StartOfYear(at)
chronos.EndOfYear(at)
```

## 五、辅助

### 5.1 是否闰年
```golang
chronos.IsLeap(at)
```

### 5.2 计算某月的天数
```golang
chronos.DaysInMonth(at)
```

