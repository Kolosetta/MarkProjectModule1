package events

import "log"

type Event struct {
	Type string
	Data interface{}
}

var eventChan = make(chan Event, 100)

func StartLogger() {
	go func() {
		for event := range eventChan {
			log.Printf("[EVENT] %s: %+v\n\n", event.Type, event.Data)
		}
	}()
}

func LogEvent(Type string, data interface{}) {
	eventChan <- Event{Type: Type, Data: data}
}
