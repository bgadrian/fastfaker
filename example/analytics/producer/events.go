package producer

type EventGenerator func(count int) []Event

// Event identifies an Analytics Event
type Event interface {
	Name() string
	Properties() map[string]string
	AddProperty(key, value string)
}

// NewSimpleEvent creates a simple event based on input
func NewSimpleEvent(name string, props map[string]string) Event {
	e := &simpleEvent{}
	props["event"] = name
	e.props = props
	return e
}

type simpleEvent struct {
	props map[string]string
}

func (e *simpleEvent) Name() string {
	return e.props["event"]
}

func (e *simpleEvent) Properties() map[string]string {
	clone := make(map[string]string, len(e.props))
	for k, v := range e.props {
		clone[k] = v
	}
	return clone
}

func (e *simpleEvent) AddProperty(key, value string) {
	e.props[key] = value
}
