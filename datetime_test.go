package datetime_test

import (
	"testing"
	"time"

	"github.com/alex-cos/datetime"
	"github.com/stretchr/testify/assert"
)

func TestSimpleFormat(t *testing.T) {
	t.Parallel()

	date := time.Date(2020, time.February, 14, 12, 12, 12, 0, time.FixedZone("Paris", 3600))
	expected := "2020-02-14 11:12:12"
	res := datetime.SimpleFormat(date)
	assert.Equal(t, expected, res)
}

func TestFormatISO8601Millis(t *testing.T) {
	t.Parallel()

	date := time.Date(2020, time.February, 14, 12, 12, 12, 0, time.UTC)
	expected := "2020-02-14T12:12:12.000Z"
	res := datetime.FormatISO8601Millis(date)
	assert.Equal(t, expected, res)
}

func TestParseISO8601Millis(t *testing.T) {
	t.Parallel()

	date := "2020-02-14T12:12:12.000Z"
	expected := time.Date(2020, time.February, 14, 12, 12, 12, 0, time.UTC)
	res, err := datetime.ParseISO8601Millis(date)
	assert.NoError(t, err)
	assert.Equal(t, expected, res)

	date = "2020-02-14 11:12:12"
	_, err = datetime.ParseISO8601Millis(date)
	assert.Error(t, err)
}

func TestMidNightUTC(t *testing.T) {
	t.Parallel()

	date := time.Date(2020, time.February, 14, 12, 12, 12, 0, time.Local)
	expected := time.Date(2020, time.February, 14, 0, 0, 0, 0, time.UTC)
	res := datetime.MidNightUTC(date)
	assert.Equal(t, expected, res)

	date = time.Date(2020, time.February, 14, 0, 35, 0, 0, time.Local)
	expected = time.Date(2020, time.February, 13, 0, 0, 0, 0, time.UTC)
	res = datetime.MidNightUTC(date)
	assert.Equal(t, expected, res)

	date = time.Date(2020, time.February, 14, 0, 35, 0, 0, time.UTC)
	expected = time.Date(2020, time.February, 14, 0, 0, 0, 0, time.UTC)
	res = datetime.MidNightUTC(date)
	assert.Equal(t, expected, res)
}

func TestIsMidNightUTC(t *testing.T) {
	t.Parallel()

	date := time.Date(2020, time.February, 14, 12, 12, 12, 0, time.Local)
	res := datetime.IsMidNightUTC(date)
	assert.False(t, res)

	date = time.Date(2020, time.February, 1, 1, 12, 0, 0, time.Local)
	res = datetime.IsMidNightUTC(date)
	assert.True(t, res)

	date = time.Date(2020, time.February, 14, 0, 35, 0, 0, time.UTC)
	res = datetime.IsMidNightUTC(date)
	assert.True(t, res)
}

func TestLastMonth(t *testing.T) {
	t.Parallel()

	date := time.Date(2020, time.February, 14, 12, 12, 12, 0, time.Local)
	expected := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	res := datetime.LastMonth(date)
	assert.Equal(t, expected, res)

	date = time.Date(2020, time.January, 14, 12, 12, 12, 0, time.Local)
	expected = time.Date(2019, time.December, 1, 0, 0, 0, 0, time.UTC)
	res = datetime.LastMonth(date)
	assert.Equal(t, expected, res)
}

func TestThisMonth(t *testing.T) {
	t.Parallel()

	date := time.Date(2020, time.February, 14, 12, 12, 12, 0, time.Local)
	expected := time.Date(2020, time.February, 1, 0, 0, 0, 0, time.UTC)
	res := datetime.ThisMonth(date)
	assert.Equal(t, expected, res)

	date = time.Date(2020, time.January, 14, 12, 12, 12, 0, time.Local)
	expected = time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	res = datetime.ThisMonth(date)
	assert.Equal(t, expected, res)
}

func TestNextMonth(t *testing.T) {
	t.Parallel()

	date := time.Date(2020, time.February, 14, 12, 12, 12, 0, time.Local)
	expected := time.Date(2020, time.March, 1, 0, 0, 0, 0, time.UTC)
	res := datetime.NextMonth(date)
	assert.Equal(t, expected, res)

	date = time.Date(2019, time.December, 14, 12, 12, 12, 0, time.Local)
	expected = time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	res = datetime.NextMonth(date)
	assert.Equal(t, expected, res)
}

func TestLastQuarter(t *testing.T) {
	t.Parallel()

	date := time.Date(2020, time.February, 14, 12, 12, 12, 0, time.Local)
	expected := time.Date(2019, time.October, 1, 0, 0, 0, 0, time.UTC)
	res := datetime.LastQuarter(date)
	assert.Equal(t, expected, res)
}

func TestThisQuarter(t *testing.T) {
	t.Parallel()

	date := time.Date(2020, time.February, 14, 12, 12, 12, 0, time.Local)
	expected := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	res := datetime.ThisQuarter(date)
	assert.Equal(t, expected, res)

	date = time.Date(2020, time.January, 2, 12, 12, 12, 0, time.Local)
	expected = time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	res = datetime.ThisQuarter(date)
	assert.Equal(t, expected, res)
}

func TestNextQuarter(t *testing.T) {
	t.Parallel()

	date := time.Date(2020, time.February, 14, 12, 12, 12, 0, time.Local)
	expected := time.Date(2020, time.April, 1, 0, 0, 0, 0, time.UTC)
	res := datetime.NextQuarter(date)
	assert.Equal(t, expected, res)

	date = time.Date(2020, time.December, 14, 12, 12, 12, 0, time.Local)
	expected = time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)
	res = datetime.NextQuarter(date)
	assert.Equal(t, expected, res)
}

func TestLastYear(t *testing.T) {
	t.Parallel()

	date := time.Date(2020, time.February, 14, 12, 12, 12, 0, time.Local)
	expected := time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)
	res := datetime.LastYear(date)
	assert.Equal(t, expected, res)
}

func TestThisYear(t *testing.T) {
	t.Parallel()

	date := time.Date(2020, time.February, 14, 12, 12, 12, 0, time.Local)
	expected := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	res := datetime.ThisYear(date)
	assert.Equal(t, expected, res)
}

func TestNextYear(t *testing.T) {
	t.Parallel()

	date := time.Date(2020, time.February, 14, 12, 12, 12, 0, time.Local)
	expected := time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)
	res := datetime.NextYear(date)
	assert.Equal(t, expected, res)
}

func TestIsBeginingMonth(t *testing.T) {
	t.Parallel()

	date := time.Date(2020, time.January, 1, 12, 12, 12, 0, time.UTC)
	res := datetime.IsBeginingMonth(date)
	assert.True(t, res)

	date = time.Date(2020, time.February, 1, 12, 12, 12, 0, time.UTC)
	res = datetime.IsBeginingMonth(date)
	assert.True(t, res)

	date = time.Date(2020, time.September, 2, 12, 12, 12, 0, time.UTC)
	res = datetime.IsBeginingMonth(date)
	assert.False(t, res)
}

func TestIsBeginingQuarter(t *testing.T) {
	t.Parallel()

	date := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	res := datetime.IsBeginingQuarter(date)
	assert.True(t, res)

	date = time.Date(2020, time.April, 1, 0, 0, 0, 0, time.UTC)
	res = datetime.IsBeginingQuarter(date)
	assert.True(t, res)

	date = time.Date(2020, time.July, 1, 0, 0, 0, 0, time.UTC)
	res = datetime.IsBeginingQuarter(date)
	assert.True(t, res)

	date = time.Date(2020, time.October, 1, 0, 0, 0, 0, time.UTC)
	res = datetime.IsBeginingQuarter(date)
	assert.True(t, res)

	date = time.Date(2020, time.February, 1, 0, 0, 0, 0, time.UTC)
	res = datetime.IsBeginingQuarter(date)
	assert.False(t, res)

	date = time.Date(2020, time.April, 11, 0, 0, 0, 0, time.UTC)
	res = datetime.IsBeginingQuarter(date)
	assert.False(t, res)
}

func TestIsBeginingYear(t *testing.T) {
	t.Parallel()

	date := time.Date(2020, time.January, 1, 12, 12, 12, 0, time.UTC)
	res := datetime.IsBeginingYear(date)
	assert.True(t, res)

	date = time.Date(2020, time.February, 1, 12, 12, 12, 0, time.UTC)
	res = datetime.IsBeginingYear(date)
	assert.False(t, res)

	date = time.Date(2020, time.September, 11, 12, 12, 12, 0, time.UTC)
	res = datetime.IsBeginingYear(date)
	assert.False(t, res)
}
