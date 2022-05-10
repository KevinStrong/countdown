package main

import (
	"fmt"
	"time"

	"github.com/KevinStrong/daycountdown"
)

func main() {
	countdown := daycountdown.New(
		daycountdown.End(time.Date(2068, 1, 0, 0, 0, 0, 0, time.UTC)))

	days := countdown.Get()
	fmt.Printf("Days Until Jan 1 2068: %d", days)
}
