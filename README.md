# countdown #

This repository helps counts down for a specified duration or to a specified day

You should specify either a duration, or an end date.  

Additionally, if you use a duration, you can specify a start date.

## Clone the project

```
$ git clone https://github.com/KevinStrong/countdown
$ cd countdown
```

Set up a 10 day countdown, then access it using Get().

This decrements every 24 hours starting at the time you created the countdown.
```
cd := countdown.New(countdown.Days(10))
daysRemaining := cd.Get() 
```

You can control when countdown decrements by setting the start time.
``` 
startTime := time.Now()
cd := countdown.New(countdown.Start(startTime), countdown.Days(10))
```

You can have the countdown display the number of hours, minutes, seconds or milliseconds remaining instead of days remaing.
```
cd := countdown.New(countdown.Days(10), countdown.Unit(time.Hour))
```

Save a countdown to a file using Save().
You can resume it later via a construtor option.
```
err := cd.Save()
// In a different program
resumedCountdown := countdown.New(countdown.FromFile())
```
