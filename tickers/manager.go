package tickers

import (
	"sync"
	"time"

	"mygotool/task"
)

const (
	// DefaultInterval 默认更新间隔最少为1分钟
	DefaultInterval = 1 * time.Minute
)

type (
	TickName string

	TickConfig struct {
		Name     TickName
		Interval time.Duration
		Handler  func(time.Time)
	}

	TickManage struct {
		lock    sync.Mutex
		tickers map[TickName]*task.RepeatedTask
	}
)

var tickManager *TickManage

// TickManager
func TickManager() *TickManage {
	return tickManager
}

// SetupTickManager
func SetupTickManager(tickConfigs ...TickConfig) {
	if tickManager == nil {
		tickManager = &TickManage{
			tickers: make(map[TickName]*task.RepeatedTask),
		}

		for _, tickConfig := range tickConfigs {
			tickManager.RegisterTicker(tickConfig)
		}
	}
}

// RegisterTicker
func (t *TickManage) RegisterTicker(config TickConfig) {
	t.lock.Lock()
	defer t.lock.Unlock()

	if _, ok := t.tickers[config.Name]; !ok {
		if config.Interval < DefaultInterval {
			config.Interval = DefaultInterval
		}
		repeatedTask := task.NewRepeatedTask(config.Interval, config.Handler)
		repeatedTask.Start()
		t.tickers[config.Name] = repeatedTask
	}
}

// DeleteTicker
func (t *TickManage) DeleteTicker(name TickName) {
	t.lock.Lock()
	defer t.lock.Unlock()

	a, ok := t.tickers[name]
	if ok {
		a.Stop()
		delete(t.tickers, name)
	}
}

// Tickers
func (t *TickManage) Tickers() map[TickName]*task.RepeatedTask {
	return t.tickers
}

// Stop
func (t *TickManage) Stop() {
	for name := range t.tickers {
		t.DeleteTicker(name)
	}
	t.tickers = nil
}
