package daycountdown

import (
	"math"
	"time"
)

type Countdown struct {
	duration time.Duration
	start    time.Time
	end      time.Time
	units    time.Duration
}

type Setup func(*Countdown)

func Duration(duration time.Duration) Setup {
	return func(c *Countdown) {
		c.duration = duration
	}
}

func Days(days int) Setup {
	return func(c *Countdown) {
		c.duration = time.Duration(days) * time.Hour * 24
	}
}

func Start(start time.Time) Setup {
	return func(c *Countdown) {
		c.start = start.UTC()
	}
}

func End(end time.Time) Setup {
	return func(c *Countdown) {
		c.end = end.UTC()
	}
}

func Unit(unit time.Duration) Setup {
	return func(c *Countdown) {
		if unit != time.Hour &&
			unit != time.Minute &&
			unit != time.Second &&
			unit != time.Millisecond {
			return
		}
		c.units = unit
	}
}

func New(setups ...Setup) Countdown {
	countdown := Countdown{
		start: time.Now().UTC(),
	}

	for i := range setups {
		setups[i](&countdown)
	}

	if countdown.end.IsZero() {
		countdown.end = countdown.start.Add(countdown.duration)
	}

	return countdown
}

func (c Countdown) Get() int64 {
	durationRemaining := c.end.Sub(time.Now().UTC())
	var unitsRemaining int64
	// day is default behavior
	if c.units == 0 {
		unitsRemaining = int64(math.Ceil(durationRemaining.Hours() / 24))
	}
	if c.units == time.Hour {
		unitsRemaining = int64(math.Ceil(durationRemaining.Hours()))
	}
	if c.units == time.Minute {
		unitsRemaining = int64(math.Ceil(durationRemaining.Minutes()))
	}
	if c.units == time.Second {
		unitsRemaining = int64(math.Ceil(durationRemaining.Seconds()))
	}
	if c.units == time.Millisecond {
		unitsRemaining = durationRemaining.Milliseconds()
	}
	if unitsRemaining < 0 {
		return 0
	}
	return unitsRemaining
}
