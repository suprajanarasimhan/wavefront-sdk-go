package internal

import (
	"sync"
	"time"
)

// metric registry for internal metrics
type RealMetricRegistry struct {
	source       string
	prefix       string
	tags         map[string]string
	reportTicker *time.Ticker
	sender       internalSender
	done         chan struct{}

	mtx     sync.Mutex
	metrics map[string]interface{}

	pointsValid   *DeltaCounter
	pointsInvalid *DeltaCounter
	pointsDropped *DeltaCounter

	histogramsValid   *DeltaCounter
	histogramsInvalid *DeltaCounter
	histogramsDropped *DeltaCounter

	spansValid   *DeltaCounter
	spansInvalid *DeltaCounter
	spansDropped *DeltaCounter

	spanLogsValid   *DeltaCounter
	spanLogsInvalid *DeltaCounter
	spanLogsDropped *DeltaCounter

	eventsValid   *DeltaCounter
	eventsInvalid *DeltaCounter
	eventsDropped *DeltaCounter
}

func (registry *RealMetricRegistry) PointsDropped() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (registry *RealMetricRegistry) HistogramsInvalid() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (registry *RealMetricRegistry) HistogramsValid() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (registry *RealMetricRegistry) HistogramsDropped() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (registry *RealMetricRegistry) SpansInvalid() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (registry *RealMetricRegistry) SpansValid() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (registry *RealMetricRegistry) SpansDropped() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (registry *RealMetricRegistry) SpanLogsInvalid() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (registry *RealMetricRegistry) SpanLogsValid() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (registry *RealMetricRegistry) SpanLogsDropped() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (registry *RealMetricRegistry) EventsInvalid() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (registry *RealMetricRegistry) EventsValid() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (registry *RealMetricRegistry) EventsDropped() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (registry *RealMetricRegistry) PointsInvalid() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (registry *RealMetricRegistry) PointsValid() Incrementer {
	return registry.pointsValid
}
