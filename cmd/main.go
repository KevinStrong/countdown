package main

import (
	"fmt"
	"github.com/KevinStrong/countdown"
	"time"
)

func main() {
	// make a constructor option that reads in a file
	// make a constructor option that reads in a command line flag

	cd := countdown.New(
		countdown.End(time.Date(2068, 1, 0, 0, 0, 0, 0, time.UTC)))

	days := cd.Get()
	fmt.Printf("Days Until Jan 1 2068: %d", days)
}
