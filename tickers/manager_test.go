package tickers_test

import (
	"testing"
	"time"

	"github.com/huangsuwei/mygotool/tickers"
	"github.com/huangsuwei/mygotool/tickers/examples"
)

func TestNewTickManage(t *testing.T) {
	tickers.SetupTickManager(
		examples.GetTickerConfig(),
	)

	defer tickers.TickManager().Stop()

	time.Sleep(time.Second * 5)
}
