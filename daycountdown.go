package daycountdown

import (
	"time"
)

type Countdown struct {
	duration time.Duration
	start    time.Time
	end      time.Time
}

type Setup func(*Countdown)

func Duration(duration int) Setup {
	return func(c *Countdown) {
		c.duration = time.Duration(duration) * time.Hour * 24
	}
}

func Start(start time.Time) Setup {
	return func(c *Countdown) {
		c.start = roundToDay(start).UTC()
	}
}

func End(end time.Time) Setup {
	return func(c *Countdown) {
		c.end = roundToDay(end).UTC()
	}
}

func roundToDay(original time.Time) time.Time {
	return time.Date(original.Year(), original.Month(), original.Day(), 0, 0, 0, 0, original.Location())
}

func New(setups ...Setup) Countdown {
	countdown := Countdown{start: time.Now().UTC()}

	for i := range setups {
		setups[i](&countdown)
	}

	if countdown.end.IsZero() {
		countdown.end = countdown.start.Add(countdown.duration)
	}

	return countdown
}

func (c Countdown) Get() int {
	return int(c.end.Sub(time.Now().UTC()).Hours()) / 24
}
