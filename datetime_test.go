package datetime_test

import (
	"testing"
	"time"

	"github.com/alex-cos/datetime"
	"github.com/stretchr/testify/assert"
)

func TestSimpleFormat(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{
			name:     "UTC time",
			input:    time.Date(2020, time.February, 14, 12, 12, 12, 0, time.UTC),
			expected: "2020-02-14 12:12:12",
		},
		{
			name:     "non-UTC timezone is converted",
			input:    time.Date(2020, time.February, 14, 12, 12, 12, 0, time.FixedZone("Paris", 3600)),
			expected: "2020-02-14 11:12:12",
		},
		{
			name:     "midnight",
			input:    time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: "2020-01-01 00:00:00",
		},
		{
			name:     "nanoseconds are truncated",
			input:    time.Date(2020, time.March, 15, 10, 30, 45, 123456789, time.UTC),
			expected: "2020-03-15 10:30:45",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res := datetime.SimpleFormat(tt.input)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestFormatISO8601Millis(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{
			name:     "zero nanoseconds",
			input:    time.Date(2020, time.February, 14, 12, 12, 12, 0, time.UTC),
			expected: "2020-02-14T12:12:12.000Z",
		},
		{
			name:     "with nanoseconds truncated to millis",
			input:    time.Date(2020, time.February, 14, 12, 12, 12, 123456789, time.UTC),
			expected: "2020-02-14T12:12:12.123Z",
		},
		{
			name:     "non-UTC timezone is converted",
			input:    time.Date(2020, time.February, 14, 13, 12, 12, 0, time.FixedZone("Paris", 3600)),
			expected: "2020-02-14T12:12:12.000Z",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res := datetime.FormatISO8601Millis(tt.input)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestParseISO8601Millis(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		input       string
		expected    time.Time
		expectError bool
	}{
		{
			name:     "valid datetime",
			input:    "2020-02-14T12:12:12.000Z",
			expected: time.Date(2020, time.February, 14, 12, 12, 12, 0, time.UTC),
		},
		{
			name:     "valid datetime with millis",
			input:    "2020-02-14T12:12:12.456Z",
			expected: time.Date(2020, time.February, 14, 12, 12, 12, 456000000, time.UTC),
		},
		{
			name:        "invalid format with space",
			input:       "2020-02-14 11:12:12",
			expectError: true,
		},
		{
			name:        "invalid format without T",
			input:       "2020-02-14 12:12:12.000Z",
			expectError: true,
		},
		{
			name:        "empty string",
			input:       "",
			expectError: true,
		},
		{
			name:        "timezone offset not supported",
			input:       "2020-02-14T12:12:12.000+01:00",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res, err := datetime.ParseISO8601Millis(tt.input)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, res)
			}
		})
	}
}

func TestMidNightUTC(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "noon UTC truncates to midnight same day",
			input:    time.Date(2020, time.February, 14, 12, 0, 0, 0, time.UTC),
			expected: time.Date(2020, time.February, 14, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "already at midnight UTC",
			input:    time.Date(2020, time.February, 14, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2020, time.February, 14, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "non-UTC timezone with afternoon time",
			input:    time.Date(2020, time.February, 14, 12, 12, 12, 0, time.FixedZone("Paris", 3600)),
			expected: time.Date(2020, time.February, 14, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "non-UTC timezone where UTC date differs",
			input:    time.Date(2020, time.February, 14, 0, 35, 0, 0, time.FixedZone("Tokyo", 9*3600)),
			expected: time.Date(2020, time.February, 13, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "with nanoseconds",
			input:    time.Date(2020, time.February, 14, 23, 59, 59, 999999999, time.UTC),
			expected: time.Date(2020, time.February, 14, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res := datetime.MidNightUTC(tt.input)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestIsMidNightUTC(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    time.Time
		expected bool
	}{
		{
			name:     "exact midnight UTC",
			input:    time.Date(2020, time.February, 14, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "35 minutes past midnight UTC",
			input:    time.Date(2020, time.February, 14, 0, 35, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "noon UTC is not midnight",
			input:    time.Date(2020, time.February, 14, 12, 12, 12, 0, time.UTC),
			expected: false,
		},
		{
			name:     "non-UTC timezone converted to midnight UTC hour",
			input:    time.Date(2020, time.February, 1, 1, 12, 0, 0, time.FixedZone("Paris", 3600)),
			expected: true,
		},
		{
			name:     "1 AM UTC is not midnight",
			input:    time.Date(2020, time.February, 14, 1, 0, 0, 0, time.UTC),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res := datetime.IsMidNightUTC(tt.input)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestLastMonth(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "February → January same year",
			input:    time.Date(2020, time.February, 14, 12, 12, 12, 0, time.UTC),
			expected: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "January → December previous year",
			input:    time.Date(2020, time.January, 14, 12, 12, 12, 0, time.UTC),
			expected: time.Date(2019, time.December, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "December → November same year",
			input:    time.Date(2020, time.December, 31, 23, 59, 59, 0, time.UTC),
			expected: time.Date(2020, time.November, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "non-UTC timezone",
			input:    time.Date(2020, time.February, 14, 12, 12, 12, 0, time.FixedZone("Tokyo", 9*3600)),
			expected: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "first day of month",
			input:    time.Date(2020, time.March, 1, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2020, time.February, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res := datetime.LastMonth(tt.input)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestThisMonth(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "mid-month",
			input:    time.Date(2020, time.February, 14, 12, 12, 12, 0, time.UTC),
			expected: time.Date(2020, time.February, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "January",
			input:    time.Date(2020, time.January, 14, 12, 12, 12, 0, time.UTC),
			expected: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "December",
			input:    time.Date(2020, time.December, 31, 23, 59, 59, 0, time.UTC),
			expected: time.Date(2020, time.December, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "already first day of month",
			input:    time.Date(2020, time.February, 1, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2020, time.February, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "non-UTC timezone",
			input:    time.Date(2020, time.February, 14, 12, 12, 12, 0, time.FixedZone("Tokyo", 9*3600)),
			expected: time.Date(2020, time.February, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res := datetime.ThisMonth(tt.input)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestNextMonth(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "February → March same year",
			input:    time.Date(2020, time.February, 14, 12, 12, 12, 0, time.UTC),
			expected: time.Date(2020, time.March, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "December → January next year",
			input:    time.Date(2019, time.December, 14, 12, 12, 12, 0, time.UTC),
			expected: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "November → December same year",
			input:    time.Date(2020, time.November, 30, 23, 59, 59, 0, time.UTC),
			expected: time.Date(2020, time.December, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "non-UTC timezone",
			input:    time.Date(2020, time.February, 14, 12, 12, 12, 0, time.FixedZone("Tokyo", 9*3600)),
			expected: time.Date(2020, time.March, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "already first day of month",
			input:    time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2020, time.February, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res := datetime.NextMonth(tt.input)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestThisQuarter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "Q1 (Feb)",
			input:    time.Date(2020, time.February, 14, 12, 12, 12, 0, time.UTC),
			expected: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q1 (Jan) first day",
			input:    time.Date(2020, time.January, 2, 12, 12, 12, 0, time.UTC),
			expected: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q2 (Apr)",
			input:    time.Date(2020, time.April, 15, 10, 0, 0, 0, time.UTC),
			expected: time.Date(2020, time.April, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q2 (Jun) last day",
			input:    time.Date(2020, time.June, 30, 23, 59, 59, 0, time.UTC),
			expected: time.Date(2020, time.April, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q3 (Jul)",
			input:    time.Date(2020, time.July, 1, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2020, time.July, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q3 (Sep)",
			input:    time.Date(2020, time.September, 15, 12, 0, 0, 0, time.UTC),
			expected: time.Date(2020, time.July, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q4 (Oct)",
			input:    time.Date(2020, time.October, 1, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2020, time.October, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q4 (Dec) last day",
			input:    time.Date(2020, time.December, 31, 23, 59, 59, 0, time.UTC),
			expected: time.Date(2020, time.October, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "non-UTC timezone",
			input:    time.Date(2020, time.February, 14, 12, 12, 12, 0, time.FixedZone("Tokyo", 9*3600)),
			expected: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res := datetime.ThisQuarter(tt.input)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestNextQuarter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "Q1 (Feb) → Q2 same year",
			input:    time.Date(2020, time.February, 14, 12, 12, 12, 0, time.UTC),
			expected: time.Date(2020, time.April, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q1 (Jan) → Q2 same year",
			input:    time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2020, time.April, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q2 (Apr) → Q3 same year",
			input:    time.Date(2020, time.April, 15, 10, 0, 0, 0, time.UTC),
			expected: time.Date(2020, time.July, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q2 (Jun) → Q3 same year",
			input:    time.Date(2020, time.June, 30, 23, 59, 59, 0, time.UTC),
			expected: time.Date(2020, time.July, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q3 (Jul) → Q4 same year",
			input:    time.Date(2020, time.July, 1, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2020, time.October, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q3 (Sep) → Q4 same year",
			input:    time.Date(2020, time.September, 15, 12, 0, 0, 0, time.UTC),
			expected: time.Date(2020, time.October, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q4 (Oct) → Q1 next year",
			input:    time.Date(2020, time.October, 1, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Q4 (Dec) → Q1 next year",
			input:    time.Date(2020, time.December, 14, 12, 12, 12, 0, time.UTC),
			expected: time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "non-UTC timezone",
			input:    time.Date(2020, time.February, 14, 12, 12, 12, 0, time.FixedZone("Tokyo", 9*3600)),
			expected: time.Date(2020, time.April, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res := datetime.NextQuarter(tt.input)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestLastYear(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "mid-year",
			input:    time.Date(2020, time.February, 14, 12, 12, 12, 0, time.UTC),
			expected: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "January 1st",
			input:    time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "December 31st",
			input:    time.Date(2020, time.December, 31, 23, 59, 59, 0, time.UTC),
			expected: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "non-UTC timezone",
			input:    time.Date(2020, time.February, 14, 12, 12, 12, 0, time.FixedZone("Tokyo", 9*3600)),
			expected: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res := datetime.LastYear(tt.input)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestThisYear(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "mid-year",
			input:    time.Date(2020, time.February, 14, 12, 12, 12, 0, time.UTC),
			expected: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "already January 1st",
			input:    time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "December 31st",
			input:    time.Date(2020, time.December, 31, 23, 59, 59, 0, time.UTC),
			expected: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "non-UTC timezone",
			input:    time.Date(2020, time.February, 14, 12, 12, 12, 0, time.FixedZone("Tokyo", 9*3600)),
			expected: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res := datetime.ThisYear(tt.input)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestNextYear(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "mid-year",
			input:    time.Date(2020, time.February, 14, 12, 12, 12, 0, time.UTC),
			expected: time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "January 1st",
			input:    time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "December 31st",
			input:    time.Date(2020, time.December, 31, 23, 59, 59, 0, time.UTC),
			expected: time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "non-UTC timezone",
			input:    time.Date(2020, time.February, 14, 12, 12, 12, 0, time.FixedZone("Tokyo", 9*3600)),
			expected: time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res := datetime.NextYear(tt.input)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestIsBeginingMonth(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    time.Time
		expected bool
	}{
		{
			name:     "January 1st",
			input:    time.Date(2020, time.January, 1, 12, 12, 12, 0, time.UTC),
			expected: true,
		},
		{
			name:     "February 1st",
			input:    time.Date(2020, time.February, 1, 12, 12, 12, 0, time.UTC),
			expected: true,
		},
		{
			name:     "December 1st",
			input:    time.Date(2020, time.December, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "January 2nd is not first day",
			input:    time.Date(2020, time.January, 2, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "September 2nd is not first day",
			input:    time.Date(2020, time.September, 2, 12, 12, 12, 0, time.UTC),
			expected: false,
		},
		{
			name:     "last day of month",
			input:    time.Date(2020, time.January, 31, 23, 59, 59, 0, time.UTC),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res := datetime.IsBeginingMonth(tt.input)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestIsBeginingQuarter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    time.Time
		expected bool
	}{
		{
			name:     "January 1st",
			input:    time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "April 1st",
			input:    time.Date(2020, time.April, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "July 1st",
			input:    time.Date(2020, time.July, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "October 1st",
			input:    time.Date(2020, time.October, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "February 1st is not quarter start",
			input:    time.Date(2020, time.February, 1, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "April 11th is not quarter start",
			input:    time.Date(2020, time.April, 11, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "March 31st is not quarter start",
			input:    time.Date(2020, time.March, 31, 23, 59, 59, 0, time.UTC),
			expected: false,
		},
		{
			name:     "December 31st is not quarter start",
			input:    time.Date(2020, time.December, 31, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res := datetime.IsBeginingQuarter(tt.input)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestIsBeginingYear(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    time.Time
		expected bool
	}{
		{
			name:     "January 1st",
			input:    time.Date(2020, time.January, 1, 12, 12, 12, 0, time.UTC),
			expected: true,
		},
		{
			name:     "January 1st at midnight",
			input:    time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "February 1st is not year start",
			input:    time.Date(2020, time.February, 1, 12, 12, 12, 0, time.UTC),
			expected: false,
		},
		{
			name:     "December 31st is not year start",
			input:    time.Date(2020, time.December, 31, 23, 59, 59, 0, time.UTC),
			expected: false,
		},
		{
			name:     "January 2nd is not year start",
			input:    time.Date(2020, time.January, 2, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			res := datetime.IsBeginingYear(tt.input)
			assert.Equal(t, tt.expected, res)
		})
	}
}
