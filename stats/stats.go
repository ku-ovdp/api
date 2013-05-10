package stats

import (
	"sync"
)

type ToNotify interface {
	SendJSON(value interface{}) error
}

var Destination ToNotify

var trigger chan struct{}

var stats struct {
	data map[string]int
	sync.Mutex
}

func ChangeStat(statName string, change int) {
	stats.Lock()
	defer stats.Unlock()
	stats.data[statName] = stats.data[statName] + change
	go func() {
		if Destination != nil {
			Destination.SendJSON(stats.data)
		}
	}()
}

func init() {
	trigger = make(chan struct{})
	stats.data = make(map[string]int)
}
