package internal

import (
	"sync"
	"time"
)

// mimics senders.MetricSender to avoid circular dependency
type internalSender interface {
	// Sends a single metric to Wavefront with optional timestamp and tags.
	SendMetric(name string, value float64, ts int64, source string, tags map[string]string) error

	// Sends a delta counter (counter aggregated at the Wavefront service) to Wavefront.
	// the timestamp for a delta counter is assigned at the server side.
	SendDeltaCounter(name string, value float64, source string, tags map[string]string) error
}

type IncrementerDecrementer interface {
	Inc()
	Dec()
}

type noOpIncrementerDecrementer struct{}

func (n noOpIncrementerDecrementer) Inc() {
	//TODO implement me
	panic("implement me")
}

func (n noOpIncrementerDecrementer) Dec() {
	//TODO implement me
	panic("implement me")
}

type MetricRegistry interface {
	Start()
	Stop()

	NewGauge(string, func() int64)

	PointsInvalid() IncrementerDecrementer
	PointsValid() IncrementerDecrementer
	PointsDropped() IncrementerDecrementer

	HistogramsInvalid() IncrementerDecrementer
	HistogramsValid() IncrementerDecrementer
	HistogramsDropped() IncrementerDecrementer

	SpansInvalid() IncrementerDecrementer
	SpansValid() IncrementerDecrementer
	SpansDropped() IncrementerDecrementer

	SpanLogsInvalid() IncrementerDecrementer
	SpanLogsValid() IncrementerDecrementer
	SpanLogsDropped() IncrementerDecrementer

	EventsInvalid() IncrementerDecrementer
	EventsValid() IncrementerDecrementer
	EventsDropped() IncrementerDecrementer
}

type NoOpRegistry struct {
}

func (n *NoOpRegistry) PointsDropped() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (n *NoOpRegistry) HistogramsInvalid() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (n *NoOpRegistry) HistogramsValid() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (n *NoOpRegistry) HistogramsDropped() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (n *NoOpRegistry) SpansInvalid() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (n *NoOpRegistry) SpansValid() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (n *NoOpRegistry) SpansDropped() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (n *NoOpRegistry) SpanLogsInvalid() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (n *NoOpRegistry) SpanLogsValid() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (n *NoOpRegistry) SpanLogsDropped() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (n *NoOpRegistry) EventsInvalid() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (n *NoOpRegistry) EventsValid() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (n *NoOpRegistry) EventsDropped() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (n *NoOpRegistry) NewGauge(s string, f func() int64) {
	//TODO implement me
	panic("implement me")
}

func (n *NoOpRegistry) PointsInvalid() IncrementerDecrementer {
	return noOpIncrementerDecrementer{}
}

func (n *NoOpRegistry) PointsValid() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (n *NoOpRegistry) Start() {
}

func (n *NoOpRegistry) Stop() {
}

func NewNoOpRegistry() MetricRegistry {
	return &NoOpRegistry{}
}

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

func (registry *RealMetricRegistry) PointsInvalid() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

func (registry *RealMetricRegistry) PointsValid() IncrementerDecrementer {
	//TODO implement me
	panic("implement me")
}

type RegistryOption func(*RealMetricRegistry)

func SetSource(source string) RegistryOption {
	return func(registry *RealMetricRegistry) {
		registry.source = source
	}
}

func SetInterval(interval time.Duration) RegistryOption {
	return func(registry *RealMetricRegistry) {
		registry.reportTicker = time.NewTicker(interval)
	}
}

func SetTags(tags map[string]string) RegistryOption {
	return func(registry *RealMetricRegistry) {
		registry.tags = tags
	}
}

func SetTag(key, value string) RegistryOption {
	return func(registry *RealMetricRegistry) {
		if registry.tags == nil {
			registry.tags = make(map[string]string)
		}
		registry.tags[key] = value
	}
}

func SetPrefix(prefix string) RegistryOption {
	return func(registry *RealMetricRegistry) {
		registry.prefix = prefix
	}
}

func NewMetricRegistry(sender internalSender, setters ...RegistryOption) *RealMetricRegistry {
	registry := &RealMetricRegistry{
		sender:       sender,
		metrics:      make(map[string]interface{}),
		reportTicker: time.NewTicker(time.Second * 60),
		done:         make(chan struct{}),
	}
	for _, setter := range setters {
		setter(registry)
	}
	return registry
}

func (registry *RealMetricRegistry) NewCounter(name string) *MetricCounter {
	return registry.getOrAdd(name, &MetricCounter{}).(*MetricCounter)
}

func (registry *RealMetricRegistry) NewDeltaCounter(name string) *DeltaCounter {
	return registry.getOrAdd(name, &DeltaCounter{MetricCounter{}}).(*DeltaCounter)
}

func (registry *RealMetricRegistry) NewGauge(name string, f func() int64) *FunctionalGauge {
	return registry.getOrAdd(name, &FunctionalGauge{value: f}).(*FunctionalGauge)
}

func (registry *RealMetricRegistry) NewGaugeFloat64(name string, f func() float64) *FunctionalGaugeFloat64 {
	return registry.getOrAdd(name, &FunctionalGaugeFloat64{value: f}).(*FunctionalGaugeFloat64)
}

func (registry *RealMetricRegistry) Start() {
	go registry.start()
}

func (registry *RealMetricRegistry) start() {
	for {
		select {
		case <-registry.reportTicker.C:
			registry.report()
		case <-registry.done:
			return
		}
	}
}

func (registry *RealMetricRegistry) Stop() {
	registry.reportTicker.Stop()
	registry.done <- struct{}{}
}

func (registry *RealMetricRegistry) report() {
	registry.mtx.Lock()
	defer registry.mtx.Unlock()

	for k, metric := range registry.metrics {
		switch metric.(type) {
		case *DeltaCounter:
			deltaCount := metric.(*DeltaCounter).count()
			registry.sender.SendDeltaCounter(registry.prefix+"."+k, float64(deltaCount), "", registry.tags)
			metric.(*DeltaCounter).dec(deltaCount)
		case *MetricCounter:
			registry.sender.SendMetric(registry.prefix+"."+k, float64(metric.(*MetricCounter).count()), 0, "", registry.tags)
		case *FunctionalGauge:
			registry.sender.SendMetric(registry.prefix+"."+k, float64(metric.(*FunctionalGauge).instantValue()), 0, "", registry.tags)
		case *FunctionalGaugeFloat64:
			registry.sender.SendMetric(registry.prefix+"."+k, metric.(*FunctionalGaugeFloat64).instantValue(), 0, "", registry.tags)
		}
	}
}

func (registry *RealMetricRegistry) getOrAdd(name string, metric interface{}) interface{} {
	registry.mtx.Lock()
	defer registry.mtx.Unlock()

	if val, ok := registry.metrics[name]; ok {
		return val
	}
	registry.metrics[name] = metric
	return metric
}
