package countdown

import (
	"io/ioutil"
	"math"
	"time"
)

type Countdown struct {
	duration   time.Duration
	start      time.Time
	end        time.Time
	units      time.Duration
	configName string
	Error      error
}

const defaultConfigFileName = "duration_countdown.config"
const timeStampFormat = time.RFC3339Nano

type Option func(*Countdown)

func Duration(duration time.Duration) Option {
	return func(c *Countdown) {
		c.duration = duration
	}
}

func Days(days int) Option {
	return func(c *Countdown) {
		c.duration = time.Duration(days) * time.Hour * 24
	}
}

func Start(start time.Time) Option {
	return func(c *Countdown) {
		c.start = start.UTC()
	}
}

func End(end time.Time) Option {
	return func(c *Countdown) {
		c.end = end.UTC()
	}
}

func Unit(unit time.Duration) Option {
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

func FromFile() Option {
	return func(c *Countdown) {
		content, err := ioutil.ReadFile(c.configName)
		if err != nil {
			c.Error = err
			return
		}
		end, err := time.Parse(timeStampFormat, string(content))
		if err != nil {
			c.Error = err
			return
		}
		c.end = end
	}
}

func New(setups ...Option) Countdown {
	countdown := Countdown{
		start:      time.Now().UTC(),
		configName: defaultConfigFileName,
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
	if c.Error != nil {
		return 0
	}
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

func (c Countdown) Save() error {
	endTime := c.end.Format(timeStampFormat)
	return ioutil.WriteFile(c.configName, []byte(endTime), 0644)
}
