package examples

import (
	"fmt"
	"time"

	"github.com/huangsuwei/mygotool/tickers"
)

const TickerName = "example_test"

// GetTickerConfig
func GetTickerConfig() tickers.TickConfig {
	return tickers.TickConfig{
		Name:     TickerName,
		Interval: 1 * time.Second,
		Handler:  getTickHandler(),
	}
}

// getTickHandler
func getTickHandler() func(time.Time) {
	return func(time.Time) { fmt.Println("this is a test print") }
}
