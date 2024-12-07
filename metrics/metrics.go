package metrics

import (
	"fmt"
	"time"
)

const (
	reportInterval = time.Second * 30
)

type metrics struct {
	calls_metrics map[string][]time.Time
}

type Metrics interface {
	IncCalls(endpoint string)
	GetReport(endpoint string) int64
	GetBatchReport() map[string]int64
}

func NewMetrics() Metrics {
	return &metrics{calls_metrics: make(map[string][]time.Time, 0)}
}

func (m *metrics) IncCalls(endpoint string) {
	val, ok := m.calls_metrics[endpoint]
	if !ok {
		m.calls_metrics[endpoint] = make([]time.Time, 0)
		val = m.calls_metrics[endpoint]
	}

	val = append(val, time.Now())
	m.calls_metrics[endpoint] = val
}

func (m *metrics) GetReport(endpoint string) int64 {
	values := m.calls_metrics[endpoint]
	res := int64(0)
	for i := 0; i < len(values); i++ {
		if time.Now().Sub(values[i]) < reportInterval {
			res++
		}
	}
	return res
}

func (m *metrics) GetBatchReport() map[string]int64 {
	result := make(map[string]int64, 0)
	for key, values := range m.calls_metrics {
		res := int64(0)

		fmt.Println(values)
		for i := 0; i < len(values); i++ {
			if time.Now().Sub(values[i]) < reportInterval {
				res++
			}
		}
		result[key] = res
	}

	return result
}
