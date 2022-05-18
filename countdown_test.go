package countdown_test

import (
	"testing"
	"time"

	"github.com/matryer/is"

	"github.com/KevinStrong/countdown"
)

func Test_countdown_CallGet_ReturnZero(t *testing.T) {
	expect := is.New(t)

	want := 0
	cd := countdown.New()

	got := cd.Get()

	expect.Equal(got, int64(want))
}

func Test_countdown_SetDays_CallGet_ReturnDurationInDays(t *testing.T) {
	expect := is.New(t)

	want := 10
	cd := countdown.New(countdown.Days(want))

	got := cd.Get()

	expect.Equal(got, int64(want))
}

func Test_countdown_SetSpecificDuration_CallGet_ReturnDuration(t *testing.T) {
	expect := is.New(t)

	want := 10
	wantDuration := time.Duration(want) * time.Hour * 24
	cd := countdown.New(countdown.Duration(wantDuration))

	got := cd.Get()

	expect.Equal(got, int64(want))
}

func Test_countdown_SetEnd_ReturnsDaysUntilEnd(t *testing.T) {
	expect := is.New(t)

	want := 5
	end := getFiveDaysFromNow()
	cd := countdown.New(countdown.End(end))

	got := cd.Get()

	expect.Equal(got, int64(want))
}

func Test_countdown_SetStart_AndDuration_ReturnsDaysUntilDurationEnds(t *testing.T) {
	expect := is.New(t)

	want := 5
	duration := 10
	start := getFiveDaysAgo()
	cd := countdown.New(countdown.Start(start), countdown.Days(duration))

	got := cd.Get()

	expect.Equal(got, int64(want))
}

// Ignore duration if end is set
func Test_countdown_SetEnd_AndSetDuration_ReturnsDaysUntilEnd(t *testing.T) {
	expect := is.New(t)

	want := 5
	end := getFiveDaysFromNow()
	ignored := 10

	cd := countdown.New(countdown.End(end), countdown.Days(ignored))
	got := cd.Get()

	expect.Equal(got, int64(want))
}

func Test_countdown_SetDuration_AndHours_ReturnDurationInHours(t *testing.T) {
	expect := is.New(t)

	wantDays := 10
	want := 10 * 24
	cd := countdown.New(countdown.Days(wantDays), countdown.Unit(time.Hour))

	got := cd.Get()

	expect.Equal(int64(want), got)
}

func Test_countdown_SetEndFromConfig_ReturnsDaysUntilEnd(t *testing.T) {
	expect := is.New(t)

	want := 5

	cd := countdown.New(countdown.FromFile())
	got := cd.Get()

	expect.Equal(got, int64(want))
}

func Test_countdown_WriteEndToConfig_AndReadInConfig(t *testing.T) {
	expect := is.New(t)

	want := 5
	end := getFiveDaysFromNow()
	cd := countdown.New(countdown.End(end))

	err := cd.Save()
	expect.NoErr(err)

	cd = countdown.New(countdown.FromFile())
	got := cd.Get()

	expect.Equal(got, int64(want))
}

func getFiveDaysAgo() time.Time {
	return time.Now().AddDate(0, 0, -5)
}

func getFiveDaysFromNow() time.Time {
	return time.Now().AddDate(0, 0, 5)
}
