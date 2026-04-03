# datetime

A lightweight Go utility package for common date and time operations.

## Installation

```bash
go get github.com/alex-cos/datetime
```

## Usage

```go
import (
    "fmt"
    "time"

    "github.com/alex-cos/datetime"
)

func main() {
    now := time.Now()

    // Formatting
    fmt.Println(datetime.SimpleFormat(now))       // "2020-02-14 12:12:12"
    fmt.Println(datetime.FormatISO8601Millis(now)) // "2020-02-14T12:12:12.000Z"

    // Parsing
    t, err := datetime.ParseISO8601Millis("2020-02-14T12:12:12.000Z")

    // Month navigation
    datetime.LastMonth(now)
    datetime.ThisMonth(now)
    datetime.NextMonth(now)

    // Quarter navigation
    datetime.LastQuarter(now)
    datetime.ThisQuarter(now)
    datetime.NextQuarter(now)

    // Year navigation
    datetime.LastYear(now)
    datetime.ThisYear(now)
    datetime.NextYear(now)

    // Checks
    datetime.IsBeginingMonth(now)
    datetime.IsBeginingQuarter(now)
    datetime.IsBeginingYear(now)
}
```

## API

### Formatting

| Function | Description |
| --- | --- |
| `SimpleFormat(t time.Time) string` | Formats as `YYYY-MM-DD HH:MM:SS` in UTC |
| `FormatISO8601Millis(t time.Time) string` | Formats as ISO 8601 with fixed milliseconds in UTC |
| `ParseISO8601Millis(str string) (time.Time, error)` | Parses `YYYY-MM-DDTHH:MM:SS.mmmZ` format |

### Navigation

All navigation functions return the **first day** of the target period at `00:00:00 UTC`.

| Function | Description |
| --- | --- |
| `LastMonth(t time.Time) time.Time` | First day of the previous month |
| `ThisMonth(t time.Time) time.Time` | First day of the current month |
| `NextMonth(t time.Time) time.Time` | First day of the next month |
| `LastQuarter(t time.Time) time.Time` | First day of the previous quarter |
| `ThisQuarter(t time.Time) time.Time` | First day of the current quarter |
| `NextQuarter(t time.Time) time.Time` | First day of the next quarter |
| `LastYear(t time.Time) time.Time` | First day of the previous year |
| `ThisYear(t time.Time) time.Time` | First day of the current year |
| `NextYear(t time.Time) time.Time` | First day of the next year |

### Checks

| Function | Description |
| --- | --- |
| `IsBeginingMonth(t time.Time) bool` | True if `t` is the first day of a month |
| `IsBeginingQuarter(t time.Time) bool` | True if `t` is the first day of a quarter (Jan, Apr, Jul, Oct) |
| `IsBeginingYear(t time.Time) bool` | True if `t` is January 1st |

### Utilities

| Function | Description |
| --- | --- |
| `MidNightUTC(t time.Time) time.Time` | Truncates to the start of the UTC day |
| `IsMidNightUTC(t time.Time) bool` | True if `t` falls within the midnight UTC hour (00:00–00:59) |

### Constants

| Constant | Value | Description |
| --- | --- | --- |
| `Day` | `24 * time.Hour` | Duration of a fixed 24-hour day |
