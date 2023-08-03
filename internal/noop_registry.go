package internal

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

func (n *NoOpRegistry) NewGauge(s string, f func() int64) *FunctionalGauge {
	//TODO implement me
	panic("implement me")
}

func (n *NoOpRegistry) PointsInvalid() IncrementerDecrementer {
	return noOpIncrementerDecrementer{}
}

func (n *NoOpRegistry) PointsValid() Incrementer {
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

type noOpIncrementerDecrementer struct{}

func (n noOpIncrementerDecrementer) Inc() {
	//TODO implement me
	panic("implement me")
}

func (n noOpIncrementerDecrementer) Dec() {
	//TODO implement me
	panic("implement me")
}
