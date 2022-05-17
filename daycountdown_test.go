package daycountdown_test

import (
	"testing"
	"time"

	"github.com/KevinStrong/daycountdown"

	"github.com/matryer/is"
)

func Test_daycountdown_CallGet_ReturnZero(t *testing.T) {
	expect := is.New(t)

	want := 0
	countdown := daycountdown.New()

	got := countdown.Get()

	expect.Equal(want, got)
}

func Test_daycountdown_SetDuration_CallGet_ReturnDuration(t *testing.T) {
	expect := is.New(t)

	want := 10
	countdown := daycountdown.New(daycountdown.Duration(want))

	got := countdown.Get()

	expect.Equal(want, got)
}

func Test_daycountdown_SetEnd_ReturnsDaysUntilEnd(t *testing.T) {
	expect := is.New(t)

	want := 5
	end := getFiveDaysFromNow()
	countdown := daycountdown.New(daycountdown.End(end))

	got := countdown.Get()

	expect.Equal(got, want)
}

func Test_daycountdown_SetStart_AndDuration_ReturnsDaysUntilDurationEnds(t *testing.T) {
	expect := is.New(t)

	want := 5
	duration := 10
	start := getFiveDaysAgo()
	countdown := daycountdown.New(daycountdown.Start(start), daycountdown.Duration(duration))

	got := countdown.Get()

	expect.Equal(got, want)
}

func Test_daycountdown_SetEnd_AndSetDuration_IgnoreDuration(t *testing.T) {
	expect := is.New(t)

	want := 5
	end := getFiveDaysFromNow()
	ignored := 10

	countdown := daycountdown.New(daycountdown.End(end), daycountdown.Duration(ignored))
	got := countdown.Get()

	expect.Equal(got, want)
}

func getFiveDaysAgo() time.Time {
	return time.Now().AddDate(0, 0, -5)
}

func getFiveDaysFromNow() time.Time {
	return time.Now().AddDate(0, 0, 5)
}
