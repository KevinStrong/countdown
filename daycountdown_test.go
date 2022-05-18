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

	expect.Equal(got, int64(want))
}

func Test_daycountdown_SetDays_CallGet_ReturnDurationInDays(t *testing.T) {
	expect := is.New(t)

	want := 10
	countdown := daycountdown.New(daycountdown.Days(want))

	got := countdown.Get()

	expect.Equal(got, int64(want))
}

func Test_daycountdown_SetSpecificDuration_CallGet_ReturnDuration(t *testing.T) {
	expect := is.New(t)

	want := 10
	wantDuration := time.Duration(want) * time.Hour * 24
	countdown := daycountdown.New(daycountdown.Duration(wantDuration))

	got := countdown.Get()

	expect.Equal(got, int64(want))
}

func Test_daycountdown_SetEnd_ReturnsDaysUntilEnd(t *testing.T) {
	expect := is.New(t)

	want := 5
	end := getFiveDaysFromNow()
	countdown := daycountdown.New(daycountdown.End(end))

	got := countdown.Get()

	expect.Equal(got, int64(want))
}

func Test_daycountdown_SetStart_AndDuration_ReturnsDaysUntilDurationEnds(t *testing.T) {
	expect := is.New(t)

	want := 5
	duration := 10
	start := getFiveDaysAgo()
	countdown := daycountdown.New(daycountdown.Start(start), daycountdown.Days(duration))

	got := countdown.Get()

	expect.Equal(got, int64(want))
}

// Ignore duration if end is set
func Test_daycountdown_SetEnd_AndSetDuration_ReturnsDaysUntilEnd(t *testing.T) {
	expect := is.New(t)

	want := 5
	end := getFiveDaysFromNow()
	ignored := 10

	countdown := daycountdown.New(daycountdown.End(end), daycountdown.Days(ignored))
	got := countdown.Get()

	expect.Equal(got, int64(want))
}

func Test_daycountdown_SetDuration_AndHours_ReturnDurationInHours(t *testing.T) {
	expect := is.New(t)

	wantDays := 10
	want := 10 * 24
	countdown := daycountdown.New(daycountdown.Days(wantDays), daycountdown.Unit(time.Hour))

	got := countdown.Get()

	expect.Equal(int64(want), got)
}

func Test_daycountdown_SetEndFromConfig_ReturnsDaysUntilEnd(t *testing.T) {
	expect := is.New(t)

	want := 5

	countdown := daycountdown.New(daycountdown.FromFile())
	got := countdown.Get()

	expect.Equal(got, int64(want))
}

func Test_daycountdown_WriteEndToConfig_AndReadInConfig(t *testing.T) {
	expect := is.New(t)

	want := 5
	end := getFiveDaysFromNow()
	countdown := daycountdown.New(daycountdown.End(end))

	err := countdown.Save()
	expect.NoErr(err)

	countdown = daycountdown.New(daycountdown.FromFile())
	got := countdown.Get()

	expect.Equal(got, int64(want))
}

func getFiveDaysAgo() time.Time {
	return time.Now().AddDate(0, 0, -5)
}

func getFiveDaysFromNow() time.Time {
	return time.Now().AddDate(0, 0, 5)
}
