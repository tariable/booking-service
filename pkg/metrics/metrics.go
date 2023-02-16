package metrics

import (
	"bookingService/config"
	"sync"
)

// pretty simple metrics. collect information about quantity of requests.

type Metrics interface {
	IncreaseCounter(topic string)
}

type DefaultMetrics struct {
	getOrdersMutex   sync.RWMutex
	getOrdersCounter int64
	makeOrderMutex   sync.RWMutex
	makeOrderCounter int64
}

func (metrics *DefaultMetrics) IncreaseCounter(topic string) {
	switch topic {
	case "get-orders":
		metrics.getOrdersMutex.Lock()
		metrics.getOrdersCounter++
		metrics.getOrdersMutex.Unlock()
	case "make-order":
		metrics.makeOrderMutex.Lock()
		metrics.getOrdersCounter++
		metrics.makeOrderMutex.Unlock()
	}
}

func New(config config.Metrics) (Metrics, error) {
	metrics := &DefaultMetrics{getOrdersCounter: 0, makeOrderCounter: 0}
	return metrics, nil
}
