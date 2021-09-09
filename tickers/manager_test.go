package tickers_test

import (
	"mygotool/tickers"
	"testing"
	"time"

	"mygotool/tickers/examples"
)

func TestNewTickManage(t *testing.T) {
	tickers.SetupTickManager(
		examples.GetTickerConfig(),
	)

	defer tickers.TickManager().Stop()

	time.Sleep(time.Second * 5)
}
