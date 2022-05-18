package main

import (
	"fmt"
	"time"

	"github.com/KevinStrong/daycountdown"
)

func main() {
	// make a constructor option that reads in a file
	// make a constructor option that reads in a command line flag

	countdown := daycountdown.New(
		daycountdown.End(time.Date(2068, 1, 0, 0, 0, 0, 0, time.UTC)))

	days := countdown.Get()
	fmt.Printf("Days Until Jan 1 2068: %d", days)
}
